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

func ReturnDate() (day, year int, err error) {
	// Return the date from the path
	pathDay, pathYear, err := ReturnDateFromPath()
	if err != nil {
		return 0, 0, err
	}

	// Return the date from the current date
	timeDay, timeMonth, timeYear, err := ReturnCurrentDateEST()
	if err != nil {
		return 0, 0, err
	}

	// Check if date from path is invalid
	if pathDay < 1 || pathDay > 25 || pathYear < 2015 || pathYear > timeYear || (pathYear == timeYear && timeMonth != 12) {
		// Check if date from time is invalid
		if timeMonth != 12 {
			timeDay = defaultDay
			timeYear = timeYear - 1
		}
		if timeDay > 25 {
			timeDay = defaultDay
		}

		return timeDay, timeYear, nil
	} else {
		return pathDay, pathYear, nil
	}
}

func ReturnDateFromPath() (day, year int, err error) {
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

	return pathDay, pathYear, err
}

func ReturnCurrentDateEST() (year, month, day int, err error) {
	// Set the timezone to EST/UTC-5 as Advent of Code uses it
	est, err := time.LoadLocation("America/New_York")
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error loading the EST/UTC-5 timezone: %v", err)
	}

	// Get the current day, month and year
	currentTimeEST := time.Now().In(est)
	day = currentTimeEST.Day()
	month = int(currentTimeEST.Month())
	year = currentTimeEST.Year()

	return day, month, year, err
}
