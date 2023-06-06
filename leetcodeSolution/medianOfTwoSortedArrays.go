package main

import (
	"fmt"
	"sort"
)

func main() {
	num1 := []int{10, 20, 30}
	num2 := []int{40, 50, 60}
	fmt.Println(findMedianSortedArrays(num1, num2))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := append(nums1, nums2...)
	sort.Ints(merged)
	l := len(merged)
	if l%2 == 0 {
		return float64(merged[l/2])/2 + float64(merged[l/2-1])/2
	} else {
		return float64(merged[(l-1)/2])
	}
}
