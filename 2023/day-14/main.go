package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	. "github.com/stackus/advent-of-code"
)

//go:embed input.txt
var input string

// puzzle1 solves the level 1 puzzle
func puzzle1(input string) int {
	field := parseInput(input)

	// // print field
	// for _, line := range field {
	// 	fmt.Println(string(line))
	// }

	// tilt north
	tiltNorth(field)

	// fmt.Println()
	// // print field
	// for _, line := range field {
	// 	fmt.Println(string(line))
	// }

	return totalField(field)
}

// puzzle2 solves the level 2 puzzle
func puzzle2(input string) int {
	field := parseInput(input)

	// tilt north
	rotations := 1_000_000_000
	seen := map[string][]int{}
	for r := 0; r < rotations; r++ {
		// rotate the field
		tiltNorth(field)
		tiltWest(field)
		tiltSouth(field)
		tiltEast(field)

		builder := strings.Builder{}
		for _, line := range field {
			builder.WriteString(string(line))
		}
		s := builder.String()

		// for the code below; add 1 to r because we've already performed the rotation above

		if hit := seen[s]; hit != nil {
			cycleLen := (r + 1) - hit[0]
			remaining := rotations - (r + 1)
			// if the remaining rotations are a multiple of the cycle length, we can just return the value
			if remaining%cycleLen == 0 {
				return hit[1]
			}
		}

		// record the rotation and total for the current field
		seen[s] = []int{r + 1, totalField(field)}
	}

	return 0
}

func totalField(field [][]rune) int {
	total := 0
	lines := len(field)
	for i, line := range field {
		total += (lines - i) * strings.Count(string(line), "O")
	}
	return total
}

func tiltNorth(field [][]rune) {
	for x := 0; x < len(field[0]); x++ {
		var open []int
		for y := 0; y < len(field); y++ {
			if field[y][x] == '.' {
				open = append(open, y)
				continue
			}
			if field[y][x] == 'O' {
				if len(open) > 0 {
					field[y][x] = '.'
					field[open[0]][x] = 'O'
					open = open[1:]
					open = append(open, y)
				}
				continue
			}
			if field[y][x] == '#' {
				open = []int{}
				continue
			}
		}
	}
}

func tiltWest(field [][]rune) {
	for y := 0; y < len(field); y++ {
		var open []int
		for x := 0; x < len(field[y]); x++ {
			if field[y][x] == '.' {
				open = append(open, x)
				continue
			}
			if field[y][x] == 'O' {
				if len(open) > 0 {
					field[y][x] = '.'
					field[y][open[0]] = 'O'
					open = open[1:]
					open = append(open, x)
				}
				continue
			}
			if field[y][x] == '#' {
				open = []int{}
				continue
			}
		}
	}
}

func tiltSouth(field [][]rune) {
	for x := 0; x < len(field[0]); x++ {
		var open []int
		for y := len(field) - 1; y >= 0; y-- {
			if field[y][x] == '.' {
				open = append(open, y)
				continue
			}
			if field[y][x] == 'O' {
				if len(open) > 0 {
					field[y][x] = '.'
					field[open[0]][x] = 'O'
					open = open[1:]
					open = append(open, y)
				}
				continue
			}
			if field[y][x] == '#' {
				open = []int{}
				continue
			}
		}
	}
}

func tiltEast(field [][]rune) {
	for y := 0; y < len(field); y++ {
		var open []int
		for x := len(field[y]) - 1; x >= 0; x-- {
			if field[y][x] == '.' {
				open = append(open, x)
				continue
			}
			if field[y][x] == 'O' {
				if len(open) > 0 {
					field[y][x] = '.'
					field[y][open[0]] = 'O'
					open = open[1:]
					open = append(open, x)
				}
				continue
			}
			if field[y][x] == '#' {
				open = []int{}
				continue
			}
		}
	}
}

// parseInput converts the input string into whatever format is needed for the puzzle
// update the return type as needed
func parseInput(input string) [][]rune {
	field := [][]rune{}

	for _, line := range strings.Split(input, "\n") {
		field = append(field, []rune(line))
	}

	return field
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

	solutionPath := filepath.Join(GetPuzzlePath(14, 2023), fmt.Sprintf("solution-%d.txt", puzzle))
	err := WriteFile(solutionPath, []byte(fmt.Sprintf("%d", solution)), true)
	if err != nil {
		log.Fatalf("Error writing solution: %v", err)
	}
	fmt.Println("Solution:", solution)
}
