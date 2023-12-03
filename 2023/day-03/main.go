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
	matches := parseInput(input)
	total := 0

	for _, match := range matches {
		for _, num := range match.nums {
			total += num.n
		}
	}

	return total
}

// puzzle2 solves the level 2 puzzle
func puzzle2(input string) int {
	matches := parseInput(input)
	total := 0

	for _, match := range matches {
		if match.r != '*' || len(match.nums) != 2 {
			continue
		}
		total += match.nums[0].n * match.nums[1].n
	}

	return total
}

type number struct {
	n int
	l int
	c int
}

type runeMatch struct {
	line int
	char int
	r    rune
	nums []number
}

// parseInput converts the input string into whatever format is needed for the puzzle
// update the return type as needed
func parseInput(input string) (matches []runeMatch) {
	var runes [][]rune
	for _, line := range strings.Split(input, "\n") {
		runes = append(runes, []rune(line))
	}

	numbers := map[int]map[int]number{}
	for lineNum, line := range strings.Split(input, "\n") {
		numbers[lineNum] = make(map[int]number)
		digits := make([]rune, 0)
		for charNum, char := range line {
			if char >= '0' && char <= '9' {
				digits = append(digits, char)
			}
			if char < '0' || char > '9' || charNum == len(line)-1 {
				if len(digits) != 0 {
					num, _ := strconv.Atoi(string(digits))
					for i := 0; i < len(digits); i++ {
						numbers[lineNum][charNum-i-1] = number{
							n: num,
							l: lineNum,
							c: charNum,
						}
					}
				}
				digits = digits[:0]
				continue
			}
		}
	}

	for lineNum, line := range runes {
		for charNum, char := range line {
			if char == '.' || (char >= '0' && char <= '9') {
				continue
			}
			match := runeMatch{
				line: lineNum,
				char: charNum,
				r:    char,
				nums: make([]number, 0),
			}

			// now check if there are any digits around it
			startLine := max(0, lineNum-1)
			endLine := min(len(runes)-1, lineNum+1)
			startChar := max(0, charNum-1)
			endChar := min(len(line)-1, charNum+1)
			for checkLineNum := startLine; checkLineNum <= endLine; checkLineNum++ {
				for checkCharNum := startChar; checkCharNum <= endChar; checkCharNum++ {
					if checkLineNum == lineNum && checkCharNum == charNum {
						continue
					}
					if num, ok := numbers[checkLineNum][checkCharNum]; ok {
						added := false
						for _, n := range match.nums {
							if n.c == num.c && n.l == num.l {
								added = true
								break
							}
						}
						if !added {
							match.nums = append(match.nums, num)
						}
					}
				}
			}
			if len(match.nums) == 0 {
				continue
			}
			matches = append(matches, match)
		}
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

	solutionPath := filepath.Join(GetPuzzlePath(3, 2023), fmt.Sprintf("solution-%d.txt", puzzle))
	err := WriteFile(solutionPath, []byte(fmt.Sprintf("%d", solution)), true)
	if err != nil {
		log.Fatalf("Error writing solution: %v", err)
	}
	fmt.Println("Solution: ", solution)
}
