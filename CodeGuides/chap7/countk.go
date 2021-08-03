package main

import "fmt"

// 两个k进制的数无进位相加，第i位上的结果为： [a(i) + b(i)] % k
// 那么k个相同的数，无进位相加，第i为上必然是0

// 将val转为32bit的k进制
// 用长除法计算，原理就是先求个位，再求百位....
func GetKSystemFromNum(val, k int) []int {
	res := make([]int, 32)
	index := 0
	for val != 0 {
		remainder := val % k
		quotient := val / k
		val = quotient
		res[index] = remainder
		index++
	}
	return res
}

// 无进位相加（二进制可以用亦或代替）
func NoCarryAdd(e []int, val, k int) {
	valK := GetKSystemFromNum(val, k)
	for i := 0; i < len(e); i++ {
		e[i] = (e[i] + valK[i]) % k
	}
}

func FindOnceNumber(arr []int, k int) int {
	e := make([]int, 32)
	for i := 0; i < len(arr); i++ {
		NoCarryAdd(e, arr[i], k)
		fmt.Println(e)
	}
	return GetNumberFromKSystem(e, k)
}

func GetNumberFromKSystem(e []int, k int) int {
	res := 0
	for i := len(e) - 1; i >= 0; i-- {
		res = res*k + e[i]
	}
	return res
}

func main() {
	a := []int{5, 5, 5, 5, 5, 2, 3, 2, 2, 2, 2}
	fmt.Println(FindOnceNumber(a, 5))
}
