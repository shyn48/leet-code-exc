func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var result [][]int

	if len(nums) < 4 {
		return result
	}

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		first := nums[i]

		for j := i + 1; j < len(nums)-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}

			second := nums[j]
			l := j + 1
			r := len(nums) - 1

			for l < r {
				sum := first + second + nums[l] + nums[r]

				if sum == target {
					quadruplets := []int{first, second, nums[l], nums[r]}

					result = append(result, quadruplets)

					l++
					r--

					for l < r && nums[l] == nums[l-1] {
						l++
					}

					for l < r && nums[r] == nums[r+1] {
						r--
					}
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}

	return result
}
