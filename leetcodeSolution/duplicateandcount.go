package main

import "fmt"

func main() {
	nums := []int{1, 0, 1, 1}
	fmt.Println(containsNearbyDuplicate(nums, 1))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	var count int
	for i := 0; i < len(nums)-1; i++ {
		count = 1
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				if count <= k {
					return true
				}
			} else {
				count++
			}
		}
	}
	return false
}
