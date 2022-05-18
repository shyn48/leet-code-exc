func bubbleSort(array[] int)[]int {
   for i:=0; i< len(array)-1; i++ {
      for j:=0; j < len(array)-i-1; j++ {
         if (array[j] > array[j+1]) {
            array[j], array[j+1] = array[j+1], array[j]
         }
      }
   }
   return array
}

func threeSum(nums []int) [][]int {
    arr := bubbleSort(nums)
    // [-4,-1,-1,0,1,2]
    var result [][]int
    
    for i := 0; i < len(arr); i++ {
        first := nums[i]
        l := i + 1
        r := len(arr) - 1
        
        for l < r {
            sum := first + nums[l] + nums[r]
            
            if sum < 0 {
                l++
            } else if sum > 0 {
                r--
            } else {
                triplet := []int{first, nums[l], nums[r]}
                
                result = append(result, triplet)
                
                for l < r && nums[l] == triplet[1] {
                    l++
                }
                
                for l < r && nums[r] == triplet[2] {
                    r--
                }
            }
        }
        
        for i + 1 < len(nums) && nums[i+1] == nums[i] {
            i++
        }
    }
    
    return result
}
