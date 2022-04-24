// such a bs problem
func convert(s string, numRows int) string {
    str := ""
    
    strLen := len(s)
    
    if strLen == 1 {
        return s
    }
    
    i := 0
    
    // ["PA", "APLS", "YI"]
    
    arr := make([]string, numRows)
    for i < strLen {
        for j := 0; j < numRows && i < strLen; j++ {
            arr[j] += string(s[i])
            i+=1
        }
        for j := numRows - 2; j >= 1 && i < strLen; j-- {
            arr[j] += string(s[i])
            i+=1
        }
    }
    
    for _, el := range arr {
        str += el
    }
    
    return str
}
