package main

import "fmt"

func SelectionSort(arr [6]int) {
	var temp int
	for i := 0; i < len(arr)-1; i++ { //T[o(n^2)] s[o(1)]
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println(arr)
}
func BubbleSort(arr [6]int) {
	var temp int
	for i := 0; i < len(arr); i++ { //	T[o(n^2)] S[o(1)]
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(arr)
}
func InsertionSort(arr [1]int) {
	for i := 1; i < len(arr); i++ { //T[o(n^2)] s[o(1)]
		for j := i - 1; j >= 0; j-- {
			temp := arr[i]
			if arr[j] > temp {
				arr[j+1] = arr[j]
				arr[j] = temp
			}

		}
	}
	fmt.Println(arr)
}
func merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	merged := make([]int, size, size)
	for k := 0; k < size; k++ {
		switch true {
		case i == len(left):
			merged[k] = right[j]
			j++
		case j == len(right):
			merged[k] = left[i]
			i++
		case left[i] < right[j]:
			merged[k] = left[i]
			i++
		default:
			merged[k] = right[j]
			j++
		}
	}
	return merged
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}

// without using slice property
func mg(arr []int, low int, mid int, high int) {
	temp := make([]int, high-low+1)
	left := low
	right := mid + 1
	k := 0

	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			temp[k] = arr[left]
			left++
		} else {
			temp[k] = arr[right]
			right++
		}
		k++
	}

	for left <= mid {
		temp[k] = arr[left]
		left++
		k++
	}

	for right <= high {
		temp[k] = arr[right]
		right++
		k++
	}

	for i := 0; i < k; i++ {
		arr[low+i] = temp[i]
	}
}

func mgs(arr []int, low int, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	mgs(arr, low, mid)
	mgs(arr, mid+1, high)
	mg(arr, low, mid, high)
}

func main() {
	arr := []int{100, 8, 106, 22, 6, 8}
	// SelectionSort(arr)
	// BubbleSort(arr)
	// InsertionSort(arr)
	fmt.Println(arr)
	mgs(arr, 0, len(arr)-1)
	fmt.Println(arr)

}
