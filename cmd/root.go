package cmd

import (
	"fmt"
	"os"
)

// Global Variables

func Root() {

	var APIROOT = "https://api.spotify.com/v1/"

	// Validate Input
	if len(os.Args) != 2 {
		fmt.Println("Please provide one file name as an argument.")
		return
	}

	// Set Variable with path of text file
	inputFile := os.Args[1]

	// Authenticate
	BearerToken := GetCreds()

	// Parse FIle
	//tracks :=
	Parse(inputFile)

	// Search
	searchType := "track"
	searchTerm := "ride the lightning"
	Search(BearerToken, APIROOT, searchTerm, searchType)

}
