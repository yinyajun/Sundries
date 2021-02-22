package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 最长公共子串问题
// 很明显，不能采用公共子序列的方式来定义dp，因为子串需要连续，可以采用最长上升子序列的方式，以...结尾的方式下
// dp[i][j]: 以s1[i-1]和s2[j-1]作为公共子串的最后一个字符的情况下，最长的公共子串

// dp[i][j] = 0, if s1[i-1]!=s2[j-1]
//			= dp[i-1][j-1] + 1, otherwise

// base
// dp[...][0] = 0, dp[0][...] = 0

func LCSS(s1, s2 string) string {
	l1, l2 := len(s1), len(s2)
	dp := make([][]int, l1+1)
	for i, _ := range dp {
		dp[i] = make([]int, l2+1)
	}

	// iterate
	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if s1[i-1] != s2[j-1] {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}
	return lcssPath(s1, s2, dp)
}

func lcssPath(s1, s2 string, dp [][]int) string {
	// 在整个dp矩阵中寻找最大值
	max := 0
	end := 0
	for i, _ := range dp {
		tmp := utils.MaxInt(dp[i]...)
		if tmp > max {
			max = tmp
			end = i
		}
	}
	// s1[end-1]==s2[*]
	// [end - (max -1) -1 , end-1]
	return s1[end-max : end]
}

// 这里如何做空间压缩呢？如果按行或者按列压缩，空间为O(min(M,N))
// 考察dp[i][j]仅依赖dp[i-1][j-1]，如果按斜线方向来计算，那么可以将空间进一步压缩
// 这里最麻烦的就是，斜线方向怎么搞？
// 先考察下
// (0,4)
// (0,3), (1,4)
// (0,2), (1,3), (2,4)
// (0,1), (1,2), (2,3), (3,4)
// (0,0), (1,1), (2,2), (3,3), (4,4)
// (1,0), (2,1), (3,2), (4,3)
// ...
// (4,0)

// 以斜线方向视角，每次都是i++,j++，直到i或j撞到边界后停下
// 斜线行的首个元素变化，先缩小col，col到0后，增加row
// 这里就不需要和之前dp设置那样，每个str都多设置1位
func LCSS2(s1, s2 string) string {
	row, col := 0, len(s2)-1 // 斜线初始位置
	max := 0
	end := -1 // s1的lcs的最后一个index
	for row < len(s1) {
		dp := 0
		i, j := row, col
		for i < len(s1) && j < len(s2) {
			if s1[i] == s2[j] {
				dp += 1
				if dp > max {
					max = dp
					end = i
				}
			} else {
				dp = 0
			}
			i++
			j++
		}
		// 更新斜线初始位置
		if col > 0 {
			col--
		} else {
			row++
		}
	}
	// s1[end] == s2[*]
	// [end - (max-1) ,end]
	return s1[end-max+1 : end+1]

}

func main() {
	fmt.Println(LCSS("abcde", "bebcd"))
	fmt.Println(LCSS2("abcde", "bebcd"))
}
