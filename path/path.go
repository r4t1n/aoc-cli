package path

import (
	"fmt"
	"os"
	"os/user"
	"regexp"
	"strconv"
	"strings"
)

const baseInputURL = "https://adventofcode.com/%d/day/%d/input"

// ReturnDateWithError returns the year and day if found in the working directory and any potential errors
type ReturnDateWithError struct {
	Year int
	Day  int
	Err  error
}

// ReturnUserHomeDirectoryWithError returns the users home directory and any potential errors
type ReturnUserHomeDirectoryWithError struct {
	UserHomeDirectory string
	Err               error
}

func ReturnDate() ReturnDateWithError {
	// Get the working directory
	workingDirectory, err := os.Getwd()
	if err != nil {
		return ReturnDateWithError{Err: fmt.Errorf("error getting the working directory: %v", err)}
	}

	// Extract the year and day from the working directory
	directoryPattern := regexp.MustCompile("[0-9]+/[0-9]+$") // [0-9]+: year, /: child directory, [0-9]+: day, $: end of string
	directories := directoryPattern.FindString(workingDirectory)
	if len(directories) > 0 {
		yearAndDay := strings.Split(directories, "/")
		year, err := strconv.Atoi(yearAndDay[0])
		if err != nil {
			return ReturnDateWithError{Err: fmt.Errorf("error converting string to int: %v", err)}
		}
		day, err := strconv.Atoi(yearAndDay[1])
		if err != nil {
			return ReturnDateWithError{Err: fmt.Errorf("error converting string to int: %v", err)}
		}

		return ReturnDateWithError{Year: year, Day: day, Err: err}
	}

	return ReturnDateWithError{Err: err}
}

func ReturnUserHomeDirectory() ReturnUserHomeDirectoryWithError {
	// Get the Advent of Code session cookie from the users home directory
	currentUser, err := user.Current()
	if err != nil {
		return ReturnUserHomeDirectoryWithError{Err: fmt.Errorf("error getting the current user: %v", err)}
	}
	userHomeDirectory := currentUser.HomeDir

	return ReturnUserHomeDirectoryWithError{UserHomeDirectory: userHomeDirectory, Err: err}
}
