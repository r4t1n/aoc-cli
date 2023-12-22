package http

import (
	"fmt"
	"io"
	"net/http"
)

// ResponseWithError represents the body of an HTTP GET request and any potential error
type ResponseWithError struct {
	Body string
	Err  error
}

func Run(url, sessionCookie string) ResponseWithError {
	// Create the HTTP client and create an HTTP GET request with the provided URL
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseWithError{Err: fmt.Errorf("error creating the HTTP GET request: %v", err)}
	}

	// Add the session cookie to the request
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	// Make the request and close the response body
	resp, err := client.Do(req)
	if err != nil {
		return ResponseWithError{Err: fmt.Errorf("error making the HTTP GET request: %v", err)}
	}
	defer resp.Body.Close()

	// Check for a successful response
	if resp.StatusCode != http.StatusOK {
		return ResponseWithError{Err: fmt.Errorf("HTTP GET request failed with status code %d", resp.StatusCode)}
	}

	// Read the response body and convert it to string
	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponseWithError{Err: fmt.Errorf("error reading the response body: %v", err)}
	}
	body := string(bodyByte)

	return ResponseWithError{Body: body, Err: nil}
}
