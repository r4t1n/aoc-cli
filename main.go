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
	baseInputURL string = "https://adventofcode.com/%d/day/%d/input"
)

var (
	blue            = color.New(color.FgBlue).SprintFunc()
	red             = color.New(color.FgRed).SprintFunc()
	argYear, argDay uint
)

func init() {
	flag.UintVar(&argYear, "year", 0, "The year used for the date")
	flag.UintVar(&argDay, "day", 0, "The day used for the date")
}

func main() {
	flag.Parse()

	// Get the date either from the path or the time
	year, day, err := date.ReturnDate()
	if err != nil {
		log.Fatal(err)
	}

	timeYear, timeMonth, _, err := date.ReturnCurrentDateEST()
	if err != nil {
		log.Fatal(err)
	}

	if argYear != 0 || argDay != 0 {
		if argYear != 0 {
			if argYear < 2015 || argYear > timeYear || (argYear == timeYear && timeMonth != 12) {
				fmt.Println(red("Year provided is invalid!"))
				os.Exit(1)
			} else {
				year = argYear
			}
		}
		if argDay != 0 {
			if argDay < 1 || argDay > 25 {
				fmt.Println(red("Day provided is invalid!"))
				os.Exit(1)
			} else {
				day = argDay
			}
		}
	}

	// Get the users home directory
	userHomeDirectory, err := file.ReturnUserHomeDirectory()
	if err != nil {
		log.Fatal(err)
	}

	GetInput(year, day, userHomeDirectory)
}

func GetInput(year, day uint, userHomeDirectory string) {
	// Check if the cached input exists
	cachedInputExists, err := file.CheckForCachedInput(year, day, userHomeDirectory)
	if err != nil {
		log.Fatal(err)
	}

	if cachedInputExists {
		// Copy the cached input
		fmt.Printf("Copying input for %s/%s...\n", blue(strconv.Itoa(int(year))), blue(strconv.Itoa(int(day))))
		err = file.CopyInput(year, day, userHomeDirectory)
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
			fmt.Printf("%s in \"%s/.adventofcode.session\"!\nPlease follow the README.md to make it\n", red("Session cookie file does not exist"), userHomeDirectory)
			os.Exit(1)
		}

		// Get the session cookie
		sessionCookie, err := file.ReturnSessionCookie(userHomeDirectory)
		if err != nil {
			log.Fatal(err)
		}

		// Make the HTTP GET request and get the response body
		inputURL := fmt.Sprintf(baseInputURL, year, day)
		fmt.Printf("Downloading input for %s/%s...\n", blue(strconv.Itoa(int(year))), blue(strconv.Itoa(int(day))))
		body, err := http.ReturnBody(inputURL, sessionCookie)
		if err != nil {
			log.Fatal(err)
		}

		// Write the response body to file in the working directory
		fmt.Printf("Writing input for %s/%s...\n", blue(strconv.Itoa(int(year))), blue(strconv.Itoa(int(day))))
		err = file.WriteInput(year, day, userHomeDirectory, body)
		if err != nil {
			log.Fatal(err)
		}
	}
}
