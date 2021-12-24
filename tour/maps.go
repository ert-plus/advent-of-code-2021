package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wc := make(map[string]int)
	fields := strings.Fields(s)
	for _, f := range fields {
		wc[f] += 1
	}
	return wc
}

func main() {
	wc.Test(WordCount)
}
