package file

import (
	"fmt"
	"io"
	"os"
	"os/user"
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
	cacheFilePath := filepath.Join(userHomeDirectory, ".cache", "aoc-cli", strconv.Itoa(year), strconv.Itoa(day), "input")
	filePath := "input"

	cacheFilePathFile, err := os.Open(cacheFilePath)
	if err != nil {
		return fmt.Errorf("error opening file at %s: %v", cacheFilePath, err)
	}
	defer cacheFilePathFile.Close()

	destinationFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file at %s: %v", filePath, err)
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, cacheFilePathFile)
	if err != nil {
		return fmt.Errorf("error copying %s to %s: %v", cacheFilePath, filePath, err)
	}

	err = destinationFile.Sync()
	if err != nil {
		return fmt.Errorf("error syncing %s: %v", filePath, err)
	}

	return nil
}

func WriteInput(day, year int, userHomeDirectory, input string) error {
	cacheFilePath := filepath.Join(userHomeDirectory, ".cache", "aoc-cli", strconv.Itoa(year), strconv.Itoa(day), "input")

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

func CheckForCachedInput(year, day int, userHomeDirectory string) (cachedInputExists bool, err error) {
	dateFolderPath := filepath.Join(userHomeDirectory, ".cache", "aoc-cli", strconv.Itoa(year), strconv.Itoa(day))
	inputFilePath := filepath.Join(dateFolderPath, "input")

	// Check if the date folder exists, if not create it
	if _, err = os.Stat(dateFolderPath); os.IsNotExist(err) {
		err = os.MkdirAll(dateFolderPath, os.ModePerm)
		return false, err
	}

	// Check if the input file exists
	if _, err = os.Stat(inputFilePath); !os.IsNotExist(err) { // os.IsExist did not seem to work for some reason, so just use os.IsNotExist and flip the bool
		return true, err
	}

	return false, nil
}

func CheckForSessionCookie(userHomeDirectory string) (sessionCookieExists bool, err error) {
	sessionCookieFilePath := filepath.Join(userHomeDirectory, ".adventofcode.session")

	// Check if the session cookie file exists
	if _, err = os.Stat(sessionCookieFilePath); !os.IsNotExist(err) {
		return true, err
	}

	return false, nil
}

func ReturnUserHomeDirectory() (userHomeDirectory string, err error) {
	// Get the Advent of Code session cookie from the users home directory
	currentUser, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("error getting the current user: %v", err)
	}
	userHomeDirectory = currentUser.HomeDir

	return userHomeDirectory, err
}
