package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

// Exercise 1.4
// Modify Dup2 to print the names of all files in which each duplicated line occurs.
type LineInfo struct {
	count int
	files []string
}

func Dup4() {
	counts := make(map[string]*LineInfo)
	files := os.Args[1:]
	duplicates := false

	if len(files) == 0 {
		countLines2(os.Stdin, "STDIN", counts)
	} else {
		for _, fname := range files {
			f, err := os.Open(fname)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Dup4 Error: %v\n", err)
				continue
			}
			countLines2(f, fname, counts)
			f.Close()
		}
	}

	for k, v := range counts {
		if v.count > 1 {
			duplicates = true
			fmt.Printf("%s\nOccurred %d times in:\t%s\n\n", k, v.count, strings.Join(v.files, " "))
		}
	}

	if !duplicates {
		fmt.Println("No duplicates found.")
	}
}

func countLines2(r io.Reader, fname string, counts map[string]*LineInfo) {
	input := bufio.NewScanner(r)
	for input.Scan() {
		line := input.Text()

		if counts[line] == nil {
			counts[line] = &LineInfo{}
		}

		if !slices.Contains(counts[line].files, fname) {
			counts[line].files = append(counts[line].files, fname)
		}

		counts[line].count++
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from %s: %v\n")
	}
}
