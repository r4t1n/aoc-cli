package file

import (
	"fmt"
	"os"
	"strings"
)

// ReadSessionCookieWithError returns the session cookie and any potential errors
type ReadSessionCookieWithError struct {
	SessionCookie string
	Err           error
}

func ReadSessionCookie(sessionCookieFilePath string) ReadSessionCookieWithError {
	// Read the file, trim any whitespace and convert it to string
	sessionCookieByte, err := os.ReadFile(sessionCookieFilePath)
	if err != nil {
		return ReadSessionCookieWithError{Err: fmt.Errorf("error reading the session cookie file at %s: %v", sessionCookieFilePath, err)}
	}
	sessionCookie := strings.TrimSpace(string(sessionCookieByte))

	return ReadSessionCookieWithError{SessionCookie: sessionCookie, Err: nil}
}

func WriteInput(input string) error {
	filePath := "input.txt"

	// Write the input to the file
	err := os.WriteFile(filePath, []byte(input), 0666)
	if err != nil {
		return fmt.Errorf("error writing to the file: %v", err)
	}

	return nil
}
