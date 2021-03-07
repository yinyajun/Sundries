package main

import "fmt"

// int8:[0,127] [-128,-1]
// int8(^uint8(0)>>1)   : 127
// -1-int8(^uint8(0)>>1) :-128

func String2Int(a string) int8 {
	if len(a) == 0 {
		return 0
	}
	if !IsValidExpression(a) {
		return 0
	}
	// 负数的绝对值比正数的绝对值大1，统一用负数表示
	pos := true
	i := 0
	var res int8
	if a[0] == '-' {
		pos = false
		i = 1
	}
	for ; i < len(a); i++ {
		res = res*10 + int8('0'-a[i])
		// overflow？
		if res > 0 {
			return 0
		}
	}
	if pos && res == -1-int8(^uint8(0)>>1) {
		return 0
	}
	if pos {
		return -res
	} else {
		return res
	}
}

func IsValidExpression(a string) bool {
	// 0开头？
	if a[0] == '0' && len(a) > 1 {
		return false
	}
	// 字母开头？
	if !((a[0] > '0' && a[0] < '9') || a[0] == '-') {
		return false
	}
	// -开头
	if (a[0] == '-' && len(a) == 1) || (a[0] == '-' && a[1] == '0') {
		return false
	}
	// 验证a[1,n-1]是否都是数字
	for i := 1; i < len(a); i++ {
		if !(a[i] > '0' && a[i] < '9') {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(String2Int("-129"))
}
