package path

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const baseInputURL = "https://adventofcode.com/%d/day/%d/input"

// ReturnInputURLWithError returns the year and day if found in the working directory and any potential errors
type ReturnInputURLWithError struct {
	InputURL string
	Err      error
}

// GetSessionCookieFilePathWithError returns session cookie file path and any potential errors
type GetSessionCookieFilePathWithError struct {
	SessionCookieFilePath string
	Err                   error
}

func ReturnInputURL() ReturnInputURLWithError {
	// Get the working directory
	workingDirectory, err := os.Getwd()
	if err != nil {
		return ReturnInputURLWithError{Err: fmt.Errorf("error getting the working directory: %v", err)}
	}

	// Extract the year and day from the working directory
	directoryPattern := regexp.MustCompile("[0-9]+/[0-9]+$") // [0-9]+: year, /: child directory, [0-9]+: day, $: end of string
	directories := directoryPattern.FindString(workingDirectory)
	if len(directories) > 0 {
		yearAndDay := strings.Split(directories, "/")
		year, err := strconv.Atoi(yearAndDay[0])
		if err != nil {
			return ReturnInputURLWithError{Err: fmt.Errorf("error converting string to int: %v", err)}
		}
		day, err := strconv.Atoi(yearAndDay[1])
		if err != nil {
			return ReturnInputURLWithError{Err: fmt.Errorf("error converting string to int: %v", err)}
		}

		inputURL := fmt.Sprintf(baseInputURL, year, day)

		return ReturnInputURLWithError{InputURL: inputURL, Err: err}
	}

	return ReturnInputURLWithError{Err: err}
}

func GetSessionCookieFilePath() GetSessionCookieFilePathWithError {
	// Get the Advent of Code session cookie from the users home directory
	currentUser, err := user.Current()
	if err != nil {
		return GetSessionCookieFilePathWithError{Err: fmt.Errorf("error getting the current user: %v", err)}
	}
	sessionCookieFilePath := filepath.Join(currentUser.HomeDir, ".adventofcode.session")

	return GetSessionCookieFilePathWithError{SessionCookieFilePath: sessionCookieFilePath, Err: err}
}
