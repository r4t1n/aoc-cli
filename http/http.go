package http

import (
	"fmt"
	"io"
	"net/http"
)

func Run(url, sessionCookie string) error {
	// Create the HTTP client and create a GET request with the provided URL
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Add the session cookie to the request
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	// Make the request and close the response body
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))
	return nil
}
