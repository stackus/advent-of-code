package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"

	. "github.com/stackus/advent-of-code"
)

//go:embed input.txt
var input string

// puzzle1 solves the level 1 puzzle
func puzzle1(input string) int {
	coords := parseInput(input, false)
	// sum all coords
	total := 0
	for _, coord := range coords {
		total += coord
	}

	return total
}

// puzzle2 solves the level 2 puzzle
func puzzle2(input string) int {
	coords := parseInput(input, true)
	// sum all coords
	total := 0
	for _, coord := range coords {
		total += coord
	}

	return total
}

// parseInput converts the input string into whatever format is needed for the puzzle
// update the return type as needed
func parseInput(input string, matchNames bool) (coords []int) {
	names := []string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
	for _, line := range strings.Split(input, "\n") {
		digits := make([]int, 0)
		for i := range line {
			if line[i] >= '0' && line[i] <= '9' {
				digit, err := strconv.Atoi(string(line[i]))
				if err != nil {
					log.Fatalf("Error parsing input as digit: %v line: %s", err, line)
				}
				digits = append(digits, digit)
				continue
			}
			if matchNames {
				for n, name := range names {
					if strings.HasPrefix(line[i:], name) {
						digits = append(digits, n+1)
						break
					}
				}
			}
		}
		if len(digits) == 0 {
			log.Fatalf("Error parsing input: no digits found in line: %s", line)
		}
		coord := digits[0] * 10
		coord += digits[len(digits)-1]

		coords = append(coords, coord)
	}

	return
}

// -- leave this code alone
func main() {
	var puzzle int
	flag.IntVar(&puzzle, "puzzle", 1, "puzzle number: 1 or 2")
	flag.Parse()

	// check puzzle is valid
	if puzzle < 1 || puzzle > 2 {
		log.Fatalf("Invalid puzzle number: %d", puzzle)
	}

	// trim input
	input = strings.TrimRight(input, "\n")

	fmt.Println("Running puzzle", puzzle)

	solution := 0
	if puzzle == 1 {
		solution = puzzle1(input)
	} else {
		solution = puzzle2(input)
	}

	solutionPath := filepath.Join(GetPuzzlePath(1, 2023), fmt.Sprintf("solution-%d.txt", puzzle))
	err := WriteFile(solutionPath, []byte(fmt.Sprintf("%d", solution)), true)
	if err != nil {
		log.Fatalf("Error writing solution: %v", err)
	}
	fmt.Println("Solution: ", solution)
}
