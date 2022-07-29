func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func min(a int, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func maxArea(height []int) int {
    l := 0
    r := len(height) - 1
    maxWaterAmount := 0
    
    for l < r {
        maxWaterAmount = max(maxWaterAmount, (r - l) * min(height[l], height[r]))
        
        if height[r] < height[l] {
            r--
        } else {
            l++
        }
    }
    
    return maxWaterAmount
}
