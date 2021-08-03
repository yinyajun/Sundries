package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 添加最少字符使得字符串整体是回文字符串

// 最少添加几个字符使得整体是回文字符串
// dp[i][j]:str[i...j]需要添加的字符数
// dp[i][j] = dp[i+1][j-1], if str[i]==str[j]
// = min(dp[i][j-1], dp[i+1][j]) +1, otherwise

// 时间复杂度为O(N^2)
func AddLeastPalindrome(a string) {
	n := len(a)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = 0
	}

	// 斜向上遍历
	for row, col := 0, 1; col < n; col++ {
		for i, j := row, col; j < n; {
			if i+1 < n && j-1 >= 0 && a[i] == a[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = utils.MinInt(dp[i][j-1], dp[i+1][j]) + 1
			}
			i++
			j++
		}
	}
	fmt.Println(dp)
	FindPath(dp, a)
}

func FindPath(dp [][]int, a string) {
	n := len(dp)
	length := dp[0][n-1] + len(a)
	ret := make([]byte, length)
	i, j := 0, n-1
	left, right := 0, length-1
	for j >= i {
		if i+1 < n && j >= 0 && dp[i][j] == dp[i+1][j]+1 { //right add a[i]
			ret[right] = a[i]
			ret[left] = a[i]
			i++
		} else if i < n && j-1 >= 0 && dp[i][j] == dp[i][j-1]+1 { // left add a[j]
			ret[left] = a[j]
			ret[right] = a[j]
			j--
		} else {
			ret[left] = a[i]
			ret[right] = a[j]
			i++
			j--
		}
		left++
		right--
	}
	fmt.Println(string(ret))
}

// 已知最长回文子序列strlps的情况下，如何将时间复杂度降低到O(N)
// 类似剥洋葱的方式
// 先找到strlps的最0层，然后将第0层外边的leftPart和rightPart
// leftPart+rightPart逆序 复制到ret的左侧
// rightPart+leftPart逆序 复制到ret的右侧
func GetPalindrome2(a string, strlps string) {
	length := 2*len(a) - len(strlps)
	ret := make([]byte, length)
	leftLps, rightLps := 0, len(strlps)-1
	leftRet, rightRet := 0, length-1
	leftA, rightA := 0, len(a)-1

	for leftLps <= rightLps {
		// 在a中寻找leftPart和rightPart（leftLps、rightLps）
		leftTmp, rightTmp := leftA, rightA
		// find leftPart [leftTmp, leftA)
		for a[leftA] != strlps[leftLps] {
			leftA++
		}
		// find rightPart (rightA, rightTmp]
		for a[rightA] != strlps[rightLps] {
			rightA--
		}
		set(ret, a, leftRet, rightRet, leftTmp, leftA, rightA, rightTmp)
		leftLps++
		rightLps--
		leftRet += leftA - leftTmp + rightTmp - rightA
		rightRet -= leftA - leftTmp + rightTmp - rightA
	}
	fmt.Println(string(ret))
}

// 将[ls, le)作为leftPart,(rs, re]作为rightPart复制到resl和resr上
func set(ret []byte, a string, resl, resr, ls, le, rs, re int) {
	for i := ls; i < le; i++ { // le - ls次
		ret[resl] = a[i]
		ret[resr] = a[i]
		resl++
		resr--
	}
	for i := re; i > rs; i-- {
		ret[resl] = a[i]
		ret[resr] = a[i]
		resl++
		resr--
	}
}

func main() {
	AddLeastPalindrome("ABCDA")
	GetPalindrome2("A1B21C", "121")
}
