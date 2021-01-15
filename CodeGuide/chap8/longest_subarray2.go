package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 未排序数组中累加和为给定值的最长子数组问题

// 1. 前缀和，sum数组维护的复杂度为O(1)。如果知道sum[i]和sum[j]，那么可以用O(1)时间求解[i+1...j]之和
// 2. 两数之差，一次遍历的过程中找到两数之差为给定值。通过hash表，空间换时间，用O(1)的时间去查找已经遍历过的

// 时间复杂度O(N)
// 注意边界条件m[0]=-1

func MaxLength(arr []int, target int) int {
	if len(arr) == 0 {
		return 0
	}
	sumMap := make(map[int]int) // 用来记录已经遍历过的sum，key为sum值，value为达到这个sum的最小index
	sum, maxLen := 0, 0
	sumMap[0] = -1 // [0... -1]的和为0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if preIdx, ok := sumMap[sum-target]; ok {
			maxLen = utils.MaxInt(maxLen, i-preIdx) // arr[i]结尾的和为target的子数组
		}
		if _, ok := sumMap[sum]; !ok { // 当前sum未记录到map中
			sumMap[sum] = i
		}
	}
	return maxLen

}

func main() {
	res := MaxLength([]int{-3, 7, 2, 1, 2, 3}, 7)
	fmt.Println(res)
}
