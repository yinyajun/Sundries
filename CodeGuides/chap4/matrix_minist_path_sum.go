package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 矩阵的最小路径之和

// 逆序解法
// dp(i, j) = min{dp(i+1,j), dp(i,j+1)} + m(i,j)
// 边界条件：dp[row-1][col-1] dp[row-1][*] dp[*][col-1]
func MinPathSum(m [][]int) int {
	row, col := len(m), len(m[0])

	dp := make([][]int, row)
	for i, _ := range dp {
		dp[i] = make([]int, col)
	}

	dp[row-1][col-1] = m[row-1][col-1]

	for j := col - 2; j >= 0; j-- {
		dp[row-1][j] += dp[row-1][j+1] + m[row-1][j]
	}

	for i := row - 2; i >= 0; i-- {
		dp[i][col-1] += dp[i+1][col-1] + m[i][col-1]
	}

	for i := row - 2; i >= 0; i-- {
		for j := col - 2; j >= 0; j-- {
			dp[i][j] = utils.MinInt(dp[i+1][j], dp[i][j+1]) + m[i][j]
		}
	}
	return dp[0][0]
}

// 顺序解法
// dp(i,j) = min{dp(i-1,j), dp(i, j-1)} + m(i,j)
// 边界条件：dp[0][0] dp[0][*] dp[*][0]
func MinPathSum2(m [][]int) int {
	row, col := len(m), len(m[0])
	dp := make([][]int, row)
	for i, _ := range dp {
		dp[i] = make([]int, col)
	}

	// init
	dp[0][0] = m[0][0]
	for j := 1; j < col; j++ {
		dp[0][j] = dp[0][j-1] + m[0][j]
	}
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + m[i][0]
	}
	// iterate
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = utils.MinInt(dp[i-1][j], dp[i][j-1]) + m[i][j]
		}
	}
	return dp[row-1][col-1]
}

// 顺序解法（空间压缩）
// dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + m[i][j]
func MinPathSum3(m [][]int) int {
	row, col := len(m), len(m[0])
	dp := make([]int, col)

	// init
	dp[0] = m[0][0]
	for j := 1; j < col; j++ {
		dp[j] = dp[j-1] + m[0][j]
	}

	// iterate
	for i := 1; i < row; i++ {
		dp[0] = dp[0] + m[i][0]
		for j := 1; j < col; j++ {
			dp[j] = utils.MinInt(dp[j], dp[j-1]) + m[i][j]
		}
	}
	return dp[col-1]
}

func MinPathSum4(m [][]int) int {
	row, col := len(m), len(m[0])
	dp := make([]int, row)

	// init
	dp[0] = m[0][0]
	for i := 1; i < row; i++ {
		dp[i] = dp[i-1] + m[i][0]
	}

	// iterate
	for j := 1; j < col; j++ {
		dp[0] = dp[0] + m[0][j]
		for i := 1; i < row; i++ {
			dp[i] = utils.MinInt(dp[i], dp[i-1]) + m[i][j]
		}
	}
	return dp[row-1]
}

func MinPathSum5(m [][]int) int {
	row, col := len(m), len(m[0])
	more := utils.MaxInt(row, col)
	less := utils.MinInt(row, col)
	rowMore := more == row

	dp := make([]int, less)
	dp[0] = m[0][0]
	for i := 1; i < less; i++ {
		if rowMore {
			dp[i] += dp[i-1] + m[0][i]
		} else {
			dp[i] += dp[i-1] + m[i][0]
		}
	}
	for i := 1; i < more; i++ {
		if rowMore {
			dp[0] += m[i][0]
		} else {
			dp[0] += m[0][i]
		}
		for j := 1; j < less; j++ {
			if rowMore {
				dp[j] = utils.MinInt(dp[j], dp[j-1]) + m[i][j]
			} else {
				dp[j] = utils.MinInt(dp[j], dp[j-1]) + m[j][i]
			}
		}
	}
	return dp[less-1]
}

func main() {
	m := [][]int{
		{1, 3, 5, 9},
		{8, 1, 3, 4},
		{5, 0, 6, 1},
		{8, 8, 4, 0},
	}
	fmt.Println(MinPathSum(m))
	fmt.Println(MinPathSum2(m))
	fmt.Println(MinPathSum3(m))
	fmt.Println(MinPathSum4(m))
	fmt.Println(MinPathSum5(m))
}
