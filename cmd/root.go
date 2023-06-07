package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// Global Variables

func Root() {

	var APIROOT = "https://api.spotify.com/v1/"

	// Validate Arguments
	if len(os.Args) != 2 {
		fmt.Println("Please provide one file name as an argument.")
		return
	}

	inputFile := os.Args[1]

	// Validate File

	// Open the file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err.Error())
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Process each line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err.Error())
	}

	// Authenticate
	BearerToken := GetCreds()

	// Search
	searchType := "artist"
	searchTerm := "metallica"
	Search(BearerToken, APIROOT, searchTerm, searchType)

}

func useRegex(s string) bool {
	re := regexp.MustCompile("(?i)[A-Za-z]+.*-.*[A-Za-z]+")
	return re.MatchString(s)
}
