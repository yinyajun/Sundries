/*
Implement wildcard pattern matching with support for '?' and '*'.
'?' Matches any single character. '*' Matches any sequence of characters (including the empty sequence).
The matching should cover the entire input string (not partial).
The function prototype should be:
bool isMatch(const char *s, const char *p)
Some examples:
isMatch("aa","a") → false
isMatch("aa","aa") → true
isMatch("aaa","aa") → false
isMatch("aa", "*") → true
isMatch("aa", "a*") → true
isMatch("ab", "?*") → true
isMatch("aab", "c*a*b") → false

* @Author: Yajun
* @Date:   2022/1/10 20:45
*/

package chap3

import "fmt"

func isMatch2(s, p string) bool {
	return match2(s, p, 0, 0)
}

// match2(s, p, i, j)  s[i...] and p[j...] match ?
func match2(s, p string, i, j int) bool {
	if j == len(p) {
		return i == len(s)
	}
	// j < len(p)
	if j < len(p) && p[j] == '*' { // pattern以*开头
		for j < len(p) && p[j] == '*' { // 跳过连续的*
			j++
		}
		// j >= len(p) || p[j] != '*'
		if j >= len(p) { // pattern全部都是*
			return true
		}
		// p[j]!= '*'， p[j-1]== '*'
		// 试探这个'*'应该匹配文本串多少个字符，应该保证剩余的文本串和模式串完全匹配（当然应该在文本串为空之前）
		for i < len(s) && !match2(s, p, i, j) { // 剩余的文本串和模式串不能完全匹配，则*多匹配一个文本串字符
			i++
		}
		// i >= len(s) || match2(s, p ,i, j) is true
		return i < len(s)
	} else { // p[j] != '*'
		return (i < len(s)) && (p[j] == s[i] || p[j] == '?') &&
			match2(s, p, i+1, j+1)
	}
}

func isMatch2B(s, p string) bool {
	return match2B(s, p, 0, 0)
}

func match2B(s, p string, i, j int) bool {
	if j == len(p) {
		return i == len(s)
	}
	if j < len(p) && p[j] == '*' {
		// 有两种选择，要么匹配0次，要么匹配1次（当前*可以匹配文本串字符，潜在要求是存在文本串字符，即文本串不为空，否则会陷入无穷调用）
		return match2B(s, p, i, j+1) || (i < len(s) && match2B(s, p, i+1, j))
		//res := match2B(s, p, i, j+1) || match2B(s, p, i+1, j)
	} else {
		return (i < len(s)) && (p[j] == s[i] || p[j] == '?') &&
			match2B(s, p, i+1, j+1)
	}
}

func isMatch2C(s, p string) bool {
	memo := make(map[pair]bool)
	return match2C(s, p, 0, 0, memo)
}

func match2C(s, p string, i, j int, memo map[pair]bool) bool {
	key := pair{i, j}
	if ans, ok := memo[key]; ok {
		return ans
	}

	if j == len(p) {
		return i == len(s)
	}

	if j < len(p) && p[j] == '*' {
		memo[key] = match2C(s, p, i, j+1, memo) ||
			(i < len(s) && match2C(s, p, i+1, j, memo))
	} else {
		memo[key] = (i < len(s)) && (p[j] == s[i] || p[j] == '?') &&
			match2C(s, p, i+1, j+1, memo)
	}

	return memo[key]
}

func isMatch2D(s, p string) bool {
	memo := make(map[pair]bool)
	return match2D(s, p, 0, 0, memo)
}

func match2D(s, p string, i, j int, memo map[pair]bool) bool {
	key := pair{i, j}
	if ans, ok := memo[key]; ok {
		return ans
	}

	if j == len(p) {
		return i == len(s)
	}

	if j < len(p) && p[j] == '*' {
		for j < len(p) && p[j] == '*' { // 连续*合并为1个
			j++
		}
		// j == len(p) || p[j]!='*'
		if j == len(p) {
			memo[key] = true
		} else {
			j = j - 1
			memo[key] = match2D(s, p, i, j+1, memo) ||
				(i < len(s) && match2D(s, p, i+1, j, memo))
		}
	} else {
		memo[key] = (i < len(s)) && (p[j] == s[i] || p[j] == '?') &&
			match2D(s, p, i+1, j+1, memo)
	}
	return memo[key]
}

var r = &Recur{
	cnt:    0,
	indent: "\t",
}

type Recur struct {
	cnt    int
	indent string
}

func (r *Recur) Push(info ...interface{}) {
	for i := 0; i < r.cnt; i++ {
		fmt.Print(r.indent)
	}
	fmt.Println(info...)
	r.cnt++
}

func (r *Recur) Pop(info ...interface{}) {
	r.cnt--
	for i := 0; i < r.cnt; i++ {
		fmt.Print(r.indent)
	}
	fmt.Println(info...)

}

// 顺序解法
// dp[i][j] : s[:i]和p[:j]是否匹配
// dp[i][j] = match && dp[i-1][j-1], if no star
//          = dp[i][j-1] || (i > 0 && dp[i-1][j]), if star
// base: dp[0][0] = true, dp[...][0] = false

func isMatch2E(s, p string) bool {
	var (
		m, n = len(s), len(p)
		dp   = make([][]bool, m+1)
	)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}

	// base
	dp[0][0] = true

	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || (i > 0 && dp[i-1][j])
			} else {
				dp[i][j] = i > 0 && (p[j-1] == s[i-1] || p[j-1] == '?') && dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

func isMatch2E2(s, p string) bool {
	var (
		m, n = len(s), len(p)
		dp   = make([][]bool, m+1)
	)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}

	// base
	for i := 0; i < m+1; i++ {
		dp[i][0] = false
	}
	dp[0][0] = true
	for j := 1; j < n+1; j++ {
		dp[0][j] = dp[0][j-1] && p[j-1] == '*'
	}

	// iterate
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else {
				dp[i][j] = (s[i-1] == p[j-1] || p[j-1] == '?') && dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

// 逆序解法
// dp[i][j] : s[i：]和p[j:]是否匹配
// dp[i][j] = match && dp[i+1][j+1], if no star
//          = dp[i][j+1] || (i < m && dp[i+1][j]), if star
// base: dp[m][n] = true, dp[...][n] = false
func isMatch2F(s, p string) bool {
	var (
		m, n = len(s), len(p)
		dp   = make([][]bool, m+1)
	)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}

	// base
	dp[m][n] = true

	for i := m; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if p[j] == '*' {
				dp[i][j] = dp[i][j+1] || (i < m && dp[i+1][j])
			} else {
				dp[i][j] = i < m && (p[j] == s[i] || p[j] == '?') && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}

func isMatch2F2(s, p string) bool {
	var (
		m, n = len(s), len(p)
		dp   = make([][]bool, m+1)
	)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}

	// base
	//for i:= 0 ; i < m+1; i ++{
	//	dp[i][n] = false
	//}
	dp[m][n] = true
	for j := n - 1; j > 0; j-- {
		dp[m][j] = dp[m][j+1] && p[j] == '*'
	}

	// iterate
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if p[j] == '*' {
				dp[i][j] = dp[i][j+1] || dp[i+1][j]
			} else {
				dp[i][j] = (s[i] == p[j] || p[j] == '?') && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}

// 状态压缩
func isMatch2G(s, p string) bool {
	var (
		m, n = len(s), len(p)
		dp   = [2][]bool{}
	)

	for i := 0; i < 2; i++ {
		dp[i] = make([]bool, n+1)
	}

	for i := m; i >= 0; i-- {
		ii := (m - i) % 2
		ii_1 := (m - i - 1) % 2
		// base !notice
		if i == m {
			dp[ii][n] = true
		} else {
			dp[ii][n] = false
		}

		for j := n - 1; j >= 0; j-- {
			if p[j] == '*' {
				dp[ii][j] = dp[ii][j+1] || (i < m && dp[ii_1][j])
			} else {
				dp[ii][j] = i < m && (p[j] == s[i] || p[j] == '?') && dp[ii_1][j+1]
			}
		}
	}
	return dp[m%2][0]
}

// 假设p的样子为"*u1*u2*u3*"
// 在s中暴力寻找u1，u2 ...
// pattern可能有如下变形
// 1. p形如"*u1*u2*u3*u4", 末尾*后还有字符
// 2. p形如"u1*u2*u3*u4"，开头*之前有字符
func isMatch2H(s, p string) bool {
	charMatch := func(u, v byte) bool {
		return u == v || v == '?'
	}
	allStar := func(str string, left, right int) bool { // str[left:right] is all star?
		for left <= right {
			if str[left] != '*' {
				return false
			}
			left++
		}
		return true
	}

	// case1: 处理p中*之后的字符, 从尾到前要和s匹配
	for len(s) > 0 && len(p) > 0 && p[len(p)-1] != '*' {
		if charMatch(s[len(s)-1], p[len(p)-1]) {
			s = s[:len(s)-1]
			p = p[:len(s)-1]
		} else {
			return false
		}
	}

	// len(s) ==0 || len(p) == 0 || p[len-1] == '*'
	if len(p) == 0 {
		return len(s) == 0
	}

	// 此时case1已经解决
	var (
		sIdx, pIdx = 0, 0   // 当前遍历到的位置
		sRec, pRec = -1, -1 // 在s中寻找到某个u_i的起始位置
	)

	// len(s) ==0 || p[len-1] == '*'
	for sIdx < len(s) && pIdx < len(p) {
		if p[pIdx] == '*' {
			pIdx++
			sRec, pRec = sIdx, pIdx
		} else if charMatch(s[sIdx], p[pIdx]) {
			sIdx++
			pIdx++
		} else { // 不匹配
			// 需要重新寻找u_i, 从s的下一个位置开始
			if sRec < len(s) {
				if sRec == -1 { // 未遇到*前就发生了不匹配
					return false
				}
				sRec++
				sIdx, pIdx = sRec, pRec
			} else {
				return false
			}
		}
	}
	// sIdx >= len(s) || pIdx >= len(p)
	// p未匹配完，剩余的p必须都是*
	// s未匹配完，由于p的最后一个为*，所以不要紧
	return allStar(p, pIdx, len(p)-1)
}

func isMatch2I(s, p string) bool {
	var (
		sIdx, pIdx int
		sRec, pRec int
		star       bool
		m, n       = len(s), len(p)
	)

	charMatch := func(u, v byte) bool {
		return u == v || v == '?'
	}

	for sIdx < m { // 保证遍历完s
		if pIdx < n && p[pIdx] == '*' {
			star = true
			sRec, pRec = sIdx, pIdx
			for pRec < n && p[pRec] == '*' { // skip continuous *
				pRec++
			}
			if pRec == n {
				return true
			}
			// p[pRec] != '*'
			sIdx, pIdx = sRec-1, pRec-1
		} else {
			if pIdx < n && !charMatch(s[sIdx], p[pIdx]) {
				if !star {
					return false
				}
				sRec++
				sIdx, pIdx = sRec-1, pRec-1
			}
		}
		//switch p[pIdx] {
		//case '*':
		//
		//default:
		//
		//}
		sIdx++
		pIdx++
	}
	for pIdx < n && p[pIdx] == '*' {
		pIdx++
	}
	return pIdx == n
}
