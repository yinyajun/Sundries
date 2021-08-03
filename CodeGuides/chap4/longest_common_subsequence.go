package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// str1和str2的最长公共子序列
// 对于两个字符串，通常的dp设计都是
// dp[i][j]为str1[0...i]和str2[0...j]的公共子串

// 为了边界条件的方便，设置如下
// dp[i][j]为 str1[0...i)和str2[0...j)的公共子串，大小为(M+1)(N+1)
// base: dp[0][...] = 0, dp[...][0] = 0
// dp[i][j] = max{
//		dp[i-1][j],
//		dp[i][j-1],
//		dp[i-1][j-1] + 1    (str1[i-1]==str2[j-1])
// }

func LCS1(s1, s2 string) string {
	l1, l2 := len(s1), len(s2)
	dp := make([][]int, l1+1)
	for i, _ := range dp {
		dp[i] = make([]int, l2+1)
	}

	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			dp[i][j] = utils.MaxInt(dp[i-1][j], dp[i][j-1])
			if s1[i-1] == s2[j-1] {
				dp[i][j] = utils.MaxInt(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}
	return lcePath(dp, s1, s2)
}

// 根据dp矩阵获取path，时间复杂度为O(M+N)
func lcePath(dp [][]int, s1, s2 string) string {
	l1 := len(dp) - 1
	utils.Assert(l1 > 0)
	l2 := len(dp[0]) - 1

	res := make([]uint8, dp[l1][l2])
	idx := len(res) - 1

	for idx >= 0 {
		if l1 > 0 && dp[l1][l2] == dp[l1-1][l2] {
			l1--
		} else if l2 > 0 && dp[l1][l2] == dp[l1][l2-1] {
			l2--
		} else {
			res[idx] = s1[l1-1] // s1[l1-1] == s2[l2-1]
			l1--
			l2--
			idx--
		}
	}
	return string(res)
}

func main() {
	fmt.Println(LCS1("1A2C3D4B56", "B1D23CA45B6A"))

}
