package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wc := make(map[string]int)
	for _, v := range strings.Fields(s) {
		wc[v] = wc[v] + 1
	}
	return wc
}

func main() {
	wc.Test(WordCount)
}
