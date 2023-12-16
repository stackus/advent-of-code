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
	total := energizeNodes(grid, [][4]int{{0, -1, 0, 1}})

	return total
}

// puzzle2 solves the level 2 puzzle
func puzzle2(input string) int {
	grid := parseInput(input)

	total := 0

	for y := 0; y < len(grid); y++ {
		// in from the left
		total = max(total, energizeNodes(grid, [][4]int{{y, -1, 0, 1}}))
		// in from the right
		total = max(total, energizeNodes(grid, [][4]int{{y, len(grid[0]), 0, -1}}))
	}
	for x := 0; x < len(grid[0]); x++ {
		// in from the top
		total = max(total, energizeNodes(grid, [][4]int{{-1, x, 1, 0}}))
		// in from the bottom
		total = max(total, energizeNodes(grid, [][4]int{{len(grid), x, -1, 0}}))
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

func energizeNodes(grid [][]rune, queue [][4]int) int {
	// map of point+direction to boolean
	// [y,x, yDir, xDir]
	// yDir and xDir are -1, 0, or 1
	visited := map[[4]int]bool{}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		dirY := item[2]
		dirX := item[3]
		nextY := item[0] + dirY
		nextX := item[1] + dirX

		move := [4]int{nextY, nextX, dirY, dirX}
		if visited[move] {
			continue
		}
		if nextY < 0 || nextX < 0 || nextY >= len(grid) || nextX >= len(grid[0]) {
			continue
		}
		visited[move] = true
		switch grid[nextY][nextX] {
		case '.':
			// moving left or right is not changed
			// moving up and down is not changed
			queue = append(queue, [4]int{nextY, nextX, dirY, dirX})
		case '/':
			// moving to the right (0, 1) becomes up (-1, 0)
			// moving to the left (0, -1) becomes down (1, 0)
			// moving up (-1, 0) becomes right (0, 1)
			// moving down (1, 0) becomes left (0, -1)
			dirY, dirX = -dirX, -dirY
			queue = append(queue, [4]int{nextY, nextX, dirY, dirX})
		case '\\':
			// moving to the right (0, 1) becomes down (1, 0)
			// moving to the left (0, -1) becomes up (-1, 0)
			// moving up (-1, 0) becomes left (0, -1)
			// moving down (1, 0) becomes right (0, 1)
			dirY, dirX = dirX, dirY
			queue = append(queue, [4]int{nextY, nextX, dirY, dirX})
		case '-':
			// moving left or right is not changed
			// moving up and down will split into two moves heading left and right
			queue = append(queue, [4]int{nextY, nextX, 0, 1}, [4]int{nextY, nextX, 0, -1})
		case '|':
			// moving up or down is not changed
			// moving left and right will split into two moves heading up and down
			queue = append(queue, [4]int{nextY, nextX, 1, 0}, [4]int{nextY, nextX, -1, 0})
		}
	}
	nodes := map[[2]int]bool{}
	for visited := range visited {
		nodes[[2]int{visited[0], visited[1]}] = true
	}
	return len(nodes)
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

	solutionPath := filepath.Join(GetPuzzlePath(16, 2023), fmt.Sprintf("solution-%d.txt", puzzle))
	err := WriteFile(solutionPath, []byte(fmt.Sprintf("%d", solution)), true)
	if err != nil {
		log.Fatalf("Error writing solution: %v", err)
	}
	fmt.Println("Solution:", solution)
}
