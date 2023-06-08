package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
		line := scanner.Text()

		// Validate the input format
		if !validateString(line) {
			//fmt.Printf("%s is not a valid input - skipping \n", line)
			break
		}

		// Map strings
		lineSplit := strings.Split(line, "-")
		songs[lineSplit[0]] = lineSplit[1]
	}

	return songList

}

func validateString(s string) bool {
	re := regexp.MustCompile("(?i)[A-Za-z]+.*-.*[A-Za-z]+")
	return re.MatchString(s)
}
