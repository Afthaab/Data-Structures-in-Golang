package main

import "fmt"

func main() {
	s := "gaabbcd"
	fmt.Println(removeDuplicates(s))
}

func removeDuplicates(s string) string {
	if len(s) <= 1 {
		return s
	}

	result := ""
	visited := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		char := s[i]
		if !visited[char] {
			result += string(char)
			visited[char] = true
		}
	}

	return result
}
