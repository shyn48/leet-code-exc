func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    
    reversed := 0
    temp := x
    r := 0
    
    for temp != 0 {
        r = temp % 10
        reversed = reversed * 10 + r
        temp = temp / 10
    }
        
    return x == reversed
}

// alternate

func isPalindrome(x int) bool {
    strNum := strconv.Itoa(x)
    reversed := ""
    for i := len(strNum) - 1; i >= 0; i-- {
        reversed += string(strNum[i])
    } 
        
    return strNum == reversed
}

func isPalindrome(x int) bool {
    // optimal solution by leetcode, we reverse half of the number
    
    if x < 0 || (x % 10 == 0 && x != 0) {
        return false
    }
    
    reversed := 0
    
    for x > reversed {
        reversed = reversed * 10 + x % 10
        x = x / 10
    }
    
    // When the length is an odd number, we can get rid of the middle digit by revertedNumber/10
    // For example when the input is 12321, at the end of the while loop we get x = 12, revertedNumber = 123,
    // since the middle digit doesn't matter in palidrome(it will always equal to itself), we can simply get rid of it.
    
    return x == reversed || x == reversed / 10
}


