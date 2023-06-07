package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func SendRequest(BearerToken string, url string) {
	// Create an HTTP client
	client := &http.Client{}

	// Create an HTTP request with custom headers
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating HTTP request:", err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+BearerToken)

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}

	// Read the response body

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response body:", err)
		return
	}
	// Print the response body as raw text
	fmt.Println(string(body))

}
