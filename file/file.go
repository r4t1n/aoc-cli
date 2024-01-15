package file

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ReturnSessionCookie(userHomeDirectory string) (sessionCookie string, err error) {
	// Read the file, trim any whitespace and convert it to string
	sessionCookieFilePath := filepath.Join(userHomeDirectory, ".adventofcode.session")
	sessionCookieByte, err := os.ReadFile(sessionCookieFilePath)
	if err != nil {
		return "", fmt.Errorf("error reading the session cookie file at %s: %v", sessionCookieFilePath, err)
	}
	sessionCookie = strings.TrimSpace(string(sessionCookieByte))

	return sessionCookie, nil
}

func CopyInput(day, year int, userHomeDirectory string) error {
	cacheFilePath := filepath.Join(userHomeDirectory, ".cache", "aoc-cli", strconv.Itoa(year), strconv.Itoa(day), "input.txt")

	cacheFilePathFile, err := os.Open(cacheFilePath)
	if err != nil {
		return fmt.Errorf("error opening file at %s: %v", cacheFilePath, err)
	}
	defer cacheFilePathFile.Close()

	destinationFile, err := os.Create("input.txt")
	if err != nil {
		return fmt.Errorf("error creating file at %s: %v", "input.txt", err)
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, cacheFilePathFile)
	if err != nil {
		return fmt.Errorf("error copying %s to %s: %v", cacheFilePath, "input.txt", err)
	}

	err = destinationFile.Sync()
	if err != nil {
		return fmt.Errorf("error syncing %s: %v", "input.txt", err)
	}

	return nil
}

func WriteInput(day, year int, userHomeDirectory, input string) error {
	cacheFilePath := filepath.Join(userHomeDirectory, ".cache", "aoc-cli", strconv.Itoa(year), strconv.Itoa(day), "input.txt")

	// Write the input to the file
	err := os.WriteFile(cacheFilePath, []byte(input), 0644)
	if err != nil {
		return fmt.Errorf("error writing to file at %s: %v", cacheFilePath, err)
	}

	err = CopyInput(day, year, userHomeDirectory)
	if err != nil {
		return err
	}

	return nil
}
