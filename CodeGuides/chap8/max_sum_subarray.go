package main

import (
	"CodeGuide/base/utils"
)

// [1,-2,3,5,-2,6,-1] 中 [3,5,-2,6]

// 用o(N^2)时间确定子数组边界[i,j]，然后再花O(N)时间计算子数组的和
// 是 O(N^3)时间复杂度
func maxSumSubArray(nums []int) int {
	var ans int
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += nums[k]
			}
			if sum > ans {
				ans = sum
			}
		}
	}
	return ans
}

// sum(2,4)和sum(2,5)是有重复计算的，不需要从头重新算
// sum(2,5) = sum(2,4) + nums[5]，可以O(1)的时间得到新的区间的和
// 是O(N^2)的时间复杂度
func maxSumSubArray2(nums []int) int {
	var ans int
	var sum int
	for i := 0; i < len(nums); i++ {
		sum = 0
		for j := i; j < len(nums); j++ {
			sum += nums[j] // sum(i, j)
			if sum > ans {
				ans = sum
			}
		}

	}
	return ans
}

// 处理成前缀和数组的形式，这样可以用O(1)的时间得到给定区间的区间和
// 在不断遍历的过程中，相当于确定了区间的右边界，那么怎么寻找到左边界呢？
// max{以a[0]为右边界的最大子数组和，..., 以a[i]为右边界的最大子数组和}
// 在右边界左边的前缀和数组中寻找最小的值即可，通过维护minSum变量可以O(1)的时间获得最小的值
// O(N)的时间复杂度
func maxSumSubArray3(nums []int) int {
	var ans int
	var minSum int // sum(0,-1) = 0， 注意前缀和的哨兵位
	var sum int

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		// 以i为右边界的子区间的最大值
		if sum-minSum > ans {
			ans = sum - minSum
		}
		// update minSum
		if sum < minSum {
			minSum = sum
		}
	}
	return ans
}

// 书上的方法，如果某个区间和已经为负了，那么累加上这部分的区间和只会更小
// 所以当某个区间和为负的时候，直接记为0，从下一个元素开始重新累加。
func maxSumSubArray4(nums []int) int {
	var ans int
	var sum int

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > ans {
			ans = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return ans
}

// 从动态规划的角度考虑这个问题
// dp[i]: 以nums[i]结尾的最大子数组和
// dp[i] = max{dp[i-1] + nums[i], nums[i]}    dp[0]=nums[0]
// 不难看出，这里和书上的解法本质一样
// 因为可以这么等价， dp[i] = max{dp[i-1], 0} + nums[i]
func maxSumSubArray5(nums []int) int {
	var dp int
	var ans int

	for i := 0; i < len(nums); i++ {
		dp = utils.MaxInt(dp+nums[i], nums[i])
		if dp > ans {
			ans = dp
		}
	}
	return ans
}

//func main() {
//	a := []int{1, -2, 3, 5, -2, 6, -1}
//	fmt.Println(maxSumSubArray(a))
//	fmt.Println(maxSumSubArray2(a))
//	fmt.Println(maxSumSubArray3(a))
//	fmt.Println(maxSumSubArray4(a))
//	fmt.Println(maxSumSubArray5(a))
//}
