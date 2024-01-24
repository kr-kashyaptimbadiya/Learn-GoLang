package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	// Word Frequency Count
	for _, v := range words {
		if _, ok := m[v]; ok == true {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}