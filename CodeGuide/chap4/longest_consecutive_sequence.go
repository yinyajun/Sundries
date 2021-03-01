package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 在无序数组中，返回其中最长的连续序列的长度
// 不需要相对顺序

// 暴力方法
// 遍历每个元素，计算每个元素作为连续序列开头的序列长度
// 为了提升查找效率，先将所有元素加入到map中，空间复杂度为N，时间复杂度为N^2

// 从这里可以看出，重复计算。
// 在遍历的过程中，连续序列可能已经部分形成，不再需要每次都从当前元素去一个个找。
func LCCS(arr []int) int {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = 1
	}
	max := 1
	for i := 0; i < len(arr); i++ { // 平方级别复杂度
		ProcessLCCS(arr[i], m, &max)
	}
	fmt.Println(m)
	return max
}

func ProcessLCCS(start int, m map[int]int, max *int) int {
	if _, exist := m[start+1]; exist {
		m[start] = 1 + ProcessLCCS(start+1, m, max)
	}
	*max = utils.MaxInt(m[start], *max)
	return m[start]
}

// map除了做记录元素是否出现过外，还能记录该元素开头的序列长度
func LCCS2(arr []int) int {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = 1
	}
	max := 1
	for i := 0; i < len(arr); i++ { // 平方级别复杂度
		ProcessLCCS2(arr[i], m, &max)
	}
	fmt.Println(m)
	return max
}

// 以start开头的序列长度
// 类似于记忆化递归
func ProcessLCCS2(start int, m map[int]int, max *int) int { // 时间复杂度为O(N)，因为遍历过的中间结果会利用上，不会重复计算
	if _, exist := m[start+1]; !exist {
		return m[start]
	}
	// 有后序
	// start已经更新
	if m[start] != 1 {
		return m[start]
	}
	// 未更新
	m[start] = 1 + ProcessLCCS2(start+1, m, max)
	*max = utils.MaxInt(m[start], *max)
	return m[start]
}

// 虽然时间复杂度是线性的，但是遍历了两次
// 这里map的含义稍有变化，包含有key的序列长度
func LCCS3(arr []int) int {
	m := make(map[int]int)
	max := 1
	for i := 0; i < len(arr); i++ {
		if _, exist := m[arr[i]]; !exist { // 没有遇见过
			m[arr[i]] = 1
			if _, exist := m[arr[i]-1]; exist { // 合并左边和当前
				max = utils.MaxInt(max, LCSSMerge(m, arr[i]-1, arr[i]))
			}
			if _, exist := m[arr[i]+1]; exist { // 合并右边和当前
				max = utils.MaxInt(max, LCSSMerge(m, arr[i], arr[i]+1))
			}
		}
	}
	fmt.Println(m)
	return max
}

//
func LCSSMerge(m map[int]int, less, more int) int {
	left := less - m[less] + 1
	right := more + m[more] - 1
	length := right - left + 1
	m[left] = length
	m[right] = length
	return length
}

func main() {
	arr := []int{5, 100, 4, 200, 1, 3, 2}
	fmt.Println(LCCS(arr))
	fmt.Println(LCCS2(arr))
	fmt.Println(LCCS3(arr))
}
