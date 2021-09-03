package main

import (
	"CodeGuide/base/utils"
)

// 暴力解法
// O(NlogN)时间
func OneNumber1(n int) int {
	var res int
	for i := 1; i <= n; i++ {
		res += hasOne(i)
	}
	return res
}

// O(logN)时间
func hasOne(num int) int {
	var res int
	for num > 0 {
		if num%10 == 1 {
			res++
		}
		num /= 10
	}
	return res
}

// 规律：每10个，个位提供一个1；每100个，十位提供10个1.。。
// 个位： n/10 + min(1, max(n%10-1 + 1, 0))
// 十位： n/100*10 + min(10, max(n%100-10 + 1, 0))
// 里面的max()限制，最少提供0个，外面的min()限制，最多提供base个
// 得用minmax来限制
// 时间复杂度O(logN)
func OneNumber2(n int) int {
	base := 1
	var res int
	for n > base {
		res += n/(10*base)*base + utils.MinInt(base, utils.MaxInt(n%(10*base)-base+1, 0))
		base *= 10
	}
	return res
}

//func main() {
//	fmt.Println(OneNumber1(151))
//	fmt.Println(OneNumber2(151))
//}
