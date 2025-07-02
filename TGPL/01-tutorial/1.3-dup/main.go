package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	Dup4()
}

// Dup1 stdIn only
func Dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for k, v := range counts {
		if v > 1 {
			fmt.Printf("%s: %d\n", k, v)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

// Dup2 stdIn or any number of files
func Dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, fname := range files {
			f, err := os.Open(fname)
			if err != nil {
				fmt.Fprintf(os.Stderr, "DUP: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for k, v := range counts {
		if v > 1 {
			fmt.Printf("%s: %d\n", k, v)
		}
	}
}

// Dup3 files only
func Dup3() {
	counts := make(map[string]int)

	for _, fname := range os.Args[1:] {
		data, err := os.ReadFile(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "DUP: %v", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s: %d\n", line, n)
		}
	}
}
