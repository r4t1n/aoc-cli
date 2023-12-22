package time

import (
	"fmt"
	"time"
)

// ReturnWithError returns any values and any potential errors
type ReturnWithError struct {
	Year  int
	Month string
	Day   int
	Err   error
}

func Get() ReturnWithError {
	// Set the timezone to EST/UTC-5 as Advent of Code uses it
	est, err := time.LoadLocation("America/New_York")
	if err != nil {
		return ReturnWithError{Err: fmt.Errorf("error loading the EST/UTC-5 timezone: %v", err)}
	}

	// Get the current year, month and day
	currentTimeAOC := time.Now().In(est)
	year := currentTimeAOC.Year()
	month := currentTimeAOC.Month().String()
	day := currentTimeAOC.Day()

	return ReturnWithError{Year: year, Month: month, Day: day, Err: nil}
}
