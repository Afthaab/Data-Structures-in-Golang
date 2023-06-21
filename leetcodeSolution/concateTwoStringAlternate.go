package main

import "fmt"

func main() {
	a := "abcd"
	b := "tv"
	fmt.Println(mergeAlternately(a, b))
}

func mergeAlternately(word1 string, word2 string) string {
	k := 0
	var newstr string

	if len(word2) > len(word1) {
		k = len(word2)
	} else if len(word1) > len(word2) {
		k = len(word1)
	} else {
		k = len(word1)
	}

	for i := 0; i < k; i++ {
		if i <= len(word1)-1 && i <= len(word2)-1 {
			newstr = newstr + string(word1[i]) + string(word2[i])
		} else if i > len(word1)-1 {
			newstr = newstr + string(word2[i])
		} else if i > len(word2)-1 {
			newstr = newstr + string(word1[i])
		}
		fmt.Println(newstr)
		fmt.Println(i)
	}
	return newstr
}
