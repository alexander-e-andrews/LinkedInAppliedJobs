package main

import (
	"fmt"
	"strings"
	//"github.com/alecthomas/participle/v2"
)

//If I really do this, I should split it out into its own project.

//I will need to become more familiar with curl's grammer to properly flesh this out at the end

type CURLInputMyParse struct {
	URL        string
	Headers    []MyHeaders
	Compressed *bool //So with certain properties they can probably be set to true or false, or we can have them be not set by using a pointer
}

type MyHeaders struct {
	Key   string
	Value string
}

//With this assuming we get the Google chrome copy cURL (bash)
//Also going to hope that ' doesn't appear anywhere and assume its just going to be url encoded.
func ParseMySimpleCurl(body string)(CURLInputMyParse) {
	cip := CURLInputMyParse{}
	cip.Headers = make([]MyHeaders, 0)

	//Going to go line by line
	//Maye be able to save a little effort by spliting on /\n, since each line seems to send with an escaped character except that last. But the last is also a different command
	//Will need to look at more cURL examples to confirm this
	eachLine := strings.Split(body, "\\\r\n")
	if strings.HasPrefix(eachLine[0], "curl") {
		//we can continue
		//strings.TrimSpace()

		urlStartIndex := strings.Index(eachLine[0], "'") + len("'") //Adding so we get past the character and don't include it in the final string
		urlEndIndex := strings.LastIndex(eachLine[0], "'")
		//These kind of commands should probably be ut-8 aware []runes
		cip.URL = eachLine[0][urlStartIndex:urlEndIndex]
	} else {
		fmt.Println("Need to write actual handler, but this curl command is not right")
		return CURLInputMyParse{}
	}

	for x := 1; x < len(eachLine); x++ {
		line := strings.TrimSpace(eachLine[x])

		//Now we handle each line type
		//Assuming that there is only one space.
		splitLine := EscapedSplit(line, " ", []string{"'"})
		switch splitLine[0] {
		case "-H":
			//Passing the header line info
			cip.Headers = append(cip.Headers, MySplitHeaderLine(splitLine[1]))
		case "--compressed":
			cip.Compressed = newTrue()
		}
	}
	return cip
}

func MySplitHeaderLine(headerLine string) (header MyHeaders) {
	startIndex := strings.Index(headerLine, "'") + len("'") //Adding so we get past the character and don't include it in the final string
	EndIndex := strings.LastIndex(headerLine, "'")
	//Just removing any spaces/removing the quoting string
	mainBlock := headerLine[startIndex:EndIndex]
	keyValueSplit := strings.SplitN(mainBlock, ": ", 2)
	header.Key = keyValueSplit[0]
	//Remove that opening space that seems to be there everytime. I don't know about ending spaces though
	header.Value = strings.TrimPrefix(keyValueSplit[1], " ")
	return header
}

//writing a split that will split on a character only if its not being like escaped
//Is regex a real grammer I can't remember I don't think so, it can't count
//Will return an empty array of length 0 if we fail due to things not matching up
func EscapedSplit(s string, sep string, validEscapes []string) (split []string) {
	split = make([]string, 0)
	escapesStack := make([]string, 0)

	lastSplitIndex := 0
	for x := range s {
		//we can check for the sep
		if len(escapesStack) == 0 {
			if s[x:x+len(sep)] == sep {
				split = append(split, s[lastSplitIndex:x])
				lastSplitIndex = x + len(sep)
			}
		}
		for _, esc := range validEscapes {
			//So we know that we matched an escape character, I'm going to ignore the possibility of a \ before hand
			//Although that is something I should really do in the future TODO
			if s[x:x+len(esc)] == esc {
				//If that character was the last one we saw, we can remove it from the escapement, otherwise we need to append it to the end
				if len(escapesStack)> 0 &&escapesStack[len(escapesStack)-1] == esc {
					escapesStack = escapesStack[:len(escapesStack)-1]

				} else {
					escapesStack = append(escapesStack, esc)
				}
				x += len(esc) -1
			}
		}
	}
	split = append(split, s[lastSplitIndex:])

	return split
}

func newTrue() *bool {
	b := true
	return &b
}
