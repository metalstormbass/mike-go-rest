package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

// Global Variables
// Response Struct
type tableStruct struct {
	track string
	url   string
}

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

	// Build URL List
	valuesMap := make(map[string]tableStruct)
	for track, artist := range tracks {
		responseMap := Search(BearerToken, APIROOT, track, artist, searchType)

		for key, value := range responseMap {
			valuesMap[key] = tableStruct{value.track, value.url}
		}
	}
	//fmt.Println(urlList)
	FormatPlayList(valuesMap)

}

// Table Function
func FormatPlayList(valuesMap map[string]tableStruct) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Track", "Artist", "URL"})
	i := 0
	for x, y := range valuesMap {

		t.AppendRow([]interface{}{x, y.track, y.url})
		i++
	}
	t.Render()
	fmt.Printf("\n")
}
