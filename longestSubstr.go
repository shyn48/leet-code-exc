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


# second try

func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func lengthOfLongestSubstring2(s string) int {
    seen := make(map[string]int)
    
    maxSubStrLen := 0
    lastSubStrStartIndex := 0
    for i := 0; i < len(s);i++ {
        char := string(s[i])
        
        if _, ok := seen[char]; ok {
            lastSubStrStartIndex = max(seen[char] + 1, lastSubStrStartIndex)
        }
        
        seen[char] = i
        maxSubStrLen = max(maxSubStrLen, i - lastSubStrStartIndex + 1)
    }
    
    return maxSubStrLen
}
