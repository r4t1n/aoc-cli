package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/r4t1n/aoc-cli/http"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	sessionCookieFilePath := filepath.Join(currentUser.HomeDir, ".adventofcode.session")

	sessionCookieByte, err := os.ReadFile(sessionCookieFilePath)
	if err != nil {
		log.Fatal(err)
	}
	sessionCookie := strings.TrimSpace(string(sessionCookieByte))

	url := "https://adventofcode.com/2015/day/1/input"
	err = http.Run(url, sessionCookie)
	if err != nil {
		log.Fatal(err)
	}
}
