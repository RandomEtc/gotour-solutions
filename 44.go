package main

import (
	"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
	m:= map[string]int{}
	prevw := 0
	for i := 1; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			w := s[prevw:i]
			m[w]++
			prevw = i+1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}