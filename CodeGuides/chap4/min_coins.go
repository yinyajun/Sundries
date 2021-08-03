package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 完全背包问题
// dp[i][j] = min(dp[i-1][j], dp[i][j-arr[i]] + 1)
func MinCoins(arr []int, aim int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, aim+1)
	}
	// init
	max := 10000000
	for j := 1; j <= aim; j++ { // dp[0][0]=0
		dp[0][j] = max
		if j >= arr[0] && dp[0][j-arr[0]] != max { // j= k * arr[0]
			dp[0][j] = dp[0][j-arr[0]] + 1
		}
	}
	//iterate
	for i := 1; i < n; i++ {
		for j := 1; j <= aim; j++ {
			left := max
			if j >= arr[i] && dp[i][j-arr[i]] != max { // j= k * arr[i]
				left = dp[i][j-arr[i]] + 1
			}
			dp[i][j] = utils.MinInt(left, dp[i-1][j])
		}
	}

	ret := dp[n-1][aim]
	if ret != max {
		return ret
	}
	return -1
}

func MinCoins2(arr []int, aim int) int {
	n := len(arr)
	dp := make([]int, aim+1)

	// init
	max := int(^uint(0) >> 1)
	for j := 1; j <= aim; j++ {
		dp[j] = max
		if j >= arr[0] && dp[j-arr[0]] != max {
			dp[j] = dp[j-arr[0]] + 1
		}
	}
	// iterate
	for i := 1; i < n; i++ {
		for j := 1; j <= aim; j++ {
			left := max
			if j >= arr[i] && dp[j-arr[i]] != max {
				left = dp[j-arr[i]] + 1
			}
			dp[j] = utils.MinInt(dp[j], left)
		}
	}

	if dp[aim] != max {
		return dp[aim]
	}
	return -1
}

func main() {
	fmt.Println(MinCoins([]int{5, 2, 3}, 0))
	fmt.Println(MinCoins2([]int{5, 2, 3}, 0))
}
