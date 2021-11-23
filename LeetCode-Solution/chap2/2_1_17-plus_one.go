/*
Given a number represented as an array of digits, plus one to the number.

* @Author: Yajun
* @Date:   2021/11/22 23:27
*/

package chap2

// time:O(n); space: O(1)
func plusOne(digits []int) []int {
	var (
		carry int
		n     = len(digits)
	)
	carry = 1

	for i := n - 1; i >= 0; i-- {
		digits[i] += carry
		carry = digits[i] / 10
		digits[i] %= 10
	}
	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}
