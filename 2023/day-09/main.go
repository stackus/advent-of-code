package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	. "github.com/stackus/advent-of-code"
)

//go:embed input.txt
var input string

// puzzle1 solves the level 1 puzzle
func puzzle1(input string) int {
	histories := parseInput(input)

	total := 0

	for _, history := range histories {
		sequences := [][]int{history}
		for {
			sequence := sequences[len(sequences)-1]
			differences := createDifferences(sequence)
			sequences = append(sequences, differences)
			if allSame(differences) {
				break
			}
		}
		inc := 0
		for j := len(sequences) - 1; j >= 0; j-- {
			inc = sequences[j][len(sequences[j])-1] + inc
		}
		total += inc
	}

	return total
}

// puzzle2 solves the level 2 puzzle
func puzzle2(input string) int {
	histories := parseInput(input)

	total := 0

	for _, history := range histories {
		sequences := [][]int{history}
		for {
			sequence := sequences[len(sequences)-1]
			differences := createDifferences(sequence)
			sequences = append(sequences, differences)
			if allSame(differences) {
				break
			}
		}
		inc := 0
		for j := len(sequences) - 1; j >= 0; j-- {
			inc = sequences[j][0] - inc
		}
		total += inc
	}

	return total
}

func createDifferences(sequence []int) []int {
	differences := make([]int, 0, len(sequence)-1)
	num := sequence[0]
	for _, n := range sequence[1:] {
		diff := n - num
		differences = append(differences, diff)
		num = n
	}
	return differences
}

// allSame will be true just prior to all zero sequences
func allSame(sequence []int) bool {
	for i := 1; i < len(sequence); i++ {
		if sequence[i] != sequence[0] {
			return false
		}
	}
	return true
}

// parseInput converts the input string into whatever format is needed for the puzzle
// update the return type as needed
func parseInput(input string) [][]int {
	var histories [][]int

	// match all numbers in a line (positive and negative)
	re := regexp.MustCompile(`-?\d+`)

	for _, line := range strings.Split(input, "\n") {
		matches := re.FindAllString(line, -1)
		var numbers []int
		for _, match := range matches {
			num, _ := strconv.Atoi(match)
			numbers = append(numbers, num)
		}
		histories = append(histories, numbers)

	}

	return histories
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

	started := time.Now()
	solution := 0
	if puzzle == 1 {
		solution = puzzle1(input)
	} else {
		solution = puzzle2(input)
	}
	fmt.Println("Completed in", time.Since(started))

	solutionPath := filepath.Join(GetPuzzlePath(9, 2023), fmt.Sprintf("solution-%d.txt", puzzle))
	err := WriteFile(solutionPath, []byte(fmt.Sprintf("%d", solution)), true)
	if err != nil {
		log.Fatalf("Error writing solution: %v", err)
	}
	fmt.Println("Solution:", solution)
}
