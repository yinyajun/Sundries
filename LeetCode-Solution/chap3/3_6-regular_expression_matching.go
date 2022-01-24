/*
Implement regular expression matching with support for '.' and '*'.
'.' Matches any single character. '*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).
The function prototype should be:
bool isMatch(const char *s, const char *p)
Some examples:
isMatch("aa","a") → false
isMatch("aa","aa") → true
isMatch("aaa","aa") → false
isMatch("aa", "a*") → true
isMatch("aa", ".*") → true
isMatch("ab", ".*") → true
isMatch("aab", "c*a*b") → true

* @Author: Yajun
* @Date:   2022/1/10 17:26
*/

package chap3

import "fmt"

// dp[i][j]=> s[i:] 与 p[j:]是否匹配
// target: dp[0][0]
// choose: 匹配一次 or 两次
// base: dp[i][n] = s[i:] == ""
// dp[i][j] = dp[i][j+2] || (match && dp[i+1][j]) , if p[j+1] == "*"
//          = match && dp[i+1][j+1]
func isMatch(s, p string) bool {
	return match(s, p, 0, 0)
}

func match(s, p string, i, j int) bool {
	if j >= len(p) {
		return i == len(s)
	}
	// j < len(p) !note: 此时允许s为空（s为空，不是base条件）
	_match := i < len(s) && (p[j] == '.' || p[j] == s[i])
	hasStar := j+1 < len(p) && p[j+1] == '*'
	if hasStar {
		return match(s, p, i, j+2) || (_match && match(s, p, i+1, j))
	} else {
		return _match && match(s, p, i+1, j+1)
	}
}

func isMatchB(s, p string) bool {
	return matchB(s, p, 0, 0, make(map[pair]bool))
}

func matchB(s, p string, i, j int, memo map[pair]bool) bool {
	key := pair{i, j}
	if ans, ok := memo[key]; ok {
		return ans
	}

	if j >= len(p) {
		return i == len(s)
	}

	_match := i < len(s) && (p[j] == '.' || p[j] == s[i])
	hasStar := j+1 < len(p) && p[j+1] == '*'

	if hasStar {
		memo[key] = matchB(s, p, i, j+2, memo) || (_match && matchB(s, p, i+1, j, memo)) // 类似左递归
	} else {
		memo[key] = _match && match(s, p, i+1, j+1)
	}

	return memo[key]
}

func isMatchC(s, p string) bool {
	return matchC(s, p, 0, 0)
}

func matchC(s, p string, i, j int) bool {
	if j == len(p) {
		return i == len(s)
	}

	if j+1 < len(p) && p[j+1] == '*' {
		for (i < len(s)) && (p[j] == s[i] || p[j] == '.') {
			if matchC(s, p, i, j+2) { // s[i]能匹配，但是看看当前不匹配能否ok
				return true
			}
			i++
		}
		// s[i]不能匹配了
		return matchC(s, p, i, j+2)
	} else {
		return (i < len(s)) && (p[j] == s[i] || p[j] == '.') &&
			matchC(s, p, i+1, j+1)
	}
}

// 逆序解法
// dp[i][j] : s[i:]是否匹配p[j:]
// dp[i][j] = match && dp[i+1][j+1], if not has star
//			= dp[i][j+2] || match && dp[i+1][j], if has star
// base: dp[...][n] = false, dp[m][n] = true

func isMatchD(s, p string) bool {
	// init dp
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}

	match := func(i, j int) bool {
		return i < m && (p[j] == s[i] || p[j] == '.')
	}

	// base
	dp[m][n] = true

	for i := m; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if j+1 < n && p[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || match(i, j) && dp[i+1][j]
			} else {
				dp[i][j] = match(i, j) && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}

// 顺序解法
// dp[i][j] : s[:i] 是否匹配 p[:j]
// dp[i][j] = dp[i][j-2] || match && dp[i-1][j] , if star
//          = match && dp[i-1][j-1], if not star
// base dp[...][0] = false    dp[0][0] = true
func isMatchE(s, p string) bool {
	// init dp
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}

	match := func(i, j int) bool {
		return i >= 1 && (p[j-1] == s[i-1] || p[j-1] == '.')
	}

	// base
	dp[0][0] = true

	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// dp[*][j] = [...j-1]
			if j-1 > 0 && p[j-1] == '*' { // 并没有将a*当做一个整体
				dp[i][j] = dp[i][j-2] || match(i, j-1) && dp[i-1][j]
			} else {
				dp[i][j] = match(i, j) && dp[i-1][j-1]
			}
		}
	}

	for i := 0; i < len(dp); i++ {
		fmt.Println(dp[i])
	}
	return dp[m][n]
}
