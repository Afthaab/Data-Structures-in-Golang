package main

import "fmt"

func main() {
	a := 1
	count := 0
	for a != 0 {
		if a%2 == 0 {
			a = a / 2
		} else {
			a--
		}
		count++
	}
	fmt.Println(count)
}
