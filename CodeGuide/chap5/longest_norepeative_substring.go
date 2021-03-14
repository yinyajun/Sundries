package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 滑动窗口
func LongestNoRepeativeSubstring(a string) {

	// 每次遍历新元素的时候扩展右边界
	// 发现重复元素的时候收缩左边界
	m := make(map[byte]int) // 记录重复元素的位置
	maxLen := 0
	left, right := 0, 0 // [left, right)

	for right < len(a) { // right是下一个要探索的位置
		// 探索right位置，区间变为[left, right]
		// 收缩左边界
		if idx, ok := m[a[right]]; ok { // 发现和a[i]重复元素a[idx]
			for j := left; j <= idx; j++ {
				delete(m, a[j]) // 删除idx之前的重复元素，因为左边界要收缩到idx+1
			}
			left = idx + 1 // 左边界收缩到重复元素的下一个位置
		}
		// 更新窗口
		m[a[right]] = right // 将当前元素记录到map中
		maxLen = utils.MaxInt(maxLen, right-left+1)
		// debug
		fmt.Println(string(a[right]), left, right, right-left+1)
		// 扩张右边界
		right++ // [left, right)
	}
	fmt.Println(maxLen)
}

func LongestNoRepeativeSubstring2(a string) {
	m := make(map[byte]int) // 记录重复元素的位置
	maxLen := 0
	left, right := 0, -1 // [left, right] = [0, -1] is empty

	for right+1 < len(a) { // right +1代表下一个要探索的位置
		// 扩张右边界
		right++ // [left, right]
		// 收缩左边界
		if idx, ok := m[a[right]]; ok { // 发现和a[i]重复元素a[idx]
			for j := left; j <= idx; j++ {
				delete(m, a[j]) // 删除idx之前的重复元素，因为左边界要收缩到idx+1
			}
			left = idx + 1 // 左边界收缩到重复元素的下一个位置
		}
		// 更新窗口内的值
		m[a[right]] = right // 将当前元素记录到map中
		maxLen = utils.MaxInt(maxLen, right-left+1)
		// debug
		fmt.Println(string(a[right]), left, right, right-left+1)
	}
	fmt.Println(maxLen)
}

func LongestNoRepeativeSubstring3(a string) {
	m := make([]int, 1<<8)
	for i := range m {
		m[i] = -1 // 如果没有重复元素的初始值，left=-1， left+1=0
	}
	maxLen := 0

	left, right := 0, 0             // [left, right)
	for ; right < len(a); right++ { // right为下一个要探索的位置
		// [left, right]
		// 收缩左边界
		left = utils.MaxInt(m[a[right]]+1, left) // 要么是上一个重复元素的下一个位置，要么是当前left位置
		// update window
		m[a[right]] = right
		maxLen = utils.MaxInt(maxLen, right-left+1)
		// debug
		fmt.Println(string(a[right]), left, right, maxLen)
	}
	fmt.Println(maxLen)
}

func main() {
	a := "abcd"
	LongestNoRepeativeSubstring(a)
	LongestNoRepeativeSubstring2(a)
	LongestNoRepeativeSubstring3(a)
}
