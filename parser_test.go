package main

import (
	"fmt"
	"os"
	"testing"
)

func TestParsingFile(t *testing.T) {
	dat, err := os.ReadFile("./Example.txt")
    if err != nil{
		t.Error(err)
	}
    ParseMySimpleCurl(string(dat))
}

func TestEscapedSplit(t *testing.T){
	testString := `-H 'x-li-track: {"clientVersion":"1.9.5950","mpVersion":"1.9.5950","osName":"web","timezoneOffset":-4,"timezone":"America/New_York","deviceFormFactor":"DESKTOP","mpName":"voyager-web","displayDensity":1,"displayWidth":1920,"displayHeight":1080}'`

	splitStrings := EscapedSplit(testString, " ", []string{"'","\""})
	for x := range splitStrings{
		fmt.Println(splitStrings[x])
	}
}