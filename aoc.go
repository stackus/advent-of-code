package advent_of_code

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func ParseFlags() (int, int) {
	today := time.Now()

	// use ET timezone
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalf("Error loading timezone: %v", err)
	}
	today = today.In(loc)

	day := flag.Int("day", today.Day(), "day of the month")
	year := flag.Int("year", today.Year(), "year")
	flag.Parse()

	// day must be between 1 and 25
	if *day < 1 || *day > 25 {
		log.Fatalf("Invalid day of the month: %d", *day)
	}

	// year must be between 2015 and 2023
	if *year < 2015 || *year > 2023 {
		log.Fatalf("Invalid year: %d", *year)
	}

	return *day, *year
}

func GetPuzzlePath(day, year int) string {
	_, caller, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatalf("Error getting caller")
	}

	return filepath.Join(filepath.Dir(caller), fmt.Sprintf("../../%d/day-%02d", year, day))
}

func MakeDir(day, year int) string {
	_, caller, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatalf("Error getting caller")
	}

	puzzlePath := filepath.Join(filepath.Dir(caller), fmt.Sprintf("../../%d/day-%02d", year, day))

	err := os.MkdirAll(puzzlePath, 0755)
	if err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	return puzzlePath
}

func WriteFile(filename string, contents []byte, allowOverwrite bool) error {
	// ensure the file doesn't already exist OR allowOverwrite is true
	if _, err := os.Stat(filename); os.IsNotExist(err) || allowOverwrite {
		// make the directory if it doesn't exist
		err := os.MkdirAll(filepath.Dir(filename), 0755)
		if err != nil {
			return fmt.Errorf("error creating directory: %w", err)
		}

		err = os.WriteFile(filename, contents, 0644)
		if err != nil {
			return fmt.Errorf("error writing file: %w", err)
		}
	} else {
		return fmt.Errorf("file already exists: %s", filename)
	}

	return nil
}

func DoGet(url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return nil, fmt.Errorf("AOC_SESSION environment variable is not set")
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func DoPost(url string, body io.Reader) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return nil, fmt.Errorf("AOC_SESSION environment variable is not set")
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
