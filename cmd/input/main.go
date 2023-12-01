package main

import (
	"fmt"
	"log"
	"os"

	. "github.com/stackus/advent-of-code"
)

func main() {
	day, year := ParseFlags()
	adventOfCodePath := MakeDir(day, year)

	// Get the puzzle input for the Advent of Code website for the given day and year
	input := getInput(day, year)

	// write the puzzle input to a file
	inputPath := fmt.Sprintf("%s%sinput.txt", adventOfCodePath, string(os.PathSeparator))
	err := WriteFile(inputPath, []byte(input), true)
	if err != nil {
		log.Fatalf("Error writing puzzle input: %v", err)
	}

	fmt.Println("Puzzle input written for day", day, "and year", year)
}

func getInput(day, year int) string {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	body, err := DoGet(url)
	if err != nil {
		log.Fatalf("Error getting puzzle input: %v", err)
	}

	return string(body)
}
