package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 在每次遍历的时候，确定以nums[i]为右边界的最大乘积，然后记录这些乘积中的最大值即可（这个思路在之前的多道题目中出现）
// max{以a[0]结尾的最大子数组乘积， 以a[1]结尾的最大子数组乘积，... ， 以a[N-1]结尾的最大子数组乘积}
// 这里其实就可以用动态规划来求解

// max[i] = max{max[i-1] * nums[i], min[i-1] * nums[i], nums[i]}
// min[i] = min{min[i-1] * nums[i], max[i-1] * nums[i], nums[i]}
func MaxProdSubAarray(nums []float32) float32 {
	if len(nums) == 0 {
		return 0
	}
	var ans float32
	max, min := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		max = utils.MaxFloat(max*nums[i], min*nums[i], nums[i])
		min = utils.MinFloat(max*nums[i], min*nums[i], nums[i])
		if max > ans {
			ans = max
		}
	}
	return ans
}

func main() {
	nums := []float32{-2.5, 4, 0, 3, 0.5, 8, -1}
	fmt.Println(MaxProdSubAarray(nums))
}
