package main

import "fmt"

func main() {

	a := []int{11, 11, 20, 30}
	fmt.Println(threeConsecutiveOdds(a))

}
func threeConsecutiveOdds(a []int) bool {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i]%2 != 0 {
			count++
		} else {
			count = 0
		}
		if count == 3 {
			return true
		}
	}
	return false
}
