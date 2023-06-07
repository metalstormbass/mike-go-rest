package cmd

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var AccessToken string

func GetCreds() (accessToken string) {
	// Spotify Client ID + Client Secret
	CLIENT_ID := os.Getenv("CLIENT_ID")
	CLIENT_SECRET := os.Getenv("CLIENT_SECRET")

	apiUrl := "https://accounts.spotify.com/api/token"
	AccessToken = getCreds(apiUrl, CLIENT_ID, CLIENT_SECRET)

	return AccessToken
}

func getCreds(apiUrl string, CLIENT_ID string, CLIENT_SECRET string) (accessToken string) {

	// Define Struct for JSON response

	type Response struct {
		AccessToken string `json:"access_token"`
	}

	// Create an HTTP client
	client := &http.Client{}

	// Create Token String
	data := CLIENT_ID + ":" + CLIENT_SECRET
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))

	// Add Required Form Data
	formData := url.Values{}
	formData.Set("grant_type", "client_credentials")

	// Create an HTTP request with custom headers
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Fatal("Error creating HTTP request:", err)
		return
	}
	req.Header.Add("Authorization", "Basic "+string(sEnc))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
	var bodyObject Response
	json.Unmarshal(body, &bodyObject)

	// Print the response body
	return string(bodyObject.AccessToken)
}
