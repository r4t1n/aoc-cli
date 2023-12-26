package path

import (
	"fmt"
	"os"
	"os/user"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/r4t1n/aoc-cli/time"
)

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

		// Return the date from the current date
		timeDate := time.ReturnDate()
		if timeDate.Err != nil {
			return ReturnDateWithError{Err: timeDate.Err}
		}

		// Check if date from path is invalid
		if year < 2015 || year > timeDate.Year || day < 1 || day > 25 {
			color.Yellow("Date from path is invalid, falling back to current date")
		} else {
			return ReturnDateWithError{Year: year, Day: day, Err: err}
		}

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
