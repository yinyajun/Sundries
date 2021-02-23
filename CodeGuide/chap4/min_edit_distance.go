package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 最小编辑代价（加权编辑距离）

// dp[i][j]: s1[...i-1]和s2[...j-1]之间的最小编辑代价
// 多了1位偏移，方便初始化

// dp[i][j] = min(
//		dp[i-1][j] + dc, # delete s1[i-1], 然后s1[...i-2]变换到s2[...j-1]
//      dp[i][j-1] + ic, # s1[...i-1]变换到s2[...j-2]，然后insert s2[j-1]
// 		dp[i-1][j-1] + rc, # s1[...i-1]变换到s2[...j-1]，然后replace s1[j-1]
//      dp[i-1][j-1] # if s1[j-1] == s2[j-1]
// )

// dp[i][0] = i # 将s1全delete
// dp[0][j] = j # 将s2全insert

func MinEditCost(s1, s2 string, ic, dc, rc int) int {
	dp := make([][]int, len(s1)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(s2)+1)
	}

	// init
	for i := 0; i < len(s1)+1; i++ {
		dp[i][0] = i * dc
	}
	for j := 0; j < len(s2)+1; j++ {
		dp[0][j] = j * ic
	}
	// iterate
	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {

			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] + rc
			}
			dp[i][j] = utils.MinInt(
				dp[i-1][j]+dc,
				dp[i][j-1]+ic,
				dp[i][j],
			)
		}
	}
	return dp[len(s1)][len(s2)]
}

// 空间压缩，为了使得空间为min(M,N)
// 需要将长度较短的字符串作为列
// 上述方法中，将s2作为列。
// 如果s2较长，那么将s1和s2对换，同时将ic和dc对调
func MinEditCost2(s1, s2 string, ic, dc, rc int) int {
	// 短的字符串作为列
	if len(s2) < len(s1) {
		s1, s2 = s2, s1
		ic, dc = dc, ic
	}
	dp := make([]int, len(s2)+1)
	// init
	for j := 0; j < len(s2)+1; j++ {
		dp[j] = j * ic // dp[0][j]: 插入s2[...j]
	}
	// iterate
	for i := 1; i < len(s1)+1; i++ {
		// dp[i-1][...]已经结束
		leftUp := dp[0] // dp[i-1][0]
		// 开始dp[i][...]的初始化
		dp[0] = i * dc // dp[i][0]
		for j := 1; j < len(s2)+1; j++ {
			temp := dp[j] // 优先缓存dp[i-1][j]
			if s1[i-1] == s2[j-1] {
				dp[j] = leftUp
			} else {
				dp[j] = leftUp + rc
			}
			leftUp = dp[j] // 将dp[i-1][j]缓存起来
			dp[j] = utils.MinInt(
				temp+dc,    // dp[i-1][j], 这里要当心，因为dp[j]已经更新过了，现在代表dp[i][j]
				dp[j-1]+ic, // dp[i][j-1]
				dp[j])
			leftUp = temp
		}
	}
	return dp[len(s2)]
}

func main() {
	fmt.Println(MinEditCost("abc", "abd", 1, 1, 3))
	fmt.Println(MinEditCost2("abc", "abd", 1, 1, 3))
}
