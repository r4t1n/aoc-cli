package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"

	"github.com/fatih/color"
	"github.com/r4t1n/aoc-cli/file"
	"github.com/r4t1n/aoc-cli/http"
	"github.com/r4t1n/aoc-cli/path"
	"github.com/r4t1n/aoc-cli/time"
)

const (
	baseInputURL = "https://adventofcode.com/%d/day/%d/input"
	defaultDay   = 1
)

var (
	blue = color.New(color.FgBlue).SprintFunc()
)

func main() {
	// Get the users home directory to get the session cookie file path
	userHomeDirectory := path.ReturnUserHomeDirectory()
	if userHomeDirectory.Err != nil {
		log.Fatal(userHomeDirectory.Err)
	}
	sessionCookieFilePath := filepath.Join(userHomeDirectory.UserHomeDirectory, ".adventofcode.session")

	// Get the session cookie
	sessionCookie := file.ReturnSessionCookie(sessionCookieFilePath)
	if sessionCookie.Err != nil {
		log.Fatal(sessionCookie.Err)
	}

	// Try to get the date from the working directory
	pathDate := path.ReturnDate()
	if pathDate.Err != nil {
		log.Fatal(pathDate.Err)
	}

	// Set the year and day for the input URL, either from the path or the current date
	var year int
	var day int
	if pathDate.Year != 0 && pathDate.Day != 0 {
		year = pathDate.Year
		day = pathDate.Day
	} else {
		// Return the date from the current date
		timeDate := time.ReturnDate()
		if timeDate.Err != nil {
			log.Fatal(timeDate.Err)
		}

		year = timeDate.Year
		day = timeDate.Day

		// Check if it is not December or if the day is over 25 (last Advent of Code puzzle)
		if timeDate.Month != "December" {
			year = timeDate.Year - 1
			day = defaultDay
		} else if timeDate.Day > 25 {
			day = defaultDay
		}
	}
	inputURL := fmt.Sprintf(baseInputURL, year, day)
	fmt.Printf("Downloading input for %s/%s...\n", blue(strconv.Itoa(year)), blue(strconv.Itoa(day)))

	// Make the HTTP GET request and get the response body
	httpResponse := http.ReturnBody(inputURL, sessionCookie.SessionCookie)
	if httpResponse.Err != nil {
		log.Fatal(httpResponse.Err)
	}

	// Write the response body to file
	err := file.WriteInput(httpResponse.Body)
	if err != nil {
		log.Fatal(err)
	}
}
