package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 整型数组，返回排序后的最大差值
// 使用排序实现，时间为O(NlogN)
// 利用桶排序思路，可以O(N)的时间复杂度，O(N)的空间复杂度

// arr中找到min，max。
// 有N个数，准备N+1个桶，max放在N+1桶中
// 桶区间大小为(max-min)/N
// 桶id=(num - min)*N/(max-min)
// maxGap = max{第一对非空桶之间差值，第二对。。。。}
// 如果O(1)时间求得一对非空桶之间差值，需要额外维护每个桶的min和max
func MaxGap(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}

	n := len(arr)
	min, max := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		min = utils.MinInt(min, arr[i])
		max = utils.MaxInt(max, arr[i])
	}
	if max == min {
		return 0
	}

	hasNum := make([]bool, n+1) // 每个桶是否有数字
	mins := make([]int, n+1)
	maxs := make([]int, n+1)

	var bid int
	for i := 0; i < len(arr); i++ {
		bid = (arr[i] - min) * n / (max - min) // 可能上溢
		if hasNum[bid] {
			mins[bid] = utils.MinInt(mins[bid], arr[i])
			maxs[bid] = utils.MaxInt(maxs[bid], arr[i])
		} else {
			mins[bid] = arr[i]
			maxs[bid] = arr[i]
		}
		hasNum[bid] = true
	}

	// 找到第一个非空桶
	var i int
	for ; i <= n; i++ {
		if hasNum[i] == true {
			break
		}
	}
	lastMax := maxs[i]

	// 从下一个桶开始
	var res int
	for i += 1; i <= n; i++ {
		if hasNum[i] == true {
			res = utils.MaxInt(res, mins[i]-lastMax)
			lastMax = maxs[i]
		}

	}
	return res
}

func main() {
	arr := []int{9, 3, 1, 10}
	fmt.Println(MaxGap(arr))
}
