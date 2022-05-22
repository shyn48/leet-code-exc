func letterCombinations(digits string) []string {    
    if len(digits) == 0 {
        return []string{}
    }
    
    digitsToChar := []string{
        "",
        "",
        "abc",
        "def",
        "ghi",
        "jkl",
        "mno",
        "pqrs",
        "tuv",
        "wxyz",
    }
    
    result := []string{""}

    
    for _, digit := range digits {
        strDigit := string(digit)
        intDigit, _ := strconv.Atoi(strDigit)
        
        digitCharacters := digitsToChar[intDigit]
        
        var newResults  []string
        for i, _ := range digitCharacters {
            for _, str := range result {
                newResults = append(newResults, string(str) + string(digitCharacters[i]))
            }
        }
        
        result = newResults
    }
    
    return result
}
