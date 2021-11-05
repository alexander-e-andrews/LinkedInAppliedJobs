Application aims to make it easier to sort through the linkedin applied jobs that you have.
Currently they are all just jammed in together. Hope to make it easy to see which jobs you have recently applied for so when you get a call back,
you can reference the job description.
Also most features don't work, can only retrieve, and then search by company name.

Sometimes jobs you apply for will not be added to the list, it's just something on LinkedIn's side.
You do need to create a curl.txt file, which is a copy from chrome of the request to actually get the list of jobs you have applied to.
It is used to get you logged in easily on this side. A example is provided in curl_example.txt

General job list is output to response.json, and any search is outputted to output.json. If any jobs do not match your search, output.json will be a blank json array.

Flags:
-l    Load. Loads the response.json instead of querying LinkedIn. If the file is not found, program will exit. Default false
-s    Save. Saves to response.json. Overwrites previous response.json files. Default true