package main

import (
	"flag"
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
	blue      = color.New(color.FgBlue).SprintFunc()
	day, year int
)

func init() {
	flag.IntVar(&day, "day", 0, "The day used for the date")
	flag.IntVar(&year, "year", 0, "The year used for the date")
}

func main() {
	flag.Parse()

	var err error
	if day == 0 || year == 0 {
		// Get the date either from the path or the time
		day, year, err = date.ReturnDate()
		if err != nil {
			log.Fatal(err)
		}
	}

	// Get the users home directory
	userHomeDirectory, err := path.ReturnUserHomeDirectory()
	if err != nil {
		log.Fatal(err)
	}

	// Check if the cached input exists
	cachedInputExists, err := path.CheckForCachedInput(year, day, userHomeDirectory)
	if err != nil {
		log.Fatal(err)
	}

	if cachedInputExists {
		fmt.Printf("Copying input for %s/%s...\n", blue(strconv.Itoa(year)), blue(strconv.Itoa(day)))
		err = file.CopyInput(day, year, userHomeDirectory)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// Get the session cookie
		sessionCookie, err := file.ReturnSessionCookie(userHomeDirectory)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Downloading input for %s/%s...\n", blue(strconv.Itoa(year)), blue(strconv.Itoa(day)))

		// Make the HTTP GET request and get the response body
		inputURL := fmt.Sprintf(baseInputURL, year, day)
		body, err := http.ReturnBody(inputURL, sessionCookie)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Writing input for %s/%s...\n", blue(strconv.Itoa(year)), blue(strconv.Itoa(day)))

		// Write the response body to file in the working directory
		err = file.WriteInput(day, year, userHomeDirectory, body)
		if err != nil {
			log.Fatal(err)
		}
	}
}
