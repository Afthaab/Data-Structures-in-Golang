package main

import "fmt"

func main() {
	a := "ab"
	b := "ab"
	fmt.Println(SwapandCheck(a, b))

}
func SwapandCheck(a string, b string) bool {
	if a == b {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				a = a[:i] + a[i+1:]
				b = b[:j] + b[j+1:]
				i--
				j--
				break
			}
		}
	}
	fmt.Println(a, b)
	if a == b {
		return true
	}
	return false
}
