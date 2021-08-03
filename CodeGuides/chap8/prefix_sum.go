package main

import "fmt"

// s[i] = sum(a[0...i]), i in [0, N-1]
// s[i] = s[i-1] + a[i], when i>=1
// s[0] = a[0]
func cumsum1(a []int) []int {
	cum := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		if i == 0 {
			cum[i] = a[i]
			continue
		}
		// i > 0
		cum[i] = cum[i-1] + a[i]
	}
	return cum
}

// s[i+1] = sum(a[0...i]), i in [-1, N-1] , i+1 in [0, N]
// s[i+1] = s[i] + a[i]
func cumsum2(a []int) []int {
	cum := make([]int, len(a)+1)
	cum[0] = 0

	for i := 0; i < len(a); i++ {
		cum[i+1] = cum[i] + a[i]
	}
	return cum
}

// s[i] = sum(a[0...i-1]), i in [0, N]
// s[i] = s[i-1] + a[i-1]
func cumsum3(a []int) []int {
	cum := make([]int, len(a)+1)
	cum[0] = 0

	for i := 1; i <= len(a); i++ {
		cum[i] = cum[i-1] + a[i-1]
	}
	return cum
}

// len(cum) == len(a)
// cum[i] = sum(a[0...i]), i in [0, N-1]
func rangeSum1(cum []int, i, j int) int {
	if i == 0 {
		return cum[j]
	}
	// i > 0
	return cum[j] - cum[i-1]
}

// cum[i] = sum(a[0...i-1]), i in [0, N]
func rangeSum2(cum []int, i, j int) int {
	return cum[j+1] - cum[i]
}

//func main() {
//	a := []int{2, 3, 6, 7, 3}
//	i,j := 0,0
//	fmt.Println(cumsum1(a))
//	fmt.Println(cumsum2(a))
//	fmt.Println(cumsum3(a))
//	fmt.Println(rangeSum1(cumsum1(a), i, j))
//	fmt.Println(rangeSum2(cumsum2(a), i, j))
//	fmt.Println(rangeSum2(cumsum3(a), i, j))
//}
