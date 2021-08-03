package main

import (
	"CodeGuide/base/utils"
	"math"
)

// 最长可整合子数组的长度
// 可整合子数组的定义：一个数组排序后，相邻两个数的差的绝对值都为1

// 一个数组中可整合数组的大小= （max-min+1）无重复数字的时候
// 当数组中没有重复数字，且max-min+1==len
func findLI(a []int) int {
	if len(a) <= 0 {
		return 0
	}

	var max, min int
	var repeat = make(map[int]struct{})
	var length int
	// [i,j]
	for i := 0; i < len(a); i++ {
		max = math.MinInt32
		min = math.MaxInt32
		for j := i; j < len(a); j++ {
			if _, ok := repeat[a[j]]; ok {
				break // 出现重复元素，必然不是可整合数组
			}
			repeat[a[j]] = struct{}{}
			max = utils.MaxInt(max, a[j])
			min = utils.MinInt(min, a[j])
			if max-min+1 == j-i+1 {
				length = utils.MaxInt(length, j-i+1)
			}
		}
		// clear repeat
		repeat = make(map[int]struct{})
	}
	return length
}
