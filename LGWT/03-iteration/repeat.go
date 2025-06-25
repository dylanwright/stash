package main

import "strings"

const repeatCount = 5

func Repeat(s string) string {
	var result strings.Builder
	for i := 0; i < repeatCount; i++ {
		result.WriteString(s)
	}
	return result.String()
}
