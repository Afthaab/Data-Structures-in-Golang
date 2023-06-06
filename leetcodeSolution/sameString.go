package main

import "fmt"

func main() {
	a := "afthab"
	b := "fthaba"
	fmt.Println(isAnagram(a, b))
}
func isAnagram(a string, b string) bool {
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
	if a == b {
		return true
	}
	return false
}
