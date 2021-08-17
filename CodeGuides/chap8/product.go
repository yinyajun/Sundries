package main

import "fmt"

func product(arr []int) []int {
	if len(arr) < 2 {
		panic("invalid arr")
	}

	all := 1
	count := 0

	for _, val := range arr {
		if val != 0 {
			all *= val
		} else {
			count++
		}
	}

	res := make([]int, len(arr))

	if count == 0 {
		for i := 0; i < len(arr); i++ {
			res[i] = all / arr[i]
		}
	}

	if count == 1 {
		for i := 0; i < len(arr); i++ {
			if arr[i] == 0 {
				res[i] = all
			}
		}
	}
	// count > 1, res[i] == 0
	return res
}

func product2(arr []int) []int {

}

func main() {
	arr := []int{2, 3, 1, 4}
	fmt.Println(product(arr))
}
