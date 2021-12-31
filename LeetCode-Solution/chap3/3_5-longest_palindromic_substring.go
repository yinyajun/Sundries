/*
Given a string S, find the longest palindromic substring in S. You may assume that the maximum length
of S is 1000, and there exists one unique longest palindromic substring.

* @Author: Yajun
* @Date:   2021/12/28 20:16
*/

package chap3

func longPalindromicSubstring(a string) string {
	return lpsRecur(a, 0, len(a)-1)
}

func lpsRecur(a string, i, j int) string {
	if i > j {
		return ""
	}
	if i == j {
		return string(a[i])
	}
	if a[i] == a[j] && lpsRecur(a, i+1, j-1) == string(a[i+1:j]) {
		return string(a[i : j+1])
	}

	p := lpsRecur(a, i+1, j)
	q := lpsRecur(a, i, j-1)
	k := lpsRecur(a, i+1, j-1)

	ret := p
	if len(q) > len(ret) {
		ret = q
	}
	if len(k) > len(ret) {
		ret = k
	}
	return ret
}

func longPalindromicSubstringB(a string) string {
	n := len(a)
	var ret string

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	// base
	for i := 0; i < n; i++ {
		dp[i][i] = true
		if i+1 < n && a[i] == a[i+1] {
			dp[i][i+1] = true
		}
	}

	// recursive
	for j := 2; j < n; j++ {
		for i := 0; i < n-j; i++ {
			dp[i][i+j] = dp[i+1][(i+j)-1] && a[i] == a[i+j]
		}
	}

	// search result
	for j := n - 1; j >= 0; j-- {
		for i := 0; i < n-j; i++ {
			if dp[i][i+j] {
				ret = a[i : i+j+1]
				return ret
			}
		}
	}
	return ret
}

// 压缩状态的动态规划
func longPalindromicSubstringC(a string) string {
	var (
		n   = len(a)
		dp  = [2][]bool{}
		ret string
	)

	for i := 0; i < 2; i++ {
		dp[i] = make([]bool, n)
	}
	for j := 0; j < n; j++ {
		row := j % 2
		for i := 0; i < n-j; i++ {
			if j == 0 {
				dp[row][i] = true
			} else if j == 1 {
				dp[row][i] = a[i] == a[i+j]
			} else { // j >= 2
				dp[row][i] = dp[(j-2)%2][i+1] && a[i] == a[i+j]
			}
			if dp[row][i] && j+1 > len(ret) {
				ret = a[i : i+j+1]
			}
		}
	}
	return ret
}
