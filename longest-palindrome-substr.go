func longestPalindrome(s string) string {
    if len(s) == 1 {
        return s
    }
    
    currentValue := ""
    maxLen := 0
    
    for i := range s {
        l, r := i,i 
        for l >= 0 && r < len(s) && s[l] == s[r] {
            lenStr := r - l + 1
            if lenStr > maxLen {
                currentValue = string(s[l:r+1])
                maxLen = lenStr
            }
            l--
            r++
        }
        
        fmt.Println(l, r, currentValue)

        
        l,r = i, i + 1 
        for l >= 0 && r < len(s) && s[l] == s[r] {
            lenStr := r - l + 1
            if lenStr > maxLen {
                currentValue = string(s[l:r+1])
                maxLen = lenStr
            }
            
            l--
            r++
        }            
      }
    
    return currentValue
}
