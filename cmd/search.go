package cmd

import (
	"encoding/json"
	"net/url"
)

// Define Structs for JSON
type SpotifyResponse struct {
	Tracks struct {
		Items []struct {
			Artists []struct {
				Name string `json:"name"`
			} `json:"artists"`
			Name        string `json:"name"`
			URI         string `json:"uri"`
			ExternalURL struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
		} `json:"items"`
	} `json:"tracks"`
}

func Search(BearerToken string, APIROOT string, track string, artist string, searchType string) (uri string) {
	/*
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
	*/
	// Build Final URL
	url := APIROOT + "search?q=" + url.QueryEscape(track) + "&type=" + searchType

	// Parse Response
	responseJson := GetRequest(BearerToken, url)

	var spotifyResponse SpotifyResponse

	json.Unmarshal(responseJson, &spotifyResponse)

	// Compare JSON against Search Terms and assemble list of URI

	for _, item := range spotifyResponse.Tracks.Items {

		artistresp := string(item.Artists[0].Name)
		if artist == artistresp || artist != "" {
			spotifyUrl := item.ExternalURL.Spotify

			return spotifyUrl
		} else {
			spotifyUrl := item.ExternalURL.Spotify
			return spotifyUrl
		}

	}

	return
}
