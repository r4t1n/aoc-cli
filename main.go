package main

import (
	"fmt"
	"log"

	"github.com/r4t1n/aoc-cli/file"
	"github.com/r4t1n/aoc-cli/http"
	"github.com/r4t1n/aoc-cli/path"
	"github.com/r4t1n/aoc-cli/time"
)

const (
	baseInputURL = "https://adventofcode.com/%d/day/%d/input"
	defaultDay   = 1
)

func main() {
	// Get the session cookie file path
	sessionCookieFilePath := path.GetSessionCookieFilePath()
	if sessionCookieFilePath.Err != nil {
		log.Fatalf("error: %v", sessionCookieFilePath.Err)
	}

	// Get the session cookie
	sessionCookie := file.ReadSessionCookie(sessionCookieFilePath.SessionCookieFilePath)
	if sessionCookie.Err != nil {
		log.Fatalf("error: %v", sessionCookie.Err)
	}

	// Get the year, month and day
	currentTimeAOC := time.Get()
	if currentTimeAOC.Err != nil {
		log.Fatalf("error: %v", currentTimeAOC.Err)
	}

	// Check if it is December and apply the day if true, else fall back to the default day
	var inputURL string
	if currentTimeAOC.Month == "December" {
		inputURL = fmt.Sprintf(baseInputURL, currentTimeAOC.Year, currentTimeAOC.Day)
	} else {
		inputURL = fmt.Sprintf(baseInputURL, currentTimeAOC.Year, defaultDay)
	}

	// Make the HTTP GET request and get the response body
	httpResponse := http.Get(inputURL, sessionCookie.SessionCookie)
	if httpResponse.Err != nil {
		log.Fatalf("error: %v", httpResponse.Err)
	}

	// Write the response body to file
	err := file.WriteInput(httpResponse.Body)
	if err != nil {
		log.Fatalf("error writing input to file: %v", err)
	}
}
