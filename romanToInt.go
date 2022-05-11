func romanToInt(s string) int {
    dict := map[string]int{
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }
    
    num := 0
    for i:=0;i<len(s) -1;i++ {
        if dict[string(s[i])]  < dict[string(s[i+1])] {
            num -= dict[string(s[i])]
        } else {
            num += dict[string(s[i])]
        }
    }
    
    num += dict[string(s[len(s) - 1])]
    
    return num
}
