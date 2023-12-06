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
	times, distances := parseInput(input)

	wins := make([]int, len(times))
	for race, tm := range times {
		for s := 1; s < tm; s++ {
			distance := s * (tm - s)
			if distance > distances[race] {
				wins[race]++
			}
		}
	}

	total := 1
	for _, v := range wins {
		total *= v
	}

	return total
}

// puzzle2 solves the level 2 puzzle
func puzzle2(input string) int {
	times, distances := parseInput(input)

	s := fmt.Sprintf("%d%d%d%d", times[0], times[1], times[2], times[3])
	totalTime, _ := strconv.Atoi(s)

	s = fmt.Sprintf("%d%d%d%d", distances[0], distances[1], distances[2], distances[3])
	totalDistance, _ := strconv.Atoi(s)

	wins := 0
	for s := 1; s < totalTime; s++ {
		distance := s * (totalTime - s)
		if distance > totalDistance {
			wins++
		}
	}

	return wins
}

// parseInput converts the input string into whatever format is needed for the puzzle
// update the return type as needed
func parseInput(input string) (times, distances []int) {
	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`\d+`)

	s := re.FindAllString(lines[0], -1)
	for _, v := range s {
		i, _ := strconv.Atoi(v)
		times = append(times, i)
	}
	s = re.FindAllString(lines[1], -1)
	for _, v := range s {
		i, _ := strconv.Atoi(v)
		distances = append(distances, i)
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

	started := time.Now()
	solution := 0
	if puzzle == 1 {
		solution = puzzle1(input)
	} else {
		solution = puzzle2(input)
	}
	fmt.Println("Completed in", time.Since(started))

	solutionPath := filepath.Join(GetPuzzlePath(6, 2023), fmt.Sprintf("solution-%d.txt", puzzle))
	err := WriteFile(solutionPath, []byte(fmt.Sprintf("%d", solution)), true)
	if err != nil {
		log.Fatalf("Error writing solution: %v", err)
	}
	fmt.Println("Solution:", solution)
}
