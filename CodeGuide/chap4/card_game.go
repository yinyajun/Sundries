package main

import "CodeGuide/base/utils"

type Result struct {
	first  int
	second int
}

// 状态变量有三个：i，j, 先手
func CardGame(arr []int) int {
	N := len(arr)
	dp := make([][]Result, N)
	for i := range dp {
		dp[i] = make([]Result, N)
	}

	for i := 0; j < N; j++ {
		for j := i; j < N; i++ {
			// 先手
			dp[i][j].first = utils.MaxInt(arr[i]+dp[i+1][j].second, arr[j]+dp[i][j-1].second)
			// 后手
			dp[i][j].second = utils.MinInt(dp[i+1][j].first, dp[i][j-1].first)
		}

	}

}
