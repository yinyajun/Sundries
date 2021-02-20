package main

import "fmt"

// 换钱的方法

// 暴力递归
func ExchangeNum1(arr []int, aim int) int {
	return process(arr, aim, 0)
}

// arr[idx...N]中组成aim的方法数
func process(arr []int, aim int, idx int) int {
	if idx == len(arr) {
		if aim == 0 {
			return 1
		}
		return 0
	}
	res := 0
	for i := 0; arr[idx]*i <= aim; i++ { // 使用arr[idx]
		res += process(arr, aim-i*arr[idx], idx+1) // arr[idx+1...N]中组成aim-i*arr[idx]的方法数，这里有较多重复计算
	}
	return res
}

// 记忆化搜索
func ExchangeNum2(arr []int, aim int) int {
	m := make([][]int, len(arr)+1)
	for i, _ := range m {
		m[i] = make([]int, aim+1)
	}
	return process2(arr, aim, 0, m)
}

func process2(arr []int, aim, index int, m [][]int) int {
	if index == len(arr) {
		return cond(aim == 0, 1, 0)
	}

	mVal := 0
	res := 0
	for i := 0; arr[index]*i <= aim; i++ {
		mVal = m[index+1][aim-arr[index]*i]
		if mVal != 0 { // 0：不存在
			res += cond(mVal == -1, 0, mVal)
		} else {
			res += process2(arr, aim-arr[index]*i, index+1, m)
		}
	}

	m[index][aim] = cond(res == 0, -1, res)
	return res
}

func cond(expr bool, a, b int) int {
	if expr {
		return a
	}
	return b
}

// dp[i][j], 使用arr[0...i-1]情况下，组成j的方法数
// dp[...][0] = 1
// dp[0][1...] = 0
// dp[i][j] = dp[i-1][j] + dp[i-1][j-arr[i-1]] + dp[i-1][j-arr[i-1]*2] + ...
func ExchangeNum3(arr []int, aim int) int {
	dp := make([][]int, len(arr)+1)
	for i, _ := range dp {
		dp[i] = make([]int, aim+1)
	}

	for i := 0; i < len(arr)+1; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < len(arr)+1; i++ {
		for j := 1; j < aim+1; j++ {
			num := 0
			for k := 0; j-k*arr[i-1] >= 0; k++ {
				num += dp[i-1][j-arr[i-1]*k]
			}
			dp[i][j] = num
		}
	}
	return dp[len(arr)][aim]
}

// 和上面一样，dp矩阵少了一行，初始化条件不同了
// dp[i][j]，使用arr[0...i]情况下组成j的方法数
// dp[...][0] = 1
// dp[0][j*arr[0]]=1
func ExchangeNum4(arr []int, aim int) int {
	dp := make([][]int, len(arr))
	for i, _ := range dp {
		dp[i] = make([]int, aim+1)
	}

	for i := 0; i < len(arr); i++ {
		dp[i][0] = 1
	}
	for k := 1; k*arr[0] <= aim; k++ {
		dp[0][k*arr[0]] = 1
	}

	for i := 1; i < len(arr); i++ {
		for j := 1; j <= aim; j++ {
			num := 0
			for k := 0; j-arr[i]*k >= 0; k++ {
				num += dp[i-1][j-k*arr[i]]
			}
			dp[i][j] = num
		}
	}
	return dp[len(arr)-1][aim]
}

// 在3的方法上，空间压缩
// dp[...][0] =1
func ExchangeNum5(arr []int, aim int) int {
	dp := make([]int, aim+1)
	// init
	dp[0] = 1
	//iterate
	for i := 1; i <= len(arr); i++ {
		//dp[0] = 1， 0位置不更新
		for j := aim; j >= 0; j-- {
			num := 0
			for k := 0; j-k*arr[i-1] >= 0; k++ {
				num += dp[j-k*arr[i-1]]
			}
			dp[j] = num
		}
	}
	return dp[aim]
}

// 类似完全背包的低复杂度解法
// dp[i][j] = dp[i-1][j] + dp[i-1][j-arr[i-1]*k] , j-arr[i-1]*k>=0
// dp[i][j-arr[i-1]] = dp[i-1][j-arr[i-1]] + dp[i-1][j-arr[i-1]-arr[i-1]*k]
// dp[i][j] = dp[i-1][j] + dp[i][j-arr[i-1]]
func ExchangeNum6(arr []int, aim int) int {
	dp := make([]int, aim+1)
	// init
	dp[0] = 1
	// iterate
	for i := 1; i <= len(arr); i++ {
		//dp[0] = 1, 没必要，因为0位置不更新
		for j := 1; j <= aim; j++ {
			left := 0
			if j-arr[i-1] >= 0 {
				left = dp[j-arr[i-1]]
			}
			dp[j] = dp[j] + left
		}
	}
	return dp[aim]
}

func main() {
	arr := []int{5, 10, 25, 1}
	aim := 15
	fmt.Println(ExchangeNum1(arr, aim))
	fmt.Println(ExchangeNum2(arr, aim))
	fmt.Println(ExchangeNum3(arr, aim))
	fmt.Println(ExchangeNum4(arr, aim))
	fmt.Println(ExchangeNum5(arr, aim))
	fmt.Println(ExchangeNum6(arr, aim))
}
