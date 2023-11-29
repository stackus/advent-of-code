package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	. "github.com/stackus/advent-of-code"
)

//go:embed input.txt
var input string

// puzzle1 solves the level 1 puzzle
func puzzle1(input string) int {
	calories := parseInput(input)

	sort.Ints(calories)

	return calories[len(calories)-1]
}

// puzzle2 solves the level 2 puzzle
func puzzle2(input string) int {
	calories := parseInput(input)

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	total := 0
	for i := 0; i < 3; i++ {
		total += calories[i]
	}

	return total
}

// parseInput converts the input string into whatever format is needed for the puzzle
// update the return type as needed
func parseInput(input string) (calories []int) {
	index := 0
	calories = append(calories, 0)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			index++
			calories = append(calories, 0)
			continue
		}
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Error parsing input: %v", err)
		}
		calories[index] += i
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

	solutionPath := filepath.Join(GetPuzzlePath(1, 2022), fmt.Sprintf("solution-%d.txt", puzzle))
	err := WriteFile(solutionPath, []byte(fmt.Sprintf("%d", solution)))
	if err != nil {
		log.Fatalf("Error writing solution: %v", err)
	}
	fmt.Println("Solution: ", solution)
}
