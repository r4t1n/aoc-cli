package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/fatih/color"
	"github.com/r4t1n/aoc-cli/date"
	"github.com/r4t1n/aoc-cli/file"
	"github.com/r4t1n/aoc-cli/http"
	"github.com/r4t1n/aoc-cli/path"
)

const (
	baseInputURL = "https://adventofcode.com/%d/day/%d/input"
)

var (
	blue = color.New(color.FgBlue).SprintFunc()
)

func main() {
	// Try to get the date from the working directory
	year, day, err := date.ReturnDate()
	if err != nil {
		log.Fatal(err)
	}

	// Get the users home directory
	userHomeDirectory, err := path.ReturnUserHomeDirectory()
	if err != nil {
		log.Fatal(err)
	}

	// Try to copy input from cache
	cachedInputExists, err := path.CheckForCachedInput(year, day, userHomeDirectory)
	if err != nil {
		log.Fatal(err)
	}

	if cachedInputExists {
		fmt.Printf("Copying input for %s/%s...\n", blue(strconv.Itoa(year)), blue(strconv.Itoa(day)))
		err = file.CopyInput(year, day, userHomeDirectory)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// Get the session cookie
		sessionCookie, err := file.ReturnSessionCookie(userHomeDirectory)
		if err != nil {
			log.Fatal(err)
		}

		inputURL := fmt.Sprintf(baseInputURL, year, day)
		fmt.Printf("Downloading input for %s/%s...\n", blue(strconv.Itoa(year)), blue(strconv.Itoa(day)))

		// Make the HTTP GET request and get the response body
		body, err := http.ReturnBody(inputURL, sessionCookie)
		if err != nil {
			log.Fatal(err)
		}

		// Write the response body to file in the working directory
		err = file.WriteInput(year, day, userHomeDirectory, body)
		if err != nil {
			log.Fatal(err)
		}
	}
}
