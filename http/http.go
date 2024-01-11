package http

import (
	"fmt"
	"io"
	"net/http"
)

// ReturnBodyWithError returns the HTTP body and any potential errors
type ReturnBodyWithError struct {
	Body string
	Err  error
}

func ReturnBody(url, sessionCookie string) ReturnBodyWithError {
	// Create the HTTP client and create an HTTP GET request with the provided URL
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ReturnBodyWithError{Err: fmt.Errorf("error creating HTTP GET request for %s: %v", url, err)}
	}

	// Add the session cookie and User Agent to the request
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})
	req.Header.Set("User-Agent", "https://github.com/r4t1n/aoc-cli by ratin")

	// Make the request and close the response body
	resp, err := client.Do(req)
	if err != nil {
		return ReturnBodyWithError{Err: fmt.Errorf("error making HTTP GET request for %s: %v", url, err)}
	}
	defer resp.Body.Close()

	// Check for a successful response
	if resp.StatusCode != http.StatusOK {
		return ReturnBodyWithError{Err: fmt.Errorf("HTTP GET request failed for %s: status code %d", url, resp.StatusCode)}
	}

	// Read the response body and convert it to string
	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return ReturnBodyWithError{Err: fmt.Errorf("error reading the response body: %v", err)}
	}
	body := string(bodyByte)

	return ReturnBodyWithError{Body: body, Err: nil}
}
