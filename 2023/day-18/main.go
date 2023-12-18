package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"math"
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
func puzzle1(input string) int64 {
	instructions := parseInput(input)

	currentX, currentY := int64(0), int64(0)
	length := int64(0)
	var points []point
	for _, instruction := range instructions {
		switch instruction.dir {
		case 'U':
			currentY -= instruction.distance
		case 'D':
			currentY += instruction.distance
		case 'L':
			currentX -= instruction.distance
		case 'R':
			currentX += instruction.distance
		}
		length += instruction.distance
		points = append(points, point{currentX, currentY})
	}

	// return interior points plus the border
	return getInterior(points) + length/2 + 1
}

// puzzle2 solves the level 2 puzzle
func puzzle2(input string) int64 {
	instructions := parseInput(input)

	dirs := map[string]rune{
		"0": 'R',
		"1": 'D',
		"2": 'L',
		"3": 'U',
	}

	currentX, currentY := int64(0), int64(0)
	length := int64(0)
	var points []point
	for _, instruction := range instructions {
		distStr := instruction.color[:5]
		direction := instruction.color[5:]
		distance, err := strconv.ParseInt(distStr, 16, 64)
		if err != nil {
			panic(err)
		}
		dir := dirs[direction]
		switch dir {
		case 'U':
			currentY -= distance
		case 'D':
			currentY += distance
		case 'L':
			currentX -= distance
		case 'R':
			currentX += distance
		}
		length += distance
		points = append(points, point{currentX, currentY})
	}

	// return interior points plus the border
	return getInterior(points) + length/2 + 1
}

type instruction struct {
	dir      rune
	distance int64
	color    string
}

type point struct {
	x, y int64
}

// getInterior returns the number of points inside the polygon defined by the points
// uses the shoelace formula
func getInterior(points []point) int64 {
	sum := int64(0)
	n := len(points)
	for i := 0; i < n-1; i++ {
		sum += (points[i].y + points[i+1].y) * (points[i].x - points[i+1].x)
	}

	sum += (points[n-1].y + points[0].y) * (points[n-1].x - points[0].x)

	return int64(math.Abs(float64(sum / 2)))
}

// parseInput converts the input string into whatever format is needed for the puzzle
// update the return type as needed
func parseInput(input string) (lines []instruction) {
	re := regexp.MustCompile(`([UDRL]) (\d+) \(#([0-9a-f]{6})\)`)
	for _, line := range strings.Split(input, "\n") {
		m := re.FindStringSubmatch(line)
		if m == nil {
			panic("invalid input")
		}

		dir := rune(m[1][0])
		distance, _ := strconv.Atoi(m[2])
		color := m[3]

		lines = append(lines, instruction{dir, int64(distance), color})
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
	solution := int64(0)
	if puzzle == 1 {
		solution = puzzle1(input)
	} else {
		solution = puzzle2(input)
	}
	fmt.Println("Completed in", time.Since(started))

	solutionPath := filepath.Join(GetPuzzlePath(18, 2023), fmt.Sprintf("solution-%d.txt", puzzle))
	err := WriteFile(solutionPath, []byte(fmt.Sprintf("%d", solution)), true)
	if err != nil {
		log.Fatalf("Error writing solution: %v", err)
	}
	fmt.Println("Solution:", solution)
}
