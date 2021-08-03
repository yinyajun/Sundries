package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// dp[i]: 剩余元素中[i...N-1]最小切割次数(str[0...i-1]已经是回文)
// dp[i] = min{
//		1 + dp[i+1], // str[i] valid, [i+1,...N-1] find min cut
//		1 + dp[i+2], // str[i...i+1] valid, [i+2...N-1] find min cut
//      ...
//		1 + dp[N-1], // str[i...N-2] valid, [N-1] find min cut
//      1 + dp[N], // str[i...N-1] valid， min cut = 0 **********
//	}
// 最后一种情况不要忽略
// 也就是说dp的状态应该有N+1中

// dp[i] = min{1+dp[j+1], i<=j<N, str[i...j] is valid}
// 那么更新dp的过程最主要花费的时间都在确定str[i...j]是否是回文上

// 为了快速确定回文，再次使用一个dp数组p
// p[i][j] = (p[i+1][j-1] == true) && str[i] == str[j]
// 初始条件为p[i,i], p[i, i+1]
// 不难看出，这是斜线向上遍历

// 考察到dp的遍历过程，i从N-1到0，j从i到N-1,从下往上，从左往右
// 左下方的依赖已经计算过，可以将p的迭代和dp的迭代放到一起。

func MinPalindromeCut(a string) int {
	dp := make([]int, len(a)+1)
	p := make([][]bool, len(a))
	for i := range p {
		p[i] = make([]bool, len(a))
	}

	// init
	dp[len(a)] = -1

	for i := len(a) - 1; i >= 0; i-- {
		dp[i] = int(^uint(0) >> 1)
		for j := i; j <= len(a)-1; j++ {
			if a[i] == a[j] && (j-i < 2 || p[i+1][j-1]) {
				p[i][j] = true
				dp[i] = utils.MinInt(dp[i], dp[j+1]+1)
			}
		}
	}
	fmt.Println(dp)
	return dp[0]
}

func MinPalindromeCut2(a string) int {
	dp := make([]int, len(a)+1)
	p := make([][]bool, len(a))
	for i := range p {
		p[i] = make([]bool, len(a))
	}

	// init
	dp[len(a)] = -1

	for i := len(a) - 1; i >= 0; i-- {
		dp[i] = int(^uint(0) >> 1)
		for j := i; j < len(a); j++ {
			if a[i] == a[j] && (j-i < 2 || p[i+1][j-1]) {
				p[i][j] = true
				dp[i] = utils.MinInt(dp[i], dp[j+1]+1)
			}
		}
	}
	fmt.Println(dp)
	return dp[0]
}

func main() {
	a := "ACDCDCDAD"
	MinPalindromeCut(a)
	MinPalindromeCut2(a)
}
