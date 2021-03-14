package main

import (
	"fmt"
	"time"
)

func GetBinaryNum(n int) {
	a := 0
	getBinaryNum(n-1, "1", &a)
	fmt.Println(a)
}

func getBinaryNum(depth int, pre string, num *int) {
	if depth == 0 {
		*num = *num + 1
		return
	}
	if pre == "0" {
		getBinaryNum(depth-1, "1", num)
	} else { // pre =1
		getBinaryNum(depth-1, "1", num)
		getBinaryNum(depth-1, "0", num)
	}
}

func GetBinaryNum2(n int) {
	a := getBinaryNum2(n-1, "1")
	fmt.Println(a)
}

func getBinaryNum2(depth int, pre string) int {
	if depth == 0 {
		return 1
	}
	if pre == "0" {
		return getBinaryNum2(depth-1, "1")
	} else {
		return getBinaryNum2(depth-1, "1") + getBinaryNum2(depth-1, "0")
	}
}

func GetBinaryNum3(n int) {
	a := getBinaryNum3(n-1, "1", make(map[Pair]int))
	fmt.Println(a)
}

type Pair struct {
	depth int
	pre   string
}

// 这里面有大量重复计算
// 比如：101**** 和 111****，星号部分的可能性是一样的，但是在2^N的算法中，它们分别处于不同的分支上，单独计算
func getBinaryNum3(depth int, pre string, memo map[Pair]int) int {
	if depth == 0 {
		return 1
	}
	if pre == "0" {
		if v1, ok := memo[Pair{depth - 1, "1"}]; ok {
			memo[Pair{depth, pre}] = v1
		} else {
			memo[Pair{depth, pre}] = getBinaryNum3(depth-1, "1", memo)
		}
		return memo[Pair{depth, pre}]
	} else {
		var v1, v2 int
		v1, ok1 := memo[Pair{depth - 1, "1"}]
		v2, ok2 := memo[Pair{depth - 1, "0"}]
		if ok1 && ok2 {
			memo[Pair{depth, pre}] = v1 + v2
		} else if ok1 {
			memo[Pair{depth, pre}] = v1 + getBinaryNum3(depth-1, "0", memo)
		} else if ok2 {
			memo[Pair{depth, pre}] = v2 + getBinaryNum3(depth-1, "1", memo)
		} else {
			memo[Pair{depth, pre}] = getBinaryNum3(depth-1, "0", memo) + getBinaryNum3(depth-1, "1", memo)
		}
		return memo[Pair{depth, pre}]
	}
}

// 既然能用记忆化递归，那么用动态规划的角度去考虑
// n=1, 1
// n=2, 2
// n=3, 3
// n=4, 5
// 用逆序的方式考虑，变化的状态有：当前所处的位置i，i位置的值是0是1
// dp[i][0]: 从i+1到N-1的所有可能性，且i=0
// dp[i][1]: 从i到N的所有可能性，且i=1
// dp[i][0] = dp[i+1][1]
// dp[i][1] = dp[i+1][1] + dp[i+1][0]
// 思路和上面其实一致，但是时间复杂度降到了O(2N)
func GetBinaryNum4(n int) {
	ones := make([]int, n)
	zeros := make([]int, n)
	// init
	ones[n-1] = 1
	zeros[n-1] = 1
	// iterate
	for i := n - 2; i >= 0; i-- {
		ones[i] = ones[i+1] + zeros[i+1]
		zeros[i] = ones[i+1]
	}
	fmt.Println(ones[0])
}

// 其实还有更优的解决方案
// dp[i][0] = dp[i+1][1]
// dp[i][1] = dp[i+1][1] + dp[i+1][0]
// 可以看到，dp[i+1][0] = dp[i+2][1]
// 而最终只关注dp[0][1], 所以可以压缩为dp[i][1] = dp[i+1][1] + dp[i+2][1]
// 进一步可以省略掉一个状态，dp[i]:从i+1到N-1且 i =1 的所有可能性
// dp[i] = dp[i+1] + dp[i+2]
// dp[N-1] = 1
// dp[N-2] = 2

func GetBinaryNum5(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp := make([]int, n)
	// init
	dp[n-1] = 1
	dp[n-2] = 2
	// iterate
	for i := n - 3; i >= 0; i-- {
		dp[i] = dp[i+1] + dp[i+2]
	}
	return dp[0]
}

// 这个其实就是斐波那契数列，可以利用矩阵乘法的方式获得O(logN)的解法
// 这里省略

func main() {
	n := 40
	start := time.Now()
	GetBinaryNum(n)
	fmt.Println(time.Since(start))
	fmt.Println("-----------------")

	start = time.Now()
	GetBinaryNum2(n)
	fmt.Println(time.Since(start))
	fmt.Println("-----------------")

	start = time.Now()
	GetBinaryNum3(n)
	fmt.Println(time.Since(start))
	fmt.Println("-----------------")

	start = time.Now()
	GetBinaryNum4(n)
	fmt.Println(time.Since(start))
	fmt.Println("-----------------")

	start = time.Now()
	fmt.Println(GetBinaryNum5(n))
	fmt.Println(time.Since(start))
	fmt.Println("-----------------")
}
