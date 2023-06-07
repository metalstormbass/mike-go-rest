package cmd

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

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

	SendRequest(BearerToken, url)
}
