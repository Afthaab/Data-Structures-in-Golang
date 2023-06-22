package main

func main() {
	k := 103

}
func getNoZeroIntegers(n int) []int {
	k := n - 1
	for hasZero(k) || hasZero(n-k) {
		k--
	}
	return []int{n - k, k}
}

func hasZero(n int) bool {
	for n > 0 {
		if n%10 == 0 {
			return true
		}
		n /= 10
	}
	return false
}
