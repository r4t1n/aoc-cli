package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/r4t1n/aoc-cli/http"
	"github.com/r4t1n/aoc-cli/time"
)

const (
	baseInputURL = "https://adventofcode.com/%d/day/%d/input"
	defaultDay   = 25
)

func main() {
	// Get the Advent of Code session cookie from the users home directory
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("error getting the current user: %v", err)
	}
	sessionCookieFilePath := filepath.Join(currentUser.HomeDir, ".adventofcode.session")

	// Read the file, trim any whitespace and convert it to string
	sessionCookieByte, err := os.ReadFile(sessionCookieFilePath)
	if err != nil {
		log.Fatalf("error reading the session cookie file at %s: %v", sessionCookieFilePath, err)
	}
	sessionCookie := strings.TrimSpace(string(sessionCookieByte))

	currentTime := time.Run()
	if currentTime.Err != nil {
		log.Fatalf("error: %v", currentTime.Err)
	}

	// Check if it is December and apply the day if true, else fall back to the default day
	var inputURL string
	if currentTime.Month == "December" {
		inputURL = fmt.Sprintf(baseInputURL, currentTime.Year, currentTime.Day)
	} else {
		inputURL = fmt.Sprintf(baseInputURL, currentTime.Year, defaultDay)
	}

	// Run the HTTP module
	httpResponse := http.Run(inputURL, sessionCookie)
	if httpResponse.Err != nil {
		log.Fatalf("error: %v", httpResponse.Err)
	}

	fmt.Println(httpResponse.Body)
}
