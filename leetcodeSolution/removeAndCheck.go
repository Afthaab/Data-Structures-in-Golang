package main

import "fmt"

func main() {
	a := []int{1, 1}
	fmt.Println(Checkk(a))
}

func Checkk(a []int) bool {
	count := 0
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			if a[i-1] > a[i+1] {
				return false
			}
			for j := i + 1; j < len(a)-1; j++ {
				if a[j] > a[j+1] {
					return false
				}
			}
		} else if a[i] == a[i+1] {
			count++
		}
	}
	if count == 1 && len(a) == 3 {
		return true
	}
	if count > 1 {
		return false
	}
	return true
}
