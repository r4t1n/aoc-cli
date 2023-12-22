package http

import (
	"fmt"
	"io"
	"net/http"
)

// ReturnWithError returns any values and any potential errors
type ReturnWithError struct {
	Body string
	Err  error
}

func Get(url, sessionCookie string) ReturnWithError {
	// Create the HTTP client and create an HTTP GET request with the provided URL
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ReturnWithError{Err: fmt.Errorf("error creating the HTTP GET request: %v", err)}
	}

	// Add the session cookie to the request
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	// Make the request and close the response body
	resp, err := client.Do(req)
	if err != nil {
		return ReturnWithError{Err: fmt.Errorf("error making the HTTP GET request: %v", err)}
	}
	defer resp.Body.Close()

	// Check for a successful response
	if resp.StatusCode != http.StatusOK {
		return ReturnWithError{Err: fmt.Errorf("HTTP GET request failed with status code %d", resp.StatusCode)}
	}

	// Read the response body and convert it to string
	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return ReturnWithError{Err: fmt.Errorf("error reading the response body: %v", err)}
	}
	body := string(bodyByte)

	return ReturnWithError{Body: body, Err: nil}
}
