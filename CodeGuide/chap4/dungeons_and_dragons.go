package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 不同于之前的顺序算法，这里使用逆序算法更加符合自然含义

// dp[i][j]:从(i,j)出发，到达右下点所需要的最少hp

// dp[i][j] = min(
//		max(dp[i+1][j] - m[i][j], 1) // choose down && keep at least 1 hp in every step
//		max(dp[i][j+1] - m[i][j], 1) // choose right && keep at least 1 hp in every step
//		)

// 特别注意，为了保证任何时刻，都至少有1hp，使用max(1, *)

// 顺序递推，从后面往前算path；逆序递推，从前面往后算path
func FindPath(dp, arr [][]int) []string {
	m, n := len(dp)-1, len(dp[0])-1
	path := make([]string, m+n)

	i, j := 0, 0
	for i <= m && j <= n && i+j < m+n { // i，j不越界且path长度满足
		down := dp[i+1][j]
		//right := dp[i][j+1]

		_down := utils.MaxInt(down-arr[i][j], 1)
		//_right := utils.MaxInt(right-arr[i][j], 1)

		if dp[i][j] == _down {
			path[i+j] = "down"
			i++
		} else {
			path[i+j] = "right"
			j++
		}
	}
	return path
}

func MinHP(arr [][]int) int {
	row := len(arr)
	utils.Assert(row > 0)
	col := len(arr[0])

	dp := make([][]int, row)
	for i, _ := range dp {
		dp[i] = make([]int, col)
	}
	for i := row - 1; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			if i == row-1 && j == col-1 {
				dp[i][j] = utils.MaxInt(1, 1-arr[i][j])
				continue
			}
			right := int(^uint(0) >> 1)
			down := int(^uint(0) >> 1)

			if i < row-1 { // 注意数组越界
				right = utils.MinInt(right, utils.MaxInt(dp[i+1][j]-arr[i][j], 1))
			}
			if j < col-1 { // 注意数组越界
				down = utils.MinInt(utils.MaxInt(dp[i][j+1]-arr[i][j], 1), down)
			}
			dp[i][j] = utils.MinInt(right, down)
		}
	}
	fmt.Println(FindPath(dp, arr))
	return dp[0][0]
}

func MinHp2(arr [][]int) int {
	row := len(arr)
	utils.Assert(row > 0)
	col := len(arr[0])

	dp := make([][]int, row)
	for i, _ := range dp {
		dp[i] = make([]int, col)
	}

	// init
	dp[row-1][col-1] = utils.MaxInt(1, 1-arr[row-1][col-1])
	for j := col - 2; j >= 0; j-- {
		dp[row-1][j] = utils.MaxInt(dp[row-1][j+1]-arr[row-1][j], 1)
	}
	for i := row - 2; i >= 0; i-- {
		dp[i][col-1] = utils.MaxInt(dp[i+1][col-1]-arr[i][col-1], 1)
	}

	// iterate
	for i := row - 2; i >= 0; i-- {
		for j := col - 2; j >= 0; j-- {
			dp[i][j] = utils.MinInt(
				utils.MaxInt(1, dp[i][j+1]-arr[i][j]),
				utils.MaxInt(1, dp[i+1][j]-arr[i][j]),
			)
		}
	}
	fmt.Println(FindPath(dp, arr))
	return dp[0][0]
}

// 空间压缩
func MinHP3(arr [][]int) int {
	row, col := len(arr), len(arr[0])
	long := utils.MaxInt(row, col)
	short := utils.MinInt(row, col)
	dp := make([]int, short)
	for i := long - 1; i >= 0; i-- {
		for j := short - 1; j >= 0; j-- {
			if i == long-1 && j == short-1 {
				dp[short-1] = utils.MaxInt(1, 1-arr[i][j])
				continue
			}
			if i == long-1 {
				dp[j] = utils.MaxInt(dp[j+1]-arr[i][j], 1)
			} else if j == short-1 {
				dp[j] = utils.MaxInt(dp[j]-arr[i][j], 1)
			} else {
				dp[j] = utils.MinInt(
					utils.MaxInt(1, dp[j]-arr[i][j]),   // dp[i+1][j]
					utils.MaxInt(1, dp[j+1]-arr[i][j]), // dp[i][j+1]
				)
			}
		}
	}
	return dp[0]
}

func main() {
	arr := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{0, 30, -5},
	}
	fmt.Println(MinHP(arr))
	fmt.Println(MinHp2(arr))
	fmt.Println(MinHP3(arr))
}
