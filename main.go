package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/r4t1n/aoc-cli/date"
	"github.com/r4t1n/aoc-cli/file"
	"github.com/r4t1n/aoc-cli/http"
)

const (
	baseInputURL = "https://adventofcode.com/%d/day/%d/input"
)

var (
	blue      = color.New(color.FgBlue).SprintFunc()
	red       = color.New(color.FgRed).SprintFunc()
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
	userHomeDirectory, err := file.ReturnUserHomeDirectory()
	if err != nil {
		log.Fatal(err)
	}

	GetInput(day, year, userHomeDirectory)

}

func GetInput(day, year int, userHomeDirectory string) {
	// Check if the cached input exists
	cachedInputExists, err := file.CheckForCachedInput(year, day, userHomeDirectory)
	if err != nil {
		log.Fatal(err)
	}

	if cachedInputExists {
		// Copy the cached input
		fmt.Printf("Copying input for %s/%s...\n", blue(strconv.Itoa(year)), blue(strconv.Itoa(day)))
		err = file.CopyInput(day, year, userHomeDirectory)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// Check if the session cookie exists
		sessionCookieExists, err := file.CheckForSessionCookie(userHomeDirectory)
		if err != nil {
			log.Fatal(err)
		}

		// If the session cookie does not exist notify the user
		if !sessionCookieExists {
			fmt.Printf(red("Session cookie file does not exist in %s/.adventofcode.session, please follow the in the README.md to make it\n"), userHomeDirectory)
			os.Exit(1)
		}

		// Get the session cookie
		sessionCookie, err := file.ReturnSessionCookie(userHomeDirectory)
		if err != nil {
			log.Fatal(err)
		}

		// Make the HTTP GET request and get the response body
		inputURL := fmt.Sprintf(baseInputURL, year, day)
		fmt.Printf("Downloading input for %s/%s...\n", blue(strconv.Itoa(year)), blue(strconv.Itoa(day)))
		body, err := http.ReturnBody(inputURL, sessionCookie)
		if err != nil {
			log.Fatal(err)
		}

		// Write the response body to file in the working directory
		fmt.Printf("Writing input for %s/%s...\n", blue(strconv.Itoa(year)), blue(strconv.Itoa(day)))
		err = file.WriteInput(day, year, userHomeDirectory, body)
		if err != nil {
			log.Fatal(err)
		}
	}
}
