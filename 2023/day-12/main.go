package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	. "github.com/stackus/advent-of-code"
)

//go:embed input.txt
var input string

// puzzle1 solves the level 1 puzzle
func puzzle1(input string) int {
	reports := parseInput(input)

	total := 0

	for _, r := range reports {
		cache = make(map[cacheKey]int)
		total += countArrangements(r.field+".", r.groups, 0, 0, 0)
	}

	return total
}

// puzzle2 solves the level 2 puzzle
func puzzle2(input string) int {
	reports := parseInput(input)

	total := 0

	for _, r := range reports {
		field := strings.Join([]string{r.field, r.field, r.field, r.field, r.field}, "?")
		var groups []int
		for i := 0; i < 5; i++ {
			groups = append(groups, r.groups...)
		}
		cache = make(map[cacheKey]int)
		count := countArrangements(field+".", groups, 0, 0, 0)
		total += count
	}

	return total
}

type report struct {
	field  string
	groups []int
}

type cacheKey struct {
	fieldIdx, groupIdx, hashLen int
}

var cache = make(map[cacheKey]int)

func countArrangements(field string, groups []int, fieldIdx, groupIdx, hashLen int) (result int) {
	// create a cache key
	key := cacheKey{fieldIdx, groupIdx, hashLen}
	// if we've already calculated the combo of (fieldIdx, groupIdx, hashLen), return the cached value
	if val, ok := cache[key]; ok {
		return val
	}
	// defer the caching of the result until the end of the function
	defer func() {
		cache[key] = result
	}()

	// we've reached the end of the field
	if fieldIdx == len(field) {
		// if we've also reached the end of the groups, we've found an arrangement
		if groupIdx == len(groups) {
			return 1
		}

		return 0
	}

	// we've encountered a hash
	if field[fieldIdx] == '#' {
		return countArrangements(field, groups, fieldIdx+1, groupIdx, hashLen+1)
	}

	// if we've encountered a dot, or we've reached the end of the groups then ...
	if field[fieldIdx] == '.' || groupIdx == len(groups) {
		if groupIdx < len(groups) && hashLen == groups[groupIdx] {
			// if the hashLen matches the current group length, we can move on to the next group
			return countArrangements(field, groups, fieldIdx+1, groupIdx+1, 0)
		} else if hashLen == 0 {
			// or if the hashLen is 0, we can move on to the next character in the field
			return countArrangements(field, groups, fieldIdx+1, groupIdx, 0)
		}

		// otherwise, we've encountered a dot and the hashLen doesn't match the current group
		return 0
	}

	// we've encountered a question mark
	hashCount := countArrangements(field, groups, fieldIdx+1, groupIdx, hashLen+1)

	var dotCount int
	if hashLen == groups[groupIdx] {
		// if the hashLen matches the current group length, we can move on to the next group
		dotCount = countArrangements(field, groups, fieldIdx+1, groupIdx+1, 0)
	} else if hashLen == 0 {
		// or if the hashLen is 0, we can move on to the next character in the field and try to match a '#' or '?'
		dotCount = countArrangements(field, groups, fieldIdx+1, groupIdx, 0)
	}

	// return the sum of the hashCount and dotCount
	return hashCount + dotCount
}

// parseInput converts the input string into whatever format is needed for the puzzle
// update the return type as needed
func parseInput(input string) []report {
	var reports []report
	for _, line := range strings.Split(input, "\n") {
		r := report{}
		parts := strings.Split(line, " ")
		r.field = parts[0]
		for _, group := range strings.Split(parts[1], ",") {
			num, _ := strconv.Atoi(group)
			r.groups = append(r.groups, num)
		}
		reports = append(reports, r)
	}

	return reports
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

	solutionPath := filepath.Join(GetPuzzlePath(12, 2023), fmt.Sprintf("solution-%d.txt", puzzle))
	err := WriteFile(solutionPath, []byte(fmt.Sprintf("%d", solution)), true)
	if err != nil {
		log.Fatalf("Error writing solution: %v", err)
	}
	fmt.Println("Solution:", solution)
}
