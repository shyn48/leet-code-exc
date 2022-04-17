package main

import (
	"math"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	seen := make(map[string]int)
	maxLen := 0
	lastFoundCharIndex := 0
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if _, ok := seen[char]; ok {
			lastFoundCharIndex = int(math.Max(float64(lastFoundCharIndex), float64(seen[char]+1)))
		}
		seen[char] = i
		maxLen = int(math.Max(float64(maxLen), float64(i-lastFoundCharIndex+1)))
	}

	return maxLen
}

func main() {
	println(lengthOfLongestSubstring("tmmzuxt"))
}
