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
	coords := parseInput(input)
	// sum all coords
	total := 0
	for _, coord := range coords {
		total += coord
	}

	return total
}

// puzzle2 solves the level 2 puzzle
func puzzle2(input string) int {
	coords := parseInput(input)
	// sum all coords
	total := 0
	for _, coord := range coords {
		total += coord
	}

	return total
}

// parseInput converts the input string into whatever format is needed for the puzzle
// update the return type as needed
func parseInput(input string) (coords []int) {
	getDigit := func(rs []rune) rune {
		num := []rune{}
		for _, r := range rs {
			if r >= 'a' && r <= 'z' {
				num = append(num, r)
				switch {
				case strings.HasSuffix(string(num), "one"):
					return '1'
				case strings.HasSuffix(string(num), "two"):
					return '2'
				case strings.HasSuffix(string(num), "three"):
					return '3'
				case strings.HasSuffix(string(num), "four"):
					return '4'
				case strings.HasSuffix(string(num), "five"):
					return '5'
				case strings.HasSuffix(string(num), "six"):
					return '6'
				case strings.HasSuffix(string(num), "seven"):
					return '7'
				case strings.HasSuffix(string(num), "eight"):
					return '8'
				case strings.HasSuffix(string(num), "nine"):
					return '9'
				}
				continue
			}
			if r < '0' || r > '9' {
				// reset num
				num = []rune{}
				continue
			}
			return r
		}
		panic("no digit found")
	}

	getDigitReversed := func(rs []rune) rune {
		num := []rune{}
		for idx := len(rs) - 1; idx >= 0; idx-- {
			r := rs[idx]
			if r >= 'a' && r <= 'z' {
				num = append(num, r)
				switch {
				case strings.HasSuffix(string(num), "eno"):
					return '1'
				case strings.HasSuffix(string(num), "owt"):
					return '2'
				case strings.HasSuffix(string(num), "eerht"):
					return '3'
				case strings.HasSuffix(string(num), "ruof"):
					return '4'
				case strings.HasSuffix(string(num), "evif"):
					return '5'
				case strings.HasSuffix(string(num), "xis"):
					return '6'
				case strings.HasSuffix(string(num), "neves"):
					return '7'
				case strings.HasSuffix(string(num), "thgie"):
					return '8'
				case strings.HasSuffix(string(num), "enin"):
					return '9'
				}
				continue
			}
			if r < '0' || r > '9' {
				// reset num
				num = []rune{}
				continue
			}
			return r
		}
		panic("no digit found")
	}

	for i, line := range strings.Split(input, "\n") {
		runes := []rune(line)
		digits := make([]rune, 0, 2)
		digits = append(digits, getDigit(runes))
		digits = append(digits, getDigitReversed(runes))
		coord, err := strconv.Atoi(string(digits))
		if err != nil {
			log.Fatalf("Error parsing input: %v", err)
		}
		if i <= 50 {
			fmt.Println("line:", line, "Coord:", coord)
		}
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
