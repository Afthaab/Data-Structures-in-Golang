package main

import "fmt"

func SelectionSort(arr [5]int) {
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
func BubbleSort(arr [5]int) {
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
func main() {
	arr := [1]int{100}
	// SelectionSort(arr)
	// BubbleSort(arr)
	InsertionSort(arr)

}
