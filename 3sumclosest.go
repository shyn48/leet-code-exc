func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    
    result := nums[0] + nums[1] + nums[len(nums) - 1]
    
    for i := 0; i < len(nums) - 2; i++ {
        l := i + 1
        r := len(nums) - 1
        
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            
            if sum == target {
                return sum
            }
            
            if sum < target {
                l++
            } else if sum > target {
                r--
            }
            
            if math.Abs(float64(sum-target)) < math.Abs(float64(result-target)) {
                result = sum
            }
        }
    }
    
    return result
}
