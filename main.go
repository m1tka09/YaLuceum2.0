package main

import (
	// "fmt"
	"sort"
)

func Permutations(input string) []string {
	var result []string
	var backtrack func(start int)
	runes := []rune(input)

	backtrack = func(start int) {
		if start == len(runes) {
			result = append(result, string(runes))
			return
		}

		seen := make(map[rune]bool)
		for i := start; i < len(runes); i++ {
			if seen[runes[i]] {
				continue
			}
			seen[runes[i]] = true

			runes[start], runes[i] = runes[i], runes[start]
			backtrack(start + 1)
			runes[start], runes[i] = runes[i], runes[start]
		}
	}

	backtrack(0)
	sort.Strings(result)

	return result
}
