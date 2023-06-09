package cmd

import (
	"encoding/json"
	"net/url"
	"strings"
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

// Response Struct
type ResponseStruct struct {
	track string
	url   string
}

func Search(BearerToken string, APIROOT string, track string, artist string, searchType string) (responseMap map[string]ResponseStruct) {
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
	url := APIROOT + "search?q=" + url.QueryEscape(track) + "&type=" + searchType + "&limit=50"

	// Parse Response
	responseJson := GetRequest(BearerToken, url)

	var spotifyResponse SpotifyResponse

	json.Unmarshal(responseJson, &spotifyResponse)

	valuesMap := make(map[string]ResponseStruct)

	// Compare JSON against Search Terms and get URL / Artist / Track

	for _, item := range spotifyResponse.Tracks.Items {

		artistresp := string(item.Artists[0].Name)

		if strings.ToLower(artist) == strings.ToLower(artistresp) {

			spotifyUrl := item.ExternalURL.Spotify
			track := item.Name

			valuesMap[artistresp] = ResponseStruct{track, spotifyUrl}

			return valuesMap
		} else if artist == "" {

			spotifyUrl := item.ExternalURL.Spotify

			valuesMap[artistresp] = ResponseStruct{track, spotifyUrl}

			return valuesMap
		}

	}

	return
}
