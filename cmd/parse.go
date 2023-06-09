package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Parse(inputFile string) (songList map[string]string) {

	// Open and process file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err.Error())
		os.Exit(2)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Define Map Variable to Store Songs

	songs := make(map[string]string)

	for scanner.Scan() {

		// Extract each line
		line := scanner.Text()

		// Trim line
		item := strings.TrimSpace(line)
		lineSplit := strings.SplitN(item, "-", 2)

		if len(lineSplit) == 1 {
			lineSplit = append(lineSplit, "")
		}
		// Assemble List
		if lineSplit[0] != "" {
			songs[strings.TrimSpace(lineSplit[0])] = strings.TrimSpace(lineSplit[1])
		}
	}
	return songs

}
