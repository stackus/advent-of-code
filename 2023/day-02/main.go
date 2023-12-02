package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strings"

	. "github.com/stackus/advent-of-code"
)

//go:embed input.txt
var input string

// puzzle1 solves the level 1 puzzle
func puzzle1(input string) int {
	games := parseInput(input)

	maxCubes := map[string]int{
		"green": 13,
		"red":   12,
		"blue":  14,
	}

	sum := 0

	for _, game := range games {
		goodGame := true
		for color, maximum := range maxCubes {
			if game.cubes[color][1] > maximum {
				goodGame = false
				break
			}
		}
		if goodGame {
			sum += game.gameNum
		}
	}

	return sum
}

// puzzle2 solves the level 2 puzzle
func puzzle2(input string) int {
	games := parseInput(input)
	sum := 0

	for _, game := range games {
		sum += game.cubes["green"][1] * game.cubes["red"][1] * game.cubes["blue"][1]
	}

	return sum
}

type gameData struct {
	gameNum int
	cubes   map[string][2]int
}

// parseInput converts the input string into whatever format is needed for the puzzle
// update the return type as needed
func parseInput(input string) (games []gameData) {
	for _, line := range strings.Split(input, "\n") {
		// Game 1: 8 green, 4 red, 4 blue; 1 green, 6 red, 4 blue; 7 red, 4 green, 1 blue; 2 blue, 8 red, 8 green
		gameNum := 0
		fmt.Sscanf(line, "Game %d:", &gameNum)
		line := strings.TrimPrefix(line, fmt.Sprintf("Game %d:", gameNum))
		rounds := map[string][]int{
			"green": make([]int, 0),
			"red":   make([]int, 0),
			"blue":  make([]int, 0),
			"total": make([]int, 0),
		}
		for _, round := range strings.Split(line, ";") {
			total := 0
			for _, cube := range strings.Split(round, ",") {
				cube = strings.TrimSpace(cube)
				if cube == "" {
					continue
				}
				color := ""
				count := 0
				fmt.Sscanf(cube, "%d %s", &count, &color)
				rounds[color] = append(rounds[color], count)
				total += count
			}
			rounds["total"] = append(rounds["total"], total)
		}
		// sort all round data
		for _, round := range rounds {
			sort.Ints(round)
		}

		// ensure each color has at least one count; otherwise set to 0
		for color, round := range rounds {
			if len(round) == 0 {
				rounds[color] = []int{0}
			}
		}

		game := gameData{
			gameNum: gameNum,
			cubes: map[string][2]int{
				"green": {rounds["green"][0], rounds["green"][len(rounds["green"])-1]},
				"red":   {rounds["red"][0], rounds["red"][len(rounds["red"])-1]},
				"blue":  {rounds["blue"][0], rounds["blue"][len(rounds["blue"])-1]},
				"total": {rounds["total"][0], rounds["total"][len(rounds["total"])-1]},
			},
		}

		games = append(games, game)
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

	solution := 0
	if puzzle == 1 {
		solution = puzzle1(input)
	} else {
		solution = puzzle2(input)
	}

	solutionPath := filepath.Join(GetPuzzlePath(2, 2023), fmt.Sprintf("solution-%d.txt", puzzle))
	err := WriteFile(solutionPath, []byte(fmt.Sprintf("%d", solution)), true)
	if err != nil {
		log.Fatalf("Error writing solution: %v", err)
	}
	fmt.Println("Solution: ", solution)
}
