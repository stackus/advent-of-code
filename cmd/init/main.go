package main

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"os"
	"text/template"

	. "github.com/stackus/advent-of-code"
)

//go:embed embeds/*
var fs embed.FS

func main() {
	day, year := ParseFlags()
	adventOfCodePath := MakeDir(day, year)

	t, err := template.ParseFS(fs, "embeds/*")
	if err != nil {
		log.Fatalf("Error parsing embeds directory: %s", err)
	}

	// list of files to create
	files := []string{
		"main.go",
	}
	for _, file := range files {
		// ensure the file doesn't already exist
		filePath := fmt.Sprintf("%s%s%s", adventOfCodePath, string(os.PathSeparator), file)

		buf := bytes.Buffer{}
		err = t.ExecuteTemplate(&buf, file, struct {
			Day  int
			Year int
		}{
			Day:  day,
			Year: year,
		})
		if err != nil {
			log.Fatalf("Error executing template: %v", err)
		}
		err = WriteFile(filePath, buf.Bytes())
		if err != nil {
			log.Fatalf("Error writing file: %v", err)
		}
	}

	fmt.Println("Initialised day", day, "for year", year)
}
