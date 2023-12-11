package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"math"
	"path/filepath"
	"strings"
	"time"

	. "github.com/stackus/advent-of-code"
)

//go:embed input.txt
var input string

// puzzle1 solves the level 1 puzzle
func puzzle1(input string) int {
	space := parseInput(input)

	galaxies := shiftGalaxies(space, 1, findGalaxies(space))

	total := 0

	for i, galaxy := range galaxies {
		if i < len(galaxies)-1 {
			for _, otherGalaxy := range galaxies[i+1:] {
				distance := math.Abs(float64(galaxy.y-otherGalaxy.y)) + math.Abs(float64(galaxy.x-otherGalaxy.x))
				total += int(distance)
			}
		}
	}

	return total
}

// puzzle2 solves the level 2 puzzle
func puzzle2(input string) int {
	space := parseInput(input)

	galaxies := shiftGalaxies(space, 1_000_000, findGalaxies(space))

	total := 0

	for i, galaxy := range galaxies {
		if i < len(galaxies)-1 {
			for _, otherGalaxy := range galaxies[i+1:] {
				distance := math.Abs(float64(galaxy.y-otherGalaxy.y)) + math.Abs(float64(galaxy.x-otherGalaxy.x))
				total += int(distance)
			}
		}
	}

	return total
}

// parseInput converts the input string into whatever format is needed for the puzzle
// update the return type as needed
func parseInput(input string) (lines [][]rune) {
	for _, line := range strings.Split(input, "\n") {
		lines = append(lines, []rune(line))
	}

	return
}

type galaxy struct {
	x  int
	y  int
	ox int
	oy int
}

func findGalaxies(space [][]rune) []galaxy {
	var galaxies []galaxy

	for y, line := range space {
		for x, char := range line {
			if char == '#' {
				galaxies = append(galaxies, galaxy{x: x, y: y, ox: x, oy: y})
			}
		}
	}

	return galaxies
}

func shiftGalaxies(space [][]rune, shift int, galaxies []galaxy) []galaxy {
	for y := 0; y < len(space); y++ {
		allDots := true
		for x := range space[y] {
			if space[y][x] != '.' {
				allDots = false
				break
			}
		}
		if !allDots {
			continue
		}
		// line is all '.'
		// shift galaxies up by `shift` if their y is greater than this line
		for i, galaxy := range galaxies {
			if galaxy.oy > y {
				galaxies[i].y += max(1, shift-1)
			}
		}
	}

	// check if the column is all '.'
	for x := 0; x < len(space[0]); x++ {
		allDots := true
		for y := 0; y < len(space); y++ {
			if space[y][x] != '.' {
				allDots = false
				break
			}
		}
		if !allDots {
			continue
		}
		for i, galaxy := range galaxies {
			if galaxy.ox > x {
				galaxies[i].x += max(1, shift-1)
			}
		}
	}
	return galaxies
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

	solutionPath := filepath.Join(GetPuzzlePath(11, 2023), fmt.Sprintf("solution-%d.txt", puzzle))
	err := WriteFile(solutionPath, []byte(fmt.Sprintf("%d", solution)), true)
	if err != nil {
		log.Fatalf("Error writing solution: %v", err)
	}
	fmt.Println("Solution:", solution)
}
