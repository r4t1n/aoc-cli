package date

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	defaultDay = 1
)

func ReturnDate() (year, day int, err error) {
	// Return the date from the path
	pathYear, pathDay, err := ReturnDateFromPath()
	if err != nil {
		return 0, 0, err
	}

	// Return the date from the current date
	timeYear, timeMonth, timeDay, err := ReturnCurrentDateEST()
	if err != nil {
		return 0, 0, err
	}

	// Check if date from path is invalid
	if pathYear < 2015 || pathYear > timeYear || pathDay < 1 || pathDay > 25 {
		// Check if date from time is invalid
		if timeMonth != 12 {
			timeYear = timeYear - 1
			timeDay = defaultDay
		}
		if timeDay > 25 {
			day = defaultDay
		}

		return timeYear, timeDay, nil
	} else {
		return pathYear, pathDay, nil
	}
}

func ReturnDateFromPath() (year, day int, err error) {
	// Get the working directory
	workingDirectory, err := os.Getwd()
	if err != nil {
		return 0, 0, fmt.Errorf("error getting the working directory: %v", err)
	}

	// Extract the year and day from the working directory
	directoryPattern := regexp.MustCompile("[0-9]+/[0-9]+") // [0-9]+: year, /: child directory, [0-9]+: day
	directories := directoryPattern.FindString(workingDirectory)
	var pathYear, pathDay int
	if len(directories) > 0 {
		pathYearAndDay := strings.Split(directories, "/")

		pathYear, err = strconv.Atoi(pathYearAndDay[0])
		if err != nil {
			return 0, 0, fmt.Errorf("error converting string to int: %v", err)
		}
		pathDay, err = strconv.Atoi(pathYearAndDay[1])
		if err != nil {
			return 0, 0, fmt.Errorf("error converting string to int: %v", err)
		}
	}

	return pathYear, pathDay, err
}

func ReturnCurrentDateEST() (year, month, day int, err error) {
	// Set the timezone to EST/UTC-5 as Advent of Code uses it
	est, err := time.LoadLocation("America/New_York")
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error loading the EST/UTC-5 timezone: %v", err)
	}

	// Get the current year, month and day
	currentTimeEST := time.Now().In(est)
	year = currentTimeEST.Year()
	month = int(currentTimeEST.Month())
	day = currentTimeEST.Day()

	return year, month, day, err
}
