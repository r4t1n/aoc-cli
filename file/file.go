package file

import (
	"fmt"
	"os"
	"strings"
)

// ReturnSessionCookieWithError returns the session cookie and any potential errors
type ReturnSessionCookieWithError struct {
	SessionCookie string
	Err           error
}

func ReturnSessionCookie(sessionCookieFilePath string) ReturnSessionCookieWithError {
	// Read the file, trim any whitespace and convert it to string
	sessionCookieByte, err := os.ReadFile(sessionCookieFilePath)
	if err != nil {
		return ReturnSessionCookieWithError{Err: fmt.Errorf("error reading the session cookie file at %s: %v", sessionCookieFilePath, err)}
	}
	sessionCookie := strings.TrimSpace(string(sessionCookieByte))

	return ReturnSessionCookieWithError{SessionCookie: sessionCookie, Err: nil}
}

func WriteInput(input string) error {
	filePath := "input.txt"
	input = strings.TrimSpace(input)

	// Write the input to the file
	err := os.WriteFile(filePath, []byte(input), 0666)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}
