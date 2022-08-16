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


// second try

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    
    for i := range strs[0] {
        for j := range strs {
            s := string(strs[j])
            if i == len(s) || s[i] != strs[0][i] {
                return strs[0][:i]
            }
        }
    }
    
    return strs[0]
}
