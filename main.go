package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"github.com/r4t1n/aoc-cli/http"
)

const (
	baseInputURL = "https://adventofcode.com/%d/day/%d/input"
	defaultDay   = 25
)

func main() {
	// Get the Advent of Code session cookie from the users home directory
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("Error getting the current user: %v", err)
	}
	sessionCookieFilePath := filepath.Join(currentUser.HomeDir, ".adventofcode.session")

	// Read the file, trim any whitespace and convert it to string
	sessionCookieByte, err := os.ReadFile(sessionCookieFilePath)
	if err != nil {
		log.Fatalf("Error reading the session cookie file at %s: %v", sessionCookieFilePath, err)
	}
	sessionCookie := strings.TrimSpace(string(sessionCookieByte))

	// Set the timezone to EST/UTC-5 as Advent of Code uses it
	est, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalf("Error loading the EST/UTC-5 timezone: %v", err)
	}

	// Get the current year, month and date
	currentTimeEST := time.Now().In(est)
	year := currentTimeEST.Year()
	month := currentTimeEST.Month()
	day := currentTimeEST.Day()

	// Check if it is December and apply the day if true, else fall back to the default day
	var inputURL string
	if month == time.December {
		inputURL = fmt.Sprintf(baseInputURL, year, day)
	} else {
		inputURL = fmt.Sprintf(baseInputURL, year, defaultDay)
	}

	// Run the HTTP module
	httpResponse := http.Run(inputURL, sessionCookie)
	if httpResponse.Err != nil {
		log.Fatalf("Error: %v", httpResponse.Err)
	}

	fmt.Println(httpResponse.Body)
}
