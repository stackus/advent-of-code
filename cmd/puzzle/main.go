package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"

	. "github.com/stackus/advent-of-code"
)

func main() {
	day, year := ParseFlags()
	adventOfCodePath := MakeDir(day, year)

	// Get the puzzle description for the Advent of Code website for the given day and year
	puzzle, err := getAOCPuzzle(day, year)
	if err != nil {
		log.Fatalf("Error getting puzzle: %v", err)
	}

	// write the puzzle description to a file
	puzzlePath := fmt.Sprintf("%s%spuzzle.md", adventOfCodePath, string(os.PathSeparator))
	err = WriteFile(puzzlePath, puzzle, true)
	if err != nil {
		log.Fatalf("Error writing puzzle: %v", err)
	}

	fmt.Println("Puzzle description written for day", day, "and year", year)
}

func getAOCPuzzle(day, year int) ([]byte, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)

	body, err := DoGet(url)
	if err != nil {
		return nil, err
	}

	// Parse the page with goquery
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	buf := bytes.Buffer{}

	// Find the nodes with the given class and extract text
	doc.Find(".day-desc").Each(func(i int, s *goquery.Selection) {
		s.Children().Each(func(_ int, child *goquery.Selection) {
			buf.WriteString(child.Text() + "\n")
		})
	})

	return buf.Bytes(), nil
}
