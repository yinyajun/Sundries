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

// f[i][j]
// = "", if i< j
// = a[i], if i==j
// = a[i:j], if f[i+1][j-1]==a[i+1:(j-1)] && a[i]==a[j]
// = max(f[i+1][j-1], f[i][j-1], f[i+1][j])
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

// dp[i][j]
// = true, i <= j
// = a[i]==a[j], j-i  == 1
// = a[i]==a[j] && dp[i+1][j-1], j-i > 1
// 沿着斜线来递推，先初始化主对角线和副对角线，然后沿着对角线向上的斜线开始遍历（上三角）
// 先确定起点，然后确定移动方向
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

// 从左到右，从上到下递推
func longPalindromicSubstringD(a string) string {
	n := len(a)

	// construct dp matrix
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	var (
		start  int
		maxLen int
	)

	// recursive
	for j := 0; j < n; j++ {
		dp[j][j] = true
		for i := 0; i < j; i++ { // [i:j]
			dp[i][j] = (a[i] == a[j]) && (j-i < 2 || dp[i+1][j-1])
			if dp[i][j] && (j-i+1) > maxLen {
				start = i
				maxLen = j - i + 1
			}
		}
	}
	return string(a[start : start+maxLen])
}

// 从对角线斜向上递推
func longPalindromicSubstringE(a string) string {
	n := len(a)

	// construct dp matrix
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	var (
		start  int
		maxLen int
	)

	// recursive
	for j := 0; j < n; j++ {
		dp[j][j] = true
		for i := 0; i+j < n; i++ { // [i : i+j]
			dp[i][i+j] = (a[i] == a[i+j]) && (j < 2 || dp[i+1][i+j-1])
			if dp[i][i+j] && j+1 > maxLen {
				start = i
				maxLen = j + 1
			}
		}
	}
	return string(a[start : start+maxLen])
}

func longPalindromicSubstringF(a string) string {
	var (
		n             = len(a)
		dp            = [2][]bool{}
		start, maxLen int
	)

	dp[0] = make([]bool, n)
	dp[1] = make([]bool, n)

	for j := 0; j < n; j++ {
		jj := j % 2
		dp[jj][j] = true
		for i := 0; i < j; i++ { // [i:j]
			dp[jj][i] = (a[i] == a[j]) && (j-i < 2 || dp[(jj+1)%2][i+1])
			if dp[jj][i] && (j-i+1) > maxLen {
				start = i
				maxLen = j - i + 1
			}
		}
	}
	return string(a[start : start+maxLen])
}

// time: O(n^2) space: O(n^2)
// 记忆化递归
func longPalindromicSubstringG(a string) string {
	return lpsRecurF(a, make(map[pair]string), 0, len(a)-1)
}

type pair struct {
	start, end int
}

func lpsRecurF(a string, memo map[pair]string, i, j int) string {
	key := pair{start: i, end: j}
	if s, ok := memo[key]; ok {
		return s
	}

	if i > j {
		return ""
	}
	if i == j {
		return string(a[i])
	}

	s1 := lpsRecurF(a, memo, i+1, j-1)
	if len(s1)+2 == j-i+1 && a[i] == a[j] {
		memo[key] = a[i : j+1]
		return memo[key]
	}

	s2 := lpsRecurF(a, memo, i+1, j)
	s3 := lpsRecurF(a, memo, i, j-1)

	memo[key] = s1
	if len(s2) > len(memo[key]) {
		memo[key] = s2
	}
	if len(s3) > len(memo[key]) {
		memo[key] = s3
	}
	return memo[key]
}

// time: O(n^2) space: O(n^2)
// 记忆化递归
func longPalindromicSubstringH(a string) string {
	return lpsRecurG(a, make(map[pair]string), 0, len(a)-1)
}

func lpsRecurG(a string, memo map[pair]string, i, j int) (ret string) {
	key := pair{start: i, end: j}
	if s, ok := memo[key]; ok {
		return s
	}

	// 利用闭包函数，捕获ret变量
	defer func() {
		memo[key] = ret
	}()

	if i > j {
		ret = ""
		return
	}
	if i == j {
		ret = string(a[i])
		return
	}

	s1 := lpsRecurG(a, memo, i+1, j-1)
	if len(s1)+2 == j-i+1 && a[i] == a[j] {
		ret = a[i : j+1]
		return
	}

	ret = s1
	s2 := lpsRecurG(a, memo, i+1, j)
	s3 := lpsRecurG(a, memo, i, j-1)
	if len(s2) > len(ret) {
		ret = s2
	}
	if len(s3) > len(ret) {
		ret = s3
	}
	return
}

// brute force: 枚举中点（可能是偶数回文，可能是奇数回文）
func longPalindromicSubstringI(a string) string {
	var (
		n             = len(a)
		start, maxLen int
	)

	for i := 0; i < n; i++ {
		// odd
		for j := 1; j <= n && i-j >= 0 && i+j < n; j++ {
			if a[i-j] == a[i+j] && 2*j+1 > maxLen {
				maxLen = 2*j + 1
				start = i - j
			} else {
				break
			}
		}

		// even
		for j := 0; j <= n && i-j >= 0 && i+j < n; j++ {
			if a[i-j] == a[i+j] && 2*(j+1) > maxLen {
				maxLen = 2 * (j + 1)
				start = i - j
			} else {
				break
			}
		}
	}
	return string(a[start : start+maxLen])
}
