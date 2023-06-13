package main

import "fmt"

func main() {
	a := "axc"
	b := "ahbgdc"
	fmt.Println(isSubsequence(a, b))
}

func isSubsequence(s string, t string) bool {
	var i, j int
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}
