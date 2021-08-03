package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

func FindStringType(a string, k int) {
	utils.Assert(k < len(a))
	i := 0
	res := ""
	for i <= len(a) {
		if isLowerCase(a[i]) {
			res = string(a[i])
			i++
		} else if isUpperCase(a[i]) && i+1 < len(a) {
			res = string(a[i : i+2])
			i += 2
		} else {
			panic("invalid")
		}
		if i > k {
			fmt.Println(res)
			return
		}
	}
}

func FindStringType2(a string, k int) {
	utils.Assert(k < len(a))
	i := 0
	for i <= len(a) {
		if isLowerCase(a[i]) {
			i++
			if i > k {
				fmt.Println(string(a[i-1]))
				return
			}

		} else if isUpperCase(a[i]) && i+1 < len(a) {
			i += 2
			if i > k {
				fmt.Println(string(a[i-2 : i]))
				return
			}
		} else {
			panic("invalid")
		}
	}
}

// 更快的方式是直接去定位第k个元素所在的类型
// 从k-1开始向左统计连续出现的大写字母
// 如果有奇数个，那么[k-1, k]
// 如果有偶数个且k=Upper，那么[k, k+1]
// 如果有偶数个且k=Lower，那么[k]
func FindStringType3(a string, k int) {
	utils.Assert(k < len(a))
	num := 0
	for i := k - 1; i >= 0; i-- {
		if isUpperCase(a[i]) {
			num++
		} else {
			if num%2 == 1 {
				fmt.Println(string(a[k-1 : k+1]))
				return
			} else if isUpperCase(a[k]) {
				fmt.Println(string(a[k : k+2]))
				return
			} else {
				fmt.Println(string(a[k]))
				return
			}
		}
	}

}

func isLowerCase(a byte) bool {
	return a >= 'a' && a <= 'z'
}

func isUpperCase(a byte) bool {
	return a >= 'A' && a <= 'Z'
}

func main() {
	a := "aaABCDEcBCg"
	b := 4
	FindStringType(a, b)
	FindStringType2(a, b)
	FindStringType3(a, b)
}
