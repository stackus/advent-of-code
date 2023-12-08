package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"golang.org/x/exp/constraints"

	. "github.com/stackus/advent-of-code"
)

//go:embed input.txt
var input string

// puzzle1 solves the level 1 puzzle
func puzzle1(input string) uint64 {
	directions, ns := parseInput(input)

	var total uint64 = 0
	label := "AAA"
	for {
		index := total % uint64(len(directions))
		dir := directions[index]
		label = ns[label][dir]
		total++
		if label == "ZZZ" {
			break
		}
	}

	return total
}

// puzzle2 solves the level 2 puzzle
func puzzle2(input string) uint64 {
	directions, ns := parseInput(input)

	var total uint64 = 0
	var labels []string
	// get all labels that end in A
	for label, _ := range ns {
		if label[2] == 'A' {
			labels = append(labels, label)
		}
	}

	values := make(chan uint64, len(labels))
	totals := make([]uint64, 0)

	var wg sync.WaitGroup

	for _, label := range labels {
		wg.Add(1)
		go func(label string) {
			total := uint64(0)
			for {
				index := total % uint64(len(directions))
				dir := directions[index]
				label = ns[label][dir]
				total++
				if strings.HasSuffix(label, "Z") {
					values <- total
					break
				}
			}
			wg.Done()
		}(label)
	}

	go func() {
		wg.Wait()
		close(values)
	}()

	for v := range values {
		totals = append(totals, v)
	}

	total = lcm(totals...)

	return total
}

// TODO: save this into some utils package
// lcm returns the least common multiple of the given numbers
func lcm[T constraints.Integer](nums ...T) T {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]
	for _, num := range nums[1:] {
		result = result / gcd(result, num) * num
	}
	return result
}

// gcd returns the greatest common divisor of the given numbers
func gcd[T constraints.Integer](a, b T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

type node map[string]string

type nodes map[string]node

// parseInput converts the input string into whatever format is needed for the puzzle
// update the return type as needed
func parseInput(input string) ([]string, nodes) {
	lines := strings.Split(input, "\n")
	directions := strings.Split(lines[0], "")

	ns := make(nodes)
	re := regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)

	for _, line := range lines[2:] {
		lines = append(lines, line)

		matches := re.FindStringSubmatch(line)
		label := matches[1]
		lft := matches[2]
		rgt := matches[3]
		ns[label] = map[string]string{
			"L": lft,
			"R": rgt,
		}
	}

	return directions, ns
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
	solution := uint64(0)
	if puzzle == 1 {
		solution = puzzle1(input)
	} else {
		solution = puzzle2(input)
	}
	fmt.Println("Completed in", time.Since(started))

	solutionPath := filepath.Join(GetPuzzlePath(8, 2023), fmt.Sprintf("solution-%d.txt", puzzle))
	err := WriteFile(solutionPath, []byte(fmt.Sprintf("%d", solution)), true)
	if err != nil {
		log.Fatalf("Error writing solution: %v", err)
	}
	fmt.Println("Solution:", solution)
}
