package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l%2 == 1 {
		return float64(findKth(nums1, nums2, l/2))
	} else {
		return (float64(findKth(nums1, nums2, l/2)) + float64(findKth(nums1, nums2, l/2-1))) / 2
	}
}

func findKth(nums1 []int, nums2 []int, k int) int {
	if len(nums1) == 0 {
		return nums2[k]
	}
	if len(nums2) == 0 {
		return nums1[k]
	}
	ia, ib := len(nums1)/2, len(nums2)/2
	ma, mb := nums1[ia], nums2[ib]

	if ia+ib < k {
		if ma > mb {
			return findKth(nums1, nums2[ib+1:], k-ib-1)
		} else {
			return findKth(nums1[ia+1:], nums2, k-ia-1)
		}
	} else {
		if ma > mb {
			return findKth(nums1[:ia], nums2, k)
		} else {
			return findKth(nums1, nums2[:ib], k)
		}
	}
}

func main() {
	nums1 := []int{10, 20, 30}
	nums2 := []int{40, 50}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
