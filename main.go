package main

import (
	"log"

	"github.com/r4t1n/aoc-cli/file"
	"github.com/r4t1n/aoc-cli/http"
	"github.com/r4t1n/aoc-cli/path"
	"github.com/r4t1n/aoc-cli/time"
)

func main() {
	// Get the session cookie file path
	sessionCookieFilePath := path.GetSessionCookieFilePath()
	if sessionCookieFilePath.Err != nil {
		log.Fatal(sessionCookieFilePath.Err)
	}

	// Get the session cookie
	sessionCookie := file.ReadSessionCookie(sessionCookieFilePath.SessionCookieFilePath)
	if sessionCookie.Err != nil {
		log.Fatal(sessionCookie.Err)
	}

	// Try to get the input URL from the working directory
	pathInputURL := path.ReturnInputURL()
	if pathInputURL.Err != nil {
		log.Fatal(pathInputURL.Err)
	}

	// Check if getting the date from the working directory was successful
	var inputURL string
	if len(pathInputURL.InputURL) != 0 {
		inputURL = pathInputURL.InputURL
	} else {
		// Return the input URL from the date
		timeInputURL := time.ReturnInputURL()
		if timeInputURL.Err != nil {
			log.Fatal(timeInputURL.Err)
		}
		inputURL = timeInputURL.InputURL
	}

	// Make the HTTP GET request and get the response body
	httpResponse := http.Get(inputURL, sessionCookie.SessionCookie)
	if httpResponse.Err != nil {
		log.Fatal(httpResponse.Err)
	}

	// Write the response body to file
	err := file.WriteInput(httpResponse.Body)
	if err != nil {
		log.Fatal(err)
	}
}
