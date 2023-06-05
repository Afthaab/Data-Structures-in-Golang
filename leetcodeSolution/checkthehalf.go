package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 1, 7, 11, 5}
	fmt.Println((DoubleCheck(arr)))
}

func DoubleCheck(arr []int) bool {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j {
				if arr[i]/2 == arr[j] && arr[i]%2 == 0 {
					return true
				}
			}
		}
	}

	return false
}
