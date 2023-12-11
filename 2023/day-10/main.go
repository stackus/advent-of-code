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
	grid := parseInput(input)
	start := findStartingPosition(grid)

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	return bfs(grid, visited, start.x, start.y, isConnected)
}

// puzzle2 solves the level 2 puzzle
func puzzle2(input string) int {
	grid := parseInput(input)
	start := findStartingPosition(grid)

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	// solve again to fill visited
	_ = bfs(grid, visited, start.x, start.y, isConnected)

	total := 0

	exploded := constructExplodedGrid(grid, visited)

	// ASCII art!
	// for _, row := range exploded {
	// 	fmt.Println(string(row))
	// }

	revisited := make([][]bool, len(exploded))
	for i := range revisited {
		revisited[i] = make([]bool, len(exploded[i]))
	}

	_ = bfs(exploded, revisited, 0, 0, func(from, to point, grid [][]rune) bool {
		if to.x < 0 || to.y < 0 || to.y >= len(grid) || to.x >= len(grid[0]) {
			return false
		}
		return grid[to.y][to.x] == ' '
	})

	// // Filled in ASCII art!
	// for y, row := range exploded {
	// 	for x, char := range row {
	// 		if char == '*' {
	// 			fmt.Print("*")
	// 			continue
	// 		}
	// 		if revisited[y][x] {
	// 			fmt.Print("O")
	// 			continue
	// 		}
	// 		fmt.Print(" ")
	// 	}
	// 	fmt.Println()
	// }

	for y, row := range grid {
		for x := range row {
			if visited[y][x] {
				continue
			}
			if allPointsAreNotRevisited([]point{
				{y: y * 3, x: x * 3}, {y: y * 3, x: x*3 + 1}, {y: y * 3, x: x*3 + 2},
				{y: y*3 + 1, x: x * 3}, {y: y*3 + 1, x: x*3 + 1}, {y: y*3 + 1, x: x*3 + 2},
				{y: y*3 + 2, x: x * 3}, {y: y*3 + 2, x: x*3 + 1}, {y: y*3 + 2, x: x*3 + 2},
			}, revisited) {
				total++
			}
		}
	}

	return total
}

type point struct {
	y, x, distance int
}

func allPointsAreNotRevisited(points []point, grid [][]bool) bool {
	for _, point := range points {
		if grid[point.y][point.x] {
			return false
		}
	}
	return true
}

func constructExplodedGrid(grid [][]rune, visited [][]bool) [][]rune {
	exploded := make([][]rune, len(grid)*3)
	for i := range exploded {
		exploded[i] = make([]rune, len(grid[0])*3)
	}

	for y, row := range grid {
		for x, char := range row {
			if !visited[y][x] {
				exploded[y*3+0][x*3+0] = ' '
				exploded[y*3+0][x*3+1] = ' '
				exploded[y*3+0][x*3+2] = ' '
				exploded[y*3+1][x*3+0] = ' '
				exploded[y*3+1][x*3+1] = ' '
				exploded[y*3+1][x*3+2] = ' '
				exploded[y*3+2][x*3+0] = ' '
				exploded[y*3+2][x*3+1] = ' '
				exploded[y*3+2][x*3+2] = ' '
				continue
			}
			switch char {
			case '|':
				exploded[y*3+0][x*3+0] = ' '
				exploded[y*3+0][x*3+1] = '*'
				exploded[y*3+0][x*3+2] = ' '
				exploded[y*3+1][x*3+0] = ' '
				exploded[y*3+1][x*3+1] = '*'
				exploded[y*3+1][x*3+2] = ' '
				exploded[y*3+2][x*3+0] = ' '
				exploded[y*3+2][x*3+1] = '*'
				exploded[y*3+2][x*3+2] = ' '
			case '-':
				exploded[y*3+0][x*3+0] = ' '
				exploded[y*3+0][x*3+1] = ' '
				exploded[y*3+0][x*3+2] = ' '
				exploded[y*3+1][x*3+0] = '*'
				exploded[y*3+1][x*3+1] = '*'
				exploded[y*3+1][x*3+2] = '*'
				exploded[y*3+2][x*3+0] = ' '
				exploded[y*3+2][x*3+1] = ' '
				exploded[y*3+2][x*3+2] = ' '
			case 'L':
				exploded[y*3+0][x*3+0] = ' '
				exploded[y*3+0][x*3+1] = '*'
				exploded[y*3+0][x*3+2] = ' '
				exploded[y*3+1][x*3+0] = ' '
				exploded[y*3+1][x*3+1] = '*'
				exploded[y*3+1][x*3+2] = '*'
				exploded[y*3+2][x*3+0] = ' '
				exploded[y*3+2][x*3+1] = ' '
				exploded[y*3+2][x*3+2] = ' '
			case 'J':
				exploded[y*3+0][x*3+0] = ' '
				exploded[y*3+0][x*3+1] = '*'
				exploded[y*3+0][x*3+2] = ' '
				exploded[y*3+1][x*3+0] = '*'
				exploded[y*3+1][x*3+1] = '*'
				exploded[y*3+1][x*3+2] = ' '
				exploded[y*3+2][x*3+0] = ' '
				exploded[y*3+2][x*3+1] = ' '
				exploded[y*3+2][x*3+2] = ' '
			case '7':
				exploded[y*3+0][x*3+0] = ' '
				exploded[y*3+0][x*3+1] = ' '
				exploded[y*3+0][x*3+2] = ' '
				exploded[y*3+1][x*3+0] = '*'
				exploded[y*3+1][x*3+1] = '*'
				exploded[y*3+1][x*3+2] = ' '
				exploded[y*3+2][x*3+0] = ' '
				exploded[y*3+2][x*3+1] = '*'
				exploded[y*3+2][x*3+2] = ' '
			case 'F':
				exploded[y*3+0][x*3+0] = ' '
				exploded[y*3+0][x*3+1] = ' '
				exploded[y*3+0][x*3+2] = ' '
				exploded[y*3+1][x*3+0] = ' '
				exploded[y*3+1][x*3+1] = '*'
				exploded[y*3+1][x*3+2] = '*'
				exploded[y*3+2][x*3+0] = ' '
				exploded[y*3+2][x*3+1] = '*'
				exploded[y*3+2][x*3+2] = ' '
			case 'S':
				exploded[y*3+0][x*3+0] = '*'
				exploded[y*3+0][x*3+1] = '*'
				exploded[y*3+0][x*3+2] = '*'
				exploded[y*3+1][x*3+0] = '*'
				exploded[y*3+1][x*3+1] = '*'
				exploded[y*3+1][x*3+2] = '*'
				exploded[y*3+2][x*3+0] = '*'
				exploded[y*3+2][x*3+1] = '*'
				exploded[y*3+2][x*3+2] = '*'

			}
		}
	}

	return exploded
}

func findStartingPosition(grid [][]rune) point {
	for y, row := range grid {
		for x, char := range row {
			if char == 'S' {
				return point{x: x, y: y}
			}
		}
	}
	return point{}
}

func bfs(grid [][]rune, visited [][]bool, x, y int, canVisitFn func(from, to point, grid [][]rune) bool) int {
	queue := []point{{x: x, y: y, distance: 0}}
	maxDistance := 0

	for len(queue) > 0 {
		from := queue[0]
		queue = queue[1:]

		if visited[from.y][from.x] {
			continue
		}

		visited[from.y][from.x] = true
		maxDistance = from.distance

		for _, dir := range getPossibleDirections(grid[from.y][from.x]) {
			to := point{x: from.x + dir.x, y: from.y + dir.y, distance: from.distance + 1}
			if canVisitFn(from, to, grid) && !visited[to.y][to.x] {
				queue = append(queue, to)
			}
		}
	}

	return maxDistance
}

func getPossibleDirections(char rune) []point {
	switch char {
	case '|':
		return []point{{y: -1, x: 0}, {y: 1, x: 0}}
	case '-':
		return []point{{y: 0, x: -1}, {y: 0, x: 1}}
	case 'L':
		return []point{{y: -1, x: 0}, {y: 0, x: 1}}
	case 'J':
		return []point{{y: -1, x: 0}, {y: 0, x: -1}}
	case '7':
		return []point{{y: 1, x: 0}, {y: 0, x: -1}}
	case 'F':
		return []point{{y: 1, x: 0}, {y: 0, x: 1}}
	case 'S', ' ':
		return []point{{y: -1, x: 0}, {y: 1, x: 0}, {y: 0, x: -1}, {y: 0, x: 1}}
	default:
		return []point{}
	}
}

func isConnected(from, to point, grid [][]rune) bool {
	if to.x < 0 || to.y < 0 || to.y >= len(grid) || to.x >= len(grid[0]) {
		return false
	}

	// if to tile is `.` then return false
	if grid[to.y][to.x] == '.' {
		return false
	}

	// Check if the pipe at (to.x, to.y) connects to the pipe at (from.x, from.y)
	// This is done by checking if the pipe at (to.x, to.y) has a direction that
	// points to (from.x, from.y)
	for _, dir := range getPossibleDirections(grid[to.y][to.x]) {
		if dir.x == from.x-to.x && dir.y == from.y-to.y {
			return true
		}
	}

	return false
}

// parseInput converts the input string into whatever format is needed for the puzzle
// update the return type as needed
func parseInput(input string) [][]rune {
	var grid [][]rune
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}

	return grid
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

	solutionPath := filepath.Join(GetPuzzlePath(10, 2023), fmt.Sprintf("solution-%d.txt", puzzle))
	err := WriteFile(solutionPath, []byte(fmt.Sprintf("%d", solution)), true)
	if err != nil {
		log.Fatalf("Error writing solution: %v", err)
	}
	fmt.Println("Solution:", solution)
}
