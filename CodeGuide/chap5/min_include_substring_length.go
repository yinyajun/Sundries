package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 暴力法为O(MN)，能否有O(N)的解法？
// 不难想出用滑动窗口
// str1 = "adabbca", str2 = "acb"

// 主要思路就是维护一个由str2中元素组成的map
// map中key为str2中的元素，value为1，代表str1的滑动窗口欠着的元素
// 首先扩张窗口right，逐步还map中欠着的元素
// 当不欠元素的时候，收缩窗口left（由于有可能多还了一些元素），找到正好不欠元素的left
// 此时window大小就是遍历到right的最小窗口
// left++，此时会又欠元素了，需要扩张right来归还元素

func MinIncludingLength(str1, str2 string) int {
	if len(str1) == 0 || len(str1) < len(str2) {
		return 0
	}

	// 通过str2构建所欠元素的map
	owe := make(map[uint8]int)
	for i := range str2 {
		owe[str2[i]] += 1
	}

	left, right := 0, 0 // [left, right) is empty
	match := len(str2)
	minLen := int(^uint(0) >> 1)

	for right < len(str1) {
		// [left, right]
		c := str1[right]
		if _, ok := owe[c]; ok {
			owe[c]-- // 归还所欠元素
			if owe[c] >= 0 {
				match--
			}
		}
		if match == 0 {
			// shrink left
			for left <= right {
				if v, ok := owe[str1[left]]; ok {
					if v == 0 { // 已经不能再借了
						break
					}
					owe[str1[left]]++ // 再次借走元素
				}
				left++
			} // owe[str1[left]] == 0
			minLen = utils.MinInt(minLen, right-left+1)
			//// debug
			//fmt.Println(str1[left : right+1])
			// next left
			owe[str1[left]]++
			left++
			match++
		}
		right++
	}
	fmt.Println(minLen)
	return minLen
}

func MinIncludingLength2(str1, str2 string) int {
	if len(str1) == 0 || len(str1) < len(str2) {
		return 0
	}
	//
	owe := make([]int, 256)
	for i := range str2 {
		owe[str2[i]]++
	}

	left, right := 0, 0 // [left, right)
	match := len(str2)
	minLen := int(^uint(0) >> 1)

	for right < len(str1) {
		// [left, right]
		// update after right extention
		owe[str1[right]]--
		if owe[str1[right]] >= 0 { // >=0的才可能是str2中的
			match--
		}
		if match == 0 {
			// shrink left
			for owe[str1[left]] < 0 {
				owe[str1[left]]++
				left++
			} //owe[str1[left]] = 0
			//debug
			fmt.Println(str1[left : right+1])
			minLen = utils.MinInt(minLen, right-left+1)
			owe[str1[left]]++
			left++
			match++
		}
		right++
	}
	fmt.Println(minLen)
	return minLen
}

func main() {
	MinIncludingLength("adabbca", "acb")
	MinIncludingLength2("adabbca", "acb")
}
