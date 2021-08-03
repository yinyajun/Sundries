package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 01背包问题
// dp[i][j] = min(dp[i-1][j], dp[i-1][j-arr[i]]+1)
func MinCoins21(arr []int, aim int) int {
	n := len(arr)
	max := int(^uint(0) >> 2)
	dp := make([]int, aim+1)

	// init
	for j := 1; j <= aim; j++ {
		dp[j] = max
		if j == arr[0] {
			dp[j] = 1
		}
	}
	// iterate
	for i := 1; i < n; i++ {
		for j := aim; j >= 1; j-- {
			leftup := max
			if j >= arr[i] && dp[j-arr[i]] != max {
				leftup = dp[j-arr[i]] + 1
			}
			dp[j] = utils.MinInt(leftup, dp[j])
		}
	}
	ret := dp[aim]
	if ret != max {
		return ret
	}
	return -1
}

func main() {
	fmt.Println(MinCoins21([]int{5, 2, 5, 3}, 15))
}
