package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"math"
	"path/filepath"
	"regexp"
	"strings"

	. "github.com/stackus/advent-of-code"
)

//go:embed input.txt
var input string

// puzzle1 solves the level 1 puzzle
func puzzle1(input string) int {
	parsed := parseInput(input)
	total := 0

	for _, line := range parsed {
		s := strings.Split(line, ":")
		numGroups := strings.Split(s[1], " | ")
		re := regexp.MustCompile(`\d+`)
		nums := re.FindAllString(numGroups[0], -1)
		winningNums := re.FindAllString(numGroups[1], -1)
		matches := 0
		for _, num := range nums {
			for _, winningNum := range winningNums {
				if num == winningNum {
					matches++
				}
			}
		}
		if matches == 0 {
			continue
		}
		total += int(math.Pow(2, float64(matches-1)))
	}

	return total
}

// puzzle2 solves the level 2 puzzle
func puzzle2(input string) int {
	parsed := parseInput(input)
	total := 0

	cards := make([]int, len(parsed))
	for lineNum, line := range parsed {
		// initial card copy
		cards[lineNum]++

		s := strings.Split(line, ":")
		numGroups := strings.Split(s[1], " | ")
		re := regexp.MustCompile(`\d+`)
		nums := re.FindAllString(numGroups[0], -1)
		winningNums := re.FindAllString(numGroups[1], -1)
		matches := 0
		for _, num := range nums {
			for _, winningNum := range winningNums {
				if num == winningNum {
					matches++
					if lineNum+matches >= len(cards) {
						break
					}
					cards[lineNum+matches] += cards[lineNum]
				}
			}
		}
	}

	return total
}

// parseInput converts the input string into whatever format is needed for the puzzle
// update the return type as needed
func parseInput(input string) (lines []string) {
	for _, line := range strings.Split(input, "\n") {
		lines = append(lines, line)
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

	solutionPath := filepath.Join(GetPuzzlePath(4, 2023), fmt.Sprintf("solution-%d.txt", puzzle))
	err := WriteFile(solutionPath, []byte(fmt.Sprintf("%d", solution)), true)
	if err != nil {
		log.Fatalf("Error writing solution: %v", err)
	}
	fmt.Println("Solution: ", solution)
}
