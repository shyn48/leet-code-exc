// simple solution similar to leetcode's

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    
    currentPrefix := strs[0]
    for _, str := range strs {
        for strings.Index(str, currentPrefix) != 0 {
            currentPrefix = currentPrefix[:len(currentPrefix) - 1]
            if len(currentPrefix) == 0 {
                return ""
            }
        }
    }
    
    return currentPrefix
}

// alternative (faster) solutions https://leetcode.com/problems/longest-common-prefix/solution/
