package cmd

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"
)

// Define Structs for JSON
type SpotifyResponse struct {
	Tracks struct {
		Items []struct {
			Artists []struct {
				Name string `json:"name"`
			} `json:"artists"`
			Name string `json:"name"`
			URI  string `json:"uri"`
		} `json:"items"`
	} `json:"tracks"`
}

func Search(BearerToken string, APIROOT string, searchTerm string, searchType string) {

	// Validate Search Type
	allowedTerms := []string{"album", "artist", "playlist", "track", "show", "episode", "audiobook"}
	validSearchTypes := strings.Join(allowedTerms, ", ")

	isValidSearchType := false

	for _, term := range allowedTerms {
		if searchType == term {
			isValidSearchType = true
		}
	}

	if !isValidSearchType {
		fmt.Println(searchType + " is not valid Search Type. Valid Search Types are: " + validSearchTypes)
		os.Exit(2)
	}

	// Build Final URL
	url := APIROOT + "search?q=" + url.QueryEscape(searchTerm) + "&type=" + searchType

	responseJson := SendRequest(BearerToken, url)

	var spotifyResponse SpotifyResponse

	json.Unmarshal(responseJson, &spotifyResponse)

	for _, item := range spotifyResponse.Tracks.Items {

		fmt.Println(item.Artists)

	}
	//fmt.Println(spotifyResponse)

}
