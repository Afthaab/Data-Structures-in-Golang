package main

import "fmt"

func splitNum(num int) []int {
	var digits []int
	for num > 0 {
		digit := num % 10
		digits = append([]int{digit}, digits...)
		num /= 10
	}
	return digits
}
func isHappy(num int) bool {
	if num == 1 {
		return true
	}
	arr := splitNum(num)
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = arr[i]*arr[i] + sum
	}
	if sum == 1 {
		return true
	} else {
		return isHappy(sum)
	}

}
func main() {
	fmt.Println(isHappy(7))
}
