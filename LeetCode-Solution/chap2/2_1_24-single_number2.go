/*
Given an array of integers, every element appears three times except for one. Find that single one.
Note: Your algorithm should have a linear runtime complexity. Could you implement it without using
extra memory?


* @Author: Yajun
* @Date:   2021/11/27 20:47
*/

package chap2

import (
	"solution/utils"
	"unsafe"
)

// time: O(n); space: O(1)
// 法1：原理类似与可以删除元素的bloom过滤器（用一个int数组而不是bit数组）
// 区别就是，bloom过滤器中一个数通过多个hash func来得到多个要插入的索引
// 而这里，通过nums[i]的二进制的位1所在的索引
// 建立一个int数组，每个位置加上对应nums[i]的二进制位1索引，如果一个位置是3的倍数就置零，最后值为1的位置能组成single number
func singleNumber2(nums []int) int {
	utils.Assert(len(nums) > 0)
	const w = int(unsafe.Sizeof(nums[0]) * 8)
	var (
		count = make([]int, w)
		res   int
	)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < w; j++ {
			count[j] += (nums[i] >> j) & 1
			count[j] %= 3 // 逢3置0
		}
	}
	// 低位 -> 高位
	for j := 0; j < w; j++ {
		res += count[j] << j
	}
	return res
}

// time: O(n); space: O(1)
// 法2非常巧妙，用二进制模拟四进制
// one代表二进制位1出现1次的索引（mod3）
// two代表二进制位1出现2次的索引（mod3）
// three = one&two, 位1索引代表出现了3次
// 00 -> 0
// 01 -> 1 (one)
// 10 -> 2 (two)
// 11 -> 3 (three)
func singleNumber2B(nums []int) int {
	var (
		one, two, three int
	)

	for _, n := range nums {
		two |= one & n // one & n 是n这个数的出现2次的bit位，然后和原来two的相或
		one ^= n       // one和n同时有1的位置代表出现两次，所以用亦或
		three = one & two
		// 逢3置0
		one &= ^three
		two &= ^three
	}
	return one
}
