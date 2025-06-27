package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(ex2(os.Args))
}

// echo 1
func Echo1(args []string) string {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}

// echo 2
func Echo2(args []string) string {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

// echo 3
// using the built in strings package is significantly
// more efficient than concatenating (strings are immutable)
func Echo3(args []string) string {
	return strings.Join(args[1:], " ")
}

// echo4 (bonus to compare strings.Builder with strings.Join)
func Echo4(args []string) string {
	var sb strings.Builder
	for _, arg := range args[1:] {
		if i > 0 {
			sb.WriteByte(' ')
		}
		sb.WriteString(arg)
	}
	return sb.String()
}

// Exercise 1.1:
// Modify the echo program to also print os.Args[0]
// the name of the command that invoked it
func ex1(args []string) string {
	binPath := strings.Split(args[0], "/")
	cmdName := binPath[len(binPath)-1]
	return cmdName + ": " + strings.Join(args[1:], " ")
}

// Exercise 1.2:
// Modify the echo program to print the index and value of each of its
// arguments, one per line.
func ex2(args []string) string {
	s := ""
	for i, k := range args[1:] {
		idx := strconv.Itoa(i)
		s += idx + ": " + k + "\n"
	}
	return s
}
