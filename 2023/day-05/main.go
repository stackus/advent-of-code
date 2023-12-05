package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"math"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"

	. "github.com/stackus/advent-of-code"
)

//go:embed input.txt
var input string

// puzzle1 solves the level 1 puzzle
func puzzle1(input string) int64 {
	plan := parseInput(input)

	var lowest int64 = math.MaxInt64

	for _, seed := range plan.seeds {
		lowest = min(lowest, plan.process(seed))
	}

	return lowest
}

// puzzle2 solves the level 2 puzzle
func puzzle2(input string) int64 {
	plan := parseInput(input)

	var lowest int64 = math.MaxInt64
	var values = make(chan int64, len(plan.seeds)/2)
	var wg sync.WaitGroup

	for i := 0; i < len(plan.seeds); i += 2 {
		wg.Add(1)
		go func(seed, rng int64) {
			defer wg.Done()
			var low int64 = math.MaxInt64

			for j := int64(0); j < rng; j++ {
				low = min(low, plan.process(seed+j))
			}
			values <- low

		}(plan.seeds[i], plan.seeds[i+1])
	}

	go func() {
		wg.Wait()
		close(values)
	}()

	for v := range values {
		lowest = min(lowest, v)
	}

	return lowest
}

type seedMap struct {
	dst int64
	src int64
	rng int64
}

type step struct {
	maps []seedMap
}

func (s *step) process(seed int64) int64 {
	for _, mapping := range s.maps {
		if seed >= mapping.src && seed < mapping.src+mapping.rng {
			return mapping.dst + (seed - mapping.src)
		}
	}
	return seed
}

type plan struct {
	seeds []int64
	steps []step
}

func (p *plan) process(seed int64) int64 {
	for _, step := range p.steps {
		seed = step.process(seed)
	}
	return seed
}

// parseInput converts the input string into whatever format is needed for the puzzle
// update the return type as needed
func parseInput(input string) *plan {
	lines := strings.Split(input, "\n")
	seedLine := strings.Split(strings.TrimLeft(lines[0], "seeds: "), " ")

	seeds := make([]int64, 0, len(seedLine))
	for _, s := range seedLine {
		seed, _ := strconv.ParseInt(s, 10, 64)
		seeds = append(seeds, seed)
	}
	plan := &plan{
		seeds: seeds,
	}

	mapRe := regexp.MustCompile(`(\w+)-to-(\w+) map:`)
	numRe := regexp.MustCompile(`(\d+) (\d+) (\d+)`)

	steps := make([]step, 0)
	var currentStep step
	for _, line := range lines[2:] {
		if len(line) == 0 {
			// next step
			steps = append(steps, currentStep)
			continue
		}
		if mapRe.MatchString(line) {
			// new mapping
			currentStep = step{}
			continue
		}
		if numRe.MatchString(line) {
			matches := numRe.FindStringSubmatch(line)
			dst, _ := strconv.ParseInt(matches[1], 10, 64)
			src, _ := strconv.ParseInt(matches[2], 10, 64)
			rng, _ := strconv.ParseInt(matches[3], 10, 64)

			currentStep.maps = append(currentStep.maps, seedMap{
				dst: dst,
				src: src,
				rng: rng,
			})
		}
	}
	steps = append(steps, currentStep)
	plan.steps = steps

	return plan
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

	var solution int64 = 0
	if puzzle == 1 {
		solution = puzzle1(input)
	} else {
		solution = puzzle2(input)
	}

	solutionPath := filepath.Join(GetPuzzlePath(5, 2023), fmt.Sprintf("solution-%d.txt", puzzle))
	err := WriteFile(solutionPath, []byte(fmt.Sprintf("%d", solution)), true)
	if err != nil {
		log.Fatalf("Error writing solution: %v", err)
	}
	fmt.Println("Solution:", solution)
}
