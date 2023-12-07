package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	. "github.com/stackus/advent-of-code"
)

//go:embed input.txt
var input string

// puzzle1 solves the level 1 puzzle
func puzzle1(input string) int {
	parsed := parseInput(input, false)
	total := 0
	for i, h := range parsed {
		total += h.bid * (i + 1)
	}

	return total
}

// puzzle2 solves the level 2 puzzle
func puzzle2(input string) int {
	parsed := parseInput(input, true)
	total := 0
	for i, h := range parsed {
		total += h.bid * (i + 1)
	}

	return total
}

type hand struct {
	cards    string
	bid      int
	strength int
	jacks    int
}

type hands []hand

var cardStrength = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'j': 1,
}

func (h hands) Len() int { return len(h) }
func (h hands) Less(i, j int) bool {
	if h[i].strength == h[j].strength {
		for k := 0; k < 5; k++ {
			if cardStrength[rune(h[i].cards[k])] == cardStrength[rune(h[j].cards[k])] {
				continue
			}
			return cardStrength[rune(h[i].cards[k])] < cardStrength[rune(h[j].cards[k])]
		}
	}

	return h[i].strength < h[j].strength
}
func (h hands) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// parseInput converts the input string into whatever format is needed for the puzzle
// update the return type as needed
func parseInput(input string, jacksWild bool) hands {
	hs := hands{}

	re := regexp.MustCompile(`[AKQJT23456789]{5} \d+`)
	for _, line := range strings.Split(input, "\n") {
		matches := re.FindAllString(line, -1)
		if len(matches) == 0 {
			continue
		}
		bid, _ := strconv.Atoi(matches[0][6:])
		h := hand{
			cards: matches[0][:5],
			bid:   bid,
		}

		// determine hand strength
		if jacksWild {
			// replace all Js with js
			h.cards = strings.ReplaceAll(h.cards, "J", "j")
		}
		cards := strings.Split(h.cards, "")
		slices.Sort(cards)
		repeats := []int{}
		card := cards[0]
		seen := 0
		jacks := 0
		if jacksWild && card == "j" {
			jacks++
		}
		for i := 1; i < len(cards); i++ {
			if jacksWild && cards[i] == "j" {
				jacks++
				continue
			}
			if cards[i] == card {
				seen++
				continue
			}
			if seen > 0 {
				repeats = append(repeats, seen)
			}
			seen = 0
			card = cards[i]
		}
		if seen > 0 {
			repeats = append(repeats, seen)
		}
		h.jacks = jacks
		switch len(repeats) {
		case 0:
			// lowest
			h.strength = 1
			switch jacks {
			case 1:
				// one pair
				h.strength = 2
			case 2:
				// trips
				h.strength = 4
			case 3:
				// quads
				h.strength = 6
			case 4, 5:
				// quints
				h.strength = 7
			}
		case 1:
			switch repeats[0] {
			case 1:
				// one pair
				h.strength = 2
				switch jacks {
				case 1:
					// trips
					h.strength = 4
				case 2:
					// quads
					h.strength = 6
				case 3:
					// quints
					h.strength = 7
				}
			case 2:
				// trips
				h.strength = 4
				switch jacks {
				case 1:
					// quads
					h.strength = 6
				case 2:
					// quints
					h.strength = 7
				}
			case 3:
				// quads
				h.strength = 6
				switch jacks {
				case 1:
					// quints
					h.strength = 7
				}
			case 4:
				// quints
				h.strength = 7
			}
		case 2:
			if repeats[0] == repeats[1] {
				// two pair
				h.strength = 3
				switch jacks {
				case 1:
					// full house
					h.strength = 5
				}
			} else {
				// full house
				h.strength = 5
			}
		}
		hs = append(hs, h)
	}

	sort.Sort(hs)

	return hs
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

	solutionPath := filepath.Join(GetPuzzlePath(7, 2023), fmt.Sprintf("solution-%d.txt", puzzle))
	err := WriteFile(solutionPath, []byte(fmt.Sprintf("%d", solution)), true)
	if err != nil {
		log.Fatalf("Error writing solution: %v", err)
	}
	fmt.Println("Solution:", solution)
}
