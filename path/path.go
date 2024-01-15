package path

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
)

func CheckForCachedInput(year, day int, userHomeDirectory string) (cachedInputExists bool, err error) {
	dateFolderPath := filepath.Join(userHomeDirectory, ".cache", "aoc-cli", strconv.Itoa(year), strconv.Itoa(day))
	inputFilePath := filepath.Join(dateFolderPath, "input.txt")

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

func ReturnUserHomeDirectory() (userHomeDirectory string, err error) {
	// Get the Advent of Code session cookie from the users home directory
	currentUser, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("error getting the current user: %v", err)
	}
	userHomeDirectory = currentUser.HomeDir

	return userHomeDirectory, err
}
