package main

import "fmt"

// 字符串的交错组成
// aim包含且仅包含s1和s2的所有字符，且都保持其相对顺序

// dp[i][j]: s1[...i-1]和s2[...j-1]能否组成aim[...i+j-1]
// dp[i][j] = dp[i-1][j], if aim[i+j-1]==s1[i-1]
//			= dp[i][j-1], if aim[i+j-1]==s2[j-1]
//          = false, aim出现非法字符

// base dp[0][0] = true

func IsCross(s1, s2, aim string) bool {
	dp := make([][]bool, len(s1)+1)
	for i, _ := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}
	//iterate
	for i := 0; i < len(s1)+1; i++ {
		for j := 0; j < len(s2)+1; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
				continue
			}
			if j > 0 && s2[j-1] == aim[i+j-1] {
				dp[i][j] = dp[i][j-1]
			} else if i > 0 && s1[i-1] == aim[i+j-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = false
			}
		}
	}
	fmt.Println(dp)
	return dp[len(s1)][len(s2)]
}

func IsCross2(s1, s2, aim string) bool {
	// 短的作为列
	if len(s2) > len(s1) {
		s1, s2 = s2, s1
	}
	dp := make([]bool, len(s2)+1)
	for i := 0; i < len(s1)+1; i++ {
		for j := 0; j < len(s2)+1; j++ {
			if i == 0 && j == 0 {
				dp[j] = true
				continue
			}
			if i > 0 && aim[i+j-1] == s1[i-1] {
				dp[j] = dp[j] // dp[i-1][j]
			} else if j > 0 && aim[i+j-1] == s2[j-1] {
				dp[j] = dp[j-1] // dp[i][j-1]
			} else {
				dp[j] = false
			}
		}
	}
	return dp[len(s2)]
}

func main() {
	fmt.Println(IsCross("ab", "12", "ab12"))
	fmt.Println(IsCross2("ab", "123", "ab123"))
}
