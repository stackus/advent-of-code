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
	codes := parseInput(input)

	total := 0
	for _, code := range codes {
		total += computeCode(code)
	}

	return total
}

// puzzle2 solves the level 2 puzzle
func puzzle2(input string) int {
	cmds := parseInput(input)

	total := 0

	boxes := boxMap{}

	for _, cmd := range cmds {
		if label, ok := strings.CutSuffix(cmd, "-"); ok {
			boxNum := computeCode(label)
			boxes.removeLens(boxNum, label)
			continue
		}

		if label, value, ok := strings.Cut(cmd, "="); ok {
			boxNum := computeCode(label)
			v, _ := strconv.Atoi(value)
			boxes.addLens(boxNum, lens{label, v})
			continue
		}
	}

	for boxNum := 0; boxNum < 256; boxNum++ {
		fmt.Printf("%d: ", boxNum)
		for _, lens := range boxes[boxNum] {
			fmt.Printf("[%s, %d]", lens.label, lens.value)
		}
		fmt.Println()
		total += boxes.sumLens(boxNum)
	}

	return total
}

type boxMap map[int][]lens

func (b *boxMap) addLens(boxNum int, l lens) {
	if _, ok := (*b)[boxNum]; !ok {
		(*b)[boxNum] = []lens{}
	}

	// replace the lens if it already exists
	for i, lens := range (*b)[boxNum] {
		if lens.label == l.label {
			(*b)[boxNum][i] = l
			return
		}
	}

	(*b)[boxNum] = append((*b)[boxNum], l)
}

func (b *boxMap) removeLens(boxNum int, label string) {
	if _, ok := (*b)[boxNum]; !ok {
		return
	}

	for i, lens := range (*b)[boxNum] {
		if lens.label == label {
			(*b)[boxNum] = append((*b)[boxNum][:i], (*b)[boxNum][i+1:]...)
			return
		}
	}
}

func (b *boxMap) sumLens(boxNum int) int {
	total := 0

	if _, ok := (*b)[boxNum]; !ok {
		return total
	}

	for i, lens := range (*b)[boxNum] {
		total += (boxNum + 1) * (i + 1) * lens.value
	}

	return total
}

type lens struct {
	label string
	value int
}

func computeCode(code string) int {
	codeTotal := 0

	for _, char := range code {
		codeTotal += int(char)
		codeTotal *= 17
		codeTotal %= 256
	}

	return codeTotal
}

// parseInput converts the input string into whatever format is needed for the puzzle
// update the return type as needed
func parseInput(input string) []string {
	for _, line := range strings.Split(input, "\n") {
		return strings.Split(line, ",")
	}

	return []string{}
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

	solutionPath := filepath.Join(GetPuzzlePath(15, 2023), fmt.Sprintf("solution-%d.txt", puzzle))
	err := WriteFile(solutionPath, []byte(fmt.Sprintf("%d", solution)), true)
	if err != nil {
		log.Fatalf("Error writing solution: %v", err)
	}
	fmt.Println("Solution:", solution)
}
