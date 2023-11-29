package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"

	. "github.com/stackus/advent-of-code"
)

func main() {
	puzzle := flag.Int("puzzle", 1, "puzzle number: 1 or 2")
	day, year := ParseFlags()

	// check puzzle is valid
	if *puzzle < 1 || *puzzle > 2 {
		log.Fatalf("Invalid puzzle number: %d", *puzzle)
	}

	// read the solution from the solution file
	puzzlePath := GetPuzzlePath(day, year)
	answerPath := filepath.Join(puzzlePath, fmt.Sprintf("solution-%d.txt", *puzzle))
	contents, err := os.ReadFile(answerPath)
	if err != nil {
		log.Fatalf("Error reading solution: %v", err)
	}

	// trim solution of all whitespace
	solution := strings.Trim(string(contents), "\n\t ")

	reply, err := submitSolution(day, year, *puzzle, solution)
	if err != nil {
		log.Fatalf("Error submitting solution: %v", err)
	}

	replyPath := filepath.Join(puzzlePath, fmt.Sprintf("reply-%d.md", *puzzle))
	err = WriteFile(replyPath, reply)
	if err != nil {
		log.Fatalf("Error writing reply: %v", err)
	}

	fmt.Println("Got reply:", string(reply), "\nThis reply has been saved to ", replyPath)

}

func submitSolution(day, year, puzzle int, answer string) ([]byte, error) {
	submit := bytes.NewReader([]byte(fmt.Sprintf("level=%d&answer=%s", puzzle, answer)))
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day)

	body, err := DoPost(url, submit)
	if err != nil {
		return nil, err
	}

	// Parse the page with goquery
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	buf := bytes.Buffer{}

	// Find the nodes
	doc.Find("main > article").Each(func(i int, s *goquery.Selection) {
		s.Children().Each(func(_ int, child *goquery.Selection) {
			buf.WriteString(child.Text() + "\n")
		})
	})

	return buf.Bytes(), nil
}
