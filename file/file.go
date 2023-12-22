package file

import (
	"fmt"
	"os"
	"strings"
)

// ReturnWithError returns any values and any potential errors
type ReturnWithError struct {
	SessionCookie string
	Err           error
}

func ReadSessionCookie(sessionCookieFilePath string) ReturnWithError {
	// Read the file, trim any whitespace and convert it to string
	sessionCookieByte, err := os.ReadFile(sessionCookieFilePath)
	if err != nil {
		return ReturnWithError{Err: fmt.Errorf("error reading the session cookie file at %s: %v", sessionCookieFilePath, err)}
	}
	sessionCookie := strings.TrimSpace(string(sessionCookieByte))

	return ReturnWithError{SessionCookie: sessionCookie, Err: nil}
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
