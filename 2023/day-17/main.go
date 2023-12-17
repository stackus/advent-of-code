package main

import (
	"container/heap"
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

	start := point{y: 0, x: 0}
	end := point{y: len(grid) - 1, x: len(grid[0]) - 1}

	return findLowestCost(grid, start, end, 1, 3)
}

// puzzle2 solves the level 2 puzzle
func puzzle2(input string) int {
	grid := parseInput(input)

	start := point{y: 0, x: 0}
	end := point{y: len(grid) - 1, x: len(grid[0]) - 1}

	return findLowestCost(grid, start, end, 4, 10)
}

type point struct {
	x, y int
}

// basically copied the https://pkg.go.dev/container/heap (PriorityQueue example) to start with
type item struct {
	index    int // The index of the item in the heap.
	cost     int // The priority of the item in the queue.
	x, y     int // The value of the item; arbitrary.
	axisLock int // The value of the item; arbitrary.
}

type priorityQueue []*item

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

var directions = [][2]int{
	{0, 1},  // right; locked on x axis (0) movement
	{1, 0},  // down; locked on y axis (1) movement
	{0, -1}, // left; locked on x axis (0) movement
	{-1, 0}, // up; locked on y axis (1) movement
}

// implement a dijkstra's algorithm with a min/max move distance
func findLowestCost(grid [][]int, start, end point, minDistance, maxDistance int) int {
	queue := &priorityQueue{}
	heap.Init(queue)
	heap.Push(queue, &item{
		cost:     0,
		x:        start.x,
		y:        start.y,
		axisLock: -1,
	})
	visited := make(map[[3]int]bool)
	costs := make(map[[3]int]int)

	for queue.Len() > 0 {
		current := heap.Pop(queue).(*item)
		// return if we've reached the end
		if current.x == end.x && current.y == end.y {
			return current.cost
		}
		// skip if we've already visited this point
		if visited[[3]int{current.x, current.y, current.axisLock}] {
			continue
		}
		visited[[3]int{current.x, current.y, current.axisLock}] = true

		for direction, dirAdjustment := range directions {
			cost := 0
			axis := direction % 2
			if axis == current.axisLock {
				continue
			}
			// push all possible moves in this direction
			for distance := 1; distance <= maxDistance; distance++ {
				newX := current.x + dirAdjustment[0]*distance
				newY := current.y + dirAdjustment[1]*distance
				// break if the move will be out of bounds; all subsequent moves will also be out of bounds
				if newX < 0 || newY < 0 || newX >= len(grid) || newY >= len(grid[0]) {
					break
				}
				// sum the cost of the move so far
				cost += grid[newX][newY]
				// skip if the distance is less than the minimum; puzzle 2 sets a minimum distance
				if distance < minDistance {
					continue
				}
				newCost := current.cost + cost
				// skip if we've already visited this point with a lower cost in this direction
				if currentCost, ok := costs[[3]int{newX, newY, direction}]; ok && currentCost <= newCost {
					continue
				}
				costs[[3]int{newX, newY, direction}] = newCost
				heap.Push(queue, &item{
					cost:     newCost,
					x:        newX,
					y:        newY,
					axisLock: axis,
				})
			}
		}
	}
	return -1 // If no path is found
}

// parseInput converts the input string into whatever format is needed for the puzzle
// update the return type as needed
func parseInput(input string) (lines [][]int) {
	for _, line := range strings.Split(input, "\n") {
		nums := make([]int, 0, len(line))
		for _, r := range line {
			nums = append(nums, int(r-'0'))
		}
		lines = append(lines, nums)
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

	solutionPath := filepath.Join(GetPuzzlePath(17, 2023), fmt.Sprintf("solution-%d.txt", puzzle))
	err := WriteFile(solutionPath, []byte(fmt.Sprintf("%d", solution)), true)
	if err != nil {
		log.Fatalf("Error writing solution: %v", err)
	}
	fmt.Println("Solution:", solution)
}
