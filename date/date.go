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

func ReturnCurrentDateEST() (year, month, day uint, err error) {
	// Set the timezone to EST/UTC-5 as Advent of Code uses it
	est, err := time.LoadLocation("America/New_York")
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error loading the EST/UTC-5 timezone: %v", err)
	}

	// Get the current day, month and year
	currentTimeEST := time.Now().In(est)
	year = uint(currentTimeEST.Year())
	month = uint(currentTimeEST.Month())
	day = uint(currentTimeEST.Day())

	return year, month, day, err
}

func ReturnDate() (year, day uint, err error) {
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
	if pathYear < 2015 || pathYear > timeYear || pathDay < 1 || pathDay > 25 || (pathYear == timeYear && timeMonth != 12) {
		// Check if date from time is invalid
		if timeMonth != 12 {
			timeYear = timeYear - 1
			timeDay = defaultDay
		}
		if timeDay > 25 {
			timeDay = defaultDay
		}
		return timeYear, timeDay, nil
	} else {
		return pathYear, pathDay, nil
	}
}

func ReturnDateFromPath() (year, day uint, err error) {
	// Get the working directory
	workingDirectory, err := os.Getwd()
	if err != nil {
		return 0, 0, fmt.Errorf("error getting the working directory: %v", err)
	}

	// Extract the year and day from the working directory
	directoryPattern := regexp.MustCompile("[0-9]+/[0-9]+") // [0-9]+: year, /: child directory, [0-9]+: day
	directories := directoryPattern.FindString(workingDirectory)
	var pathDay, pathYear int
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

	return uint(pathYear), uint(pathDay), err
}
