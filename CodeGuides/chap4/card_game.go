package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 状态变量有三个：i，j, 先手
// dp[i][j][first] = max{
//		dp[i+1][j][second] + arr[i], // 拿走arr[i], 剩下arr[i+1,...,j], 从先手变为后手
//      dp[i][j-1][second] + arr[j],
//		}
// dp[i][j][second] =
//	 	if 先手选left， dp[i+1][j][first]
//		else 先手选right，dp[i][j-1][first]

// 依赖下面和左边的状态，同时j>=i才有意义，所以上三角区域才有意义。

// 初始状态就是i==j，也就是对角线
// 可以有两种方式迭代
// 1. 斜线向上
// 2. 从左往右，从下往上
// 时间复杂度和空间复杂度都是平方级别

// 斜线向上迭代
func CardGame(arr []int) int {
	cond := func(express bool, a, b int) int {
		if express {
			return a
		}
		return b
	}
	N := len(arr)
	first := make([][]int, N)
	for i := range first {
		first[i] = make([]int, N)
	}
	second := make([][]int, N)
	for i := range second {
		second[i] = make([]int, N)
	}

	// init
	for i := 0; i < N; i++ {
		first[i][i] = arr[i]
		second[i][i] = 0
	}

	for row, col := 0, 1; col < N; col++ { //斜线初始位置，斜线结束位置，斜线更新条件
		for i, j := row, col; i < N-col; i, j = i+1, j+1 { // 从斜线开始处更新状态
			down := arr[i] + second[i+1][j]
			left := arr[j] + second[i][j-1]
			first[i][j] = utils.MaxInt(left, down)
			if left > down { // 先手选择arr[j]
				second[i][j] = first[i][j-1]
			} else {
				second[i][j] = first[i+1][j]
			}
		}
	}
	return cond(first[0][N-1] > second[0][N-1], first[0][N-1], second[0][N-1])
}

// 从左向右，从下往上迭代
func CardGame2(arr []int) int {
	cond := func(express bool, a, b int) int {
		if express {
			return a
		}
		return b
	}
	N := len(arr)
	first := make([][]int, N)
	for i := range first {
		first[i] = make([]int, N)
	}
	second := make([][]int, N)
	for i := range second {
		second[i] = make([]int, N)
	}

	// iterate
	for j := 0; j < N; j++ { // 从左向右
		// init
		first[j][j] = arr[j]
		second[j][j] = 0
		for i := j - 1; i >= 0; i-- { // 从下往上
			first[i][j] = utils.MaxInt(arr[i]+second[i+1][j], arr[j]+second[i][j-1])
			if arr[i]+second[i+1][j] > arr[j]+second[i][j-1] {
				second[i][j] = first[i+1][j]
			} else {
				second[i][j] = first[i][j-1]
			}
		}
	}
	GetCardPath(first, second, arr)
	return cond(first[0][N-1] > second[0][N-1], first[0][N-1], second[0][N-1])
}

// dp[i][j][first] = max{
//		dp[i+1][j][second] + arr[i], // 拿走arr[i], 剩下arr[i+1,...,j], 从先手变为后手
//      dp[i][j-1][second] + arr[j],
//		}
// dp[i][j][second] =
//	 	if 先手选left， dp[i+1][j][first]
//		else 先手选right，dp[i][j-1][first]
func GetCardPath(first, second [][]int, arr []int) {
	// 从目标开始向出发，回过头来，找到dp的路径
	N := len(arr)
	firstPath := []int{}
	secondPath := []int{}
	// 目标状态
	i, j, isFirst := 0, N-1, true

	for i <= j {
		if first[i][j]-arr[i] == second[i+1][j] {
			if isFirst {
				firstPath = append(firstPath, arr[i]) // A作为先手的选择
			} else {
				secondPath = append(secondPath, arr[i]) // B作为先手的选择
			}
			isFirst = !isFirst
			i++
		} else { // first[i][j]-arr[j] == second[i][j-1]
			if isFirst {
				firstPath = append(firstPath, arr[j])
			} else {
				secondPath = append(secondPath, arr[j])
			}
			isFirst = !isFirst
			j--
		}
	}
	fmt.Println("first", firstPath)
	fmt.Println("second", secondPath)
}

func main() {
	arr := []int{3, 9, 1, 2}
	fmt.Println(CardGame(arr))
	fmt.Println(CardGame2(arr))
}
