package path

import (
	"fmt"
	"os/user"
	"path/filepath"
)

type GetSessionCookieFilePathWithError struct {
	SessionCookieFilePath string
	Err                   error
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
