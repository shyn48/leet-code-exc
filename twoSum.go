package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indexArr := []int{0, 0}
	for i, el := range nums {
		indexArr[0] = i
		for j, rel := range nums[:i] {
			if el+rel == target {
				indexArr[1] = j
				return indexArr
			}
		}
	}
	return indexArr
}

func hashMapSolution(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, el := range nums {
		foundIdx, exists := hashMap[target-el]
		if exists {
			return []int{foundIdx, i}
		}
		hashMap[el] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
 	fmt.Println(hashMapSolution([]int{3, 2, 4}, 6))
}
