package time

import (
	"fmt"
	"time"
)

// ReturnWithError returns the time and any potential error
type ReturnWithError struct {
	Year  int
	Month string
	Day   int
	Err   error
}

func Run() ReturnWithError {
	// Set the timezone to EST/UTC-5 as Advent of Code uses it
	est, err := time.LoadLocation("America/New_York")
	if err != nil {
		return ReturnWithError{Err: fmt.Errorf("error loading the EST/UTC-5 timezone: %v", err)}
	}

	// Get the current year, month and date
	currentTimeEST := time.Now().In(est)
	year := currentTimeEST.Year()
	month := currentTimeEST.Month().String()
	day := currentTimeEST.Day()

	return ReturnWithError{Year: year, Month: month, Day: day, Err: nil}
}
