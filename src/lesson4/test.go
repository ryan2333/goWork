package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 4}
	result := twoSum(nums, 6)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	var index []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				fmt.Println(i, j)
				index = append(index, i, j)
			}
		}
	}
	return index
}
