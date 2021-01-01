package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

func main() {
	//Make request to xkcd.
	resp, err := http.Get("https://c.xkcd.com/random/comic/")
	if err != nil {
		log.Fatal(err)
	}

	//Read response.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//Convert body to string.
	sb := string(body)
	REEEEEEEEEEEEEEE := regexp.MustCompile("Image URL.*")
	sb = REEEEEEEEEEEEEEE.FindString(sb)
	sb = sb[38:] //Go's regex is so fucking non-sensical that i had to do this. FUCK GOOGLE.

	//Request image.
	response, err := http.Get(sb)
	if err != nil {
		log.Fatal(err)
	}

	//Create output file.
	outFile, err := os.Create(sb[29:len(sb)])
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	//Copy data from HTTP response to file.
	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
    }
