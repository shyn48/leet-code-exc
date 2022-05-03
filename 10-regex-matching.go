func isMatch(s string, p string) bool {
    if len(p) == 0 || len(s) == 0 {
        return false
    }
    
    dp := make([][]bool, len(s) + 1)
    for i := range dp {
        dp[i] = make([]bool, len(p) + 1)
    }
    
    dp[0][0] = true
    
    for i := 1; i < len(p) + 1; i++ {
        if string(p[i - 1]) == "*" {
            dp[0][i] = dp[0][i - 2]
        }
    }
    
    for i := 1; i < len(s) + 1; i++ {
        for j := 1; j < len(p) + 1;j++ {
            if string(p[j -1]) == "." || p[j - 1] == s[i - 1] {
                dp[i][j] = dp[i-1][j-1]
            } else if string(p[j-1]) == "*" {
                dp[i][j] = dp[i][j-2]
                if p[j-2] == '.' || p[j-2] == s[i-1] {
                    dp[i][j] = dp[i][j] || dp[i-1][j]
                }
            } else {
                dp[i][j] = false 
            }
        }
    }
    
    return dp[len(s)][len(p)]
}
