package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
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
	tracks := Parse(inputFile)

	// Create Playlist - Prompt for Name

	// Display Songs to Add to Playlist

	//fmt.Printf("Would you like to attempt to create a playlist with the above tracklist? (yes/no) \n")

	// Get User Confirmation
	//var input string
	//fmt.Scanln(&input)

	// Normalize Input
	//input = strings.ToLower(strings.TrimSpace(input))
	//if input != "yes" && input != "y" {
	//	fmt.Println("Exiting...")
	//	os.Exit(0)
	//}

	// Create Playlist
	//createPlaylist(BearerToken, APIROOT)

	// Search
	//Hardcoding searchType for now
	searchType := "track"
	var urlList []string

	// Build URI List
	for track, artist := range tracks {
		url := Search(BearerToken, APIROOT, track, artist, searchType)
		urlList = append(urlList, url)
	}
	fmt.Println(urlList)
	FormatPlayList(tracks, urlList)

}

// Table Function
func FormatPlayList(tracks map[string]string, urlList []string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Track", "Artist", "URL"})
	i := 0
	for x, y := range tracks {

		t.AppendRow([]interface{}{x, y, urlList[i]})
		i++
	}
	t.Render()
	fmt.Printf("\n")
}
