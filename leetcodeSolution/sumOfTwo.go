package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	temp := []int{1, 8, 9, 455}
	fmt.Println(TwoSum(temp, 456))
}
