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
	fields := parseInput(input)

	total := 0
	for i, field := range fields {
		fmt.Println("Field: ", i+1)
		// for _, row := range field.rows {
		// 	fmt.Println(row)
		// }

		found, a, b := findReflection(field.columns, false)
		if found {
			fmt.Println("Found reflection in columns", a, b)
			total += b
		}
		found, a, b = findReflection(field.rows, false)
		if found {
			fmt.Println("Found reflection in rows", a, b)
			total += b * 100
		}
		fmt.Println()
	}

	return total
}

// puzzle2 solves the level 2 puzzle
func puzzle2(input string) int {
	fields := parseInput(input)

	total := 0
	for i, field := range fields {
		fmt.Println("Field: ", i+1)
		// for _, row := range field.rows {
		// 	fmt.Println(row)
		// }

		found, a, b := findReflection(field.columns, true)
		if found {
			fmt.Println("Found reflection in columns", a, b)
			total += b
		}
		found, a, b = findReflection(field.rows, true)
		if found {
			fmt.Println("Found reflection in rows", a, b)
			total += b * 100
		}
		fmt.Println()
	}

	return total
}

type data struct {
	rows    []string
	columns []string
}

func findReflection(field []string, withSmudge bool) (bool, int, int) {
	for i := 0; i < len(field)-1; i++ {
		j := i + 1
		if match, fixedSmudge := isReflection(field, i, j, false); match {
			if withSmudge == fixedSmudge {
				return true, i, j
			}
			continue
		}
	}
	return false, 0, 0
}

func isReflection(field []string, a, b int, fixedSmudge bool) (bool, bool) {
	if a < 0 || a >= len(field) || b < 0 || b >= len(field) {
		return true, fixedSmudge
	}
	if field[a] == field[b] {
		return isReflection(field, a-1, b+1, fixedSmudge)
	}
	if fixedSmudge {
		return false, true
	}

	// check if the two strings are exactly one character different
	smudge := false
	for i := 0; i < len(field[a]); i++ {
		if field[a][i] != field[b][i] {
			if smudge {
				return false, false
			}
			smudge = true
		}
	}
	if smudge {
		return isReflection(field, a-1, b+1, true)
	}

	return false, fixedSmudge
}

// parseInput converts the input string into whatever format is needed for the puzzle
// update the return type as needed
func parseInput(input string) []data {
	var fields []data
	var rows []string

	makeField := func(rows []string) data {
		field := data{
			rows: rows,
		}
		var columns []string
		for i := 0; i < len(rows[0]); i++ {
			column := ""
			for _, row := range rows {
				column += string(row[i])
			}
			columns = append(columns, column)
		}
		field.columns = columns
		return field
	}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			field := makeField(rows)
			fields = append(fields, field)
			rows = []string{}
			continue
		}
		rows = append(rows, line)
	}
	// add last one
	field := makeField(rows)
	fields = append(fields, field)

	return fields
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

	solutionPath := filepath.Join(GetPuzzlePath(13, 2023), fmt.Sprintf("solution-%d.txt", puzzle))
	err := WriteFile(solutionPath, []byte(fmt.Sprintf("%d", solution)), true)
	if err != nil {
		log.Fatalf("Error writing solution: %v", err)
	}
	fmt.Println("Solution:", solution)
}
