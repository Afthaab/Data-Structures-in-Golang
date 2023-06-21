package main

import "fmt"

func main() {
	s := "afthaaab"
	fmt.Println(largeGroupPositions(s))
}
func largeGroupPositions(s string) [][]int {
	count := 1
	arr := [][]int{}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			count++
		} else {
			if count >= 3 {
				arr = append(arr, []int{i - count + 1, i})
			}
			count = 1
		}
	}
	return arr
}
