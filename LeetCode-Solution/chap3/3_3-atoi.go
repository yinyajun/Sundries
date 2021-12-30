/*
Implement atoi to convert a string to an integer.
Hint: Carefully consider all possible input cases. If you want a challenge, please do not see below and
ask yourself what are the possible input cases.
Notes: It is intended for this problem to be specified vaguely (ie, no given input specs). You are responsible to gather all the input requirements up front.
Requirements for atoi:
The function first discards as many whitespace characters as necessary until the first non-whitespace
character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by
as many numerical digits as possible, and interprets them as a numerical value.
The string can contain additional characters after those that form the integral number, which are ignored
and have no effect on the behavior of this function.
If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such
sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.
If no valid conversion could be performed, a zero value is returned. If the correct value is out of the
range of representable values, INT_MAX (2147483647) or INT_MIN (-2147483648) is returned.


* @Author: Yajun
* @Date:   2021/12/20 10:09
*/

package chap3

import (
	"math"
	"solution/utils"
)

// 1. 去除空白字符
// 2. 标记正负符号
// 3. 数值范围

func atoi(str string) int {
	var (
		sign = 1
		base int
		i    int
	)

	// remove white space
	for i < len(str) && str[i] == ' ' {
		i++
	} // i > n || str[i]!=' '

	// determine sign
	if i < len(str) && (str[i] == '+' || str[i] == '-') {
		sign = utils.If(str[i] == '+', 1, -1).(int)
		i++ // ! note
	}

	// to int
	for ; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			break
		}
		if overflow(base, str[i]) {
			return utils.If(sign > 0, math.MaxInt32, math.MinInt32).(int)
		}
		base = 10*base + int(str[i]-'0')
	}
	return base * sign
}

func overflow(base int, next uint8) bool {
	if base > math.MaxInt32/10 {
		return true
	}
	if base == math.MaxInt32/10 && next-'0' > math.MaxInt32%(math.MaxInt32/10) {
		return true
	}
	return false
}
