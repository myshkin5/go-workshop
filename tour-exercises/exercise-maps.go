package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	for _, field := range strings.Fields(s) {
		counts[field]++
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
