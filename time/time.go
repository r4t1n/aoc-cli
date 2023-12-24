package time

import (
	"fmt"
	"time"
)

const (
	baseInputURL = "https://adventofcode.com/%d/day/%d/input"
	defaultDay   = 1
)

// ReturnInputURLWithError returns the input URL from the date
type ReturnInputURLWithError struct {
	InputURL string
	Err      error
}

func ReturnInputURL() ReturnInputURLWithError {
	// Set the timezone to EST/UTC-5 as Advent of Code uses it
	est, err := time.LoadLocation("America/New_York")
	if err != nil {
		return ReturnInputURLWithError{Err: fmt.Errorf("error loading the EST/UTC-5 timezone: %v", err)}
	}

	// Get the current year, month and day
	currentTimeAOC := time.Now().In(est)
	year := currentTimeAOC.Year()
	month := currentTimeAOC.Month().String()
	day := currentTimeAOC.Day()

	// Check if it is December and apply the day if true, else fall back to the default day
	var inputURL string
	if month == "December" {
		inputURL = fmt.Sprintf(baseInputURL, year, day)
	} else {
		inputURL = fmt.Sprintf(baseInputURL, year, defaultDay)
	}

	return ReturnInputURLWithError{InputURL: inputURL, Err: nil}
}
