package main

import (
	"fmt"
)

// 顺序解法
// dp[i]: 定义为s[...i-1]可以变换的字母组合数
// init: dp[0] = 1
// dp[i] = 0, else
//       = dp[i-1], i>0 && s[i-1] != '0' && dp[i-1]!=0 # s[i-1]可转为单字母
//		 = dp[i-1] + dp[i-2], i>1 && s[i-2]!='0' && s[i-2,i-1]<26 && dp[i-2]!=0 && dp[i-1]!=0 # s[i-2,i-1]可转为单字母
// 可以发现对于此题，顺序解法很麻烦

// dp[i]: 定义为s[...i]可以变换的字母组合数
// dp[i] = 0, if s[i] == '0'
//       = dp[i-1], i>0 && s[i] != '0' && dp[i-1]!=0
//		 = dp[i-1] + dp[i-2], i>1 && s[i-1]!='0' && s[i-1,i]<26 && dp[i-2]!=0 && dp[i-1]!=0
// 这种定义方式有问题, 可以发现s[0]并没有被检测到，边界条件同样也应该可以从递推式中推导出来

func LetterCombinations2(s string) int {
	dp := make([]int, len(s)+1)
	// init
	dp[0] = 1 // s[0...-1]
	// iterate
	for i := 1; i < len(s)+1; i++ {
		if i > 0 && s[i-1] != '0' && dp[i-1] != 0 { // s[...i-2]有dp[i-1]次，s[i-1]可转为单字母
			dp[i] = dp[i-1]
		}
		if i > 1 && s[i-2] != '0' && (s[i-2]-'0')*10+s[i-1]-'0' < 27 && dp[i-2] != 0 { // s[...i-3]有dp[i-2]次，s[i-2,i-1]可转为单字母
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}

// 逆序算法
// dp[i]: s[i...]的组合数
// dp[i] = dp[i+1], s[i] valid
//       = dp[i+1]+dp[i+2], s[i,i+1] valid && s[i] valid
//       = dp[i+2], s[i,i+1] valid
//       = 0, s[i] invalid
// init: dp[n], n=len(s)
func LetterCombinations3(s string) int {
	dp := make([]int, len(s)+1)
	// init
	dp[len(s)] = 1
	// iterate
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
			continue
		}
		dp[i] = dp[i+1] // i+1<=len(s)
		if i+2 <= len(s) && (s[i]-'0')*10+s[i+1]-'0' < 27 {
			dp[i] += dp[i+2]
		}
	}
	a := []string{}
	Path(s, 0, "", &a)
	fmt.Println(a)

	b := []string{}
	Path2(s, 0, &[]uint8{}, &b)
	fmt.Println(b)

	return dp[0]
}

// 空间压缩
func LetterCombinations4(s string) int {
	dp, dpFirst, dpSecond := 0, 1, 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			dp = 0
			dpSecond = dpFirst
			dpFirst = dp
			continue
		}
		dp = dpFirst
		if i+2 <= len(s) && (s[i]-'0')*10+s[i+1]-'0' < 27 {
			dp += dpSecond
		}
		dpSecond = dpFirst
		dpFirst = dp
	}
	return dp
}

func LetterCombinations5(s string) int {
	dp, dpNext, tmp := 0, 1, 0
	if s[len(s)-1] == '0' {
		dp = 0
	} else {
		dp = dpNext
	}

	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == '0' {
			dpNext = dp
			dp = 0
		} else {
			tmp = dp
			if (s[i]-'0')*10+s[i+1]-'0' < 27 {
				dp += dpNext
			}
			dpNext = tmp
		}
	}
	return dp
}

// 对于dp[i] = dp[i+1]+dp[2]这种，只需要两个寄存器就够了
// 和斐波那契数列是一样的做法
// 能不能和斐波那契那样用矩阵乘法的方式压缩到O(logN)的复杂度？不行。递推表达式是有状态转移的
func LetterCombinations6(s string) int {
	dp, dpNext := 0, 1
	if s[len(s)-1] == '0' {
		dp = 0
	} else {
		dp = dpNext
	}

	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == '0' {
			dpNext = dp
			dp = 0
		} else {
			if (s[i]-'0')*10+s[i+1]-'0' < 27 {
				dp, dpNext = dp+dpNext, dp
			} else {
				dpNext = dp
			}
		}
	}
	return dp
}

func Path(s string, i int, path string, ret *[]string) {
	if i == len(s) {
		*ret = append(*ret, path)
		return
	}
	if s[i] == '0' {
		return
	}
	// s[i]!= '0'
	Path(s, i+1, path+string(Number2Letter(s[i])), ret)
	if i+1 <= len(s)-1 && (s[i]-'0')*10+s[i+1]-'0' < 27 {
		Path(s, i+2, path+string(Number2Letter(s[i], s[i+1])), ret)
	}
}

func Path2(s string, i int, path *[]uint8, ret *[]string) {
	if i == len(s) {
		*ret = append(*ret, string(*path))
		return
	}
	if s[i] == '0' {
		return
	}

	// choose1
	*path = append(*path, Number2Letter(s[i]))
	Path2(s, i+1, path, ret)
	// withdraw choose1
	*path = (*path)[:len(*path)-1] //

	//choose 2
	if i+1 <= len(s)-1 && (s[i]-'0')*10+s[i+1]-'0' < 27 {
		*path = append(*path, Number2Letter(s[i], s[i+1]))
		Path2(s, i+2, path, ret)
		// withdraw choose2
		*path = (*path)[:len(*path)-1] // 由于只添加了一位，所以只能回撤一位
	}
}

func Number2Letter(a ...uint8) uint8 {
	if len(a) == 1 {
		return (a[0] - '0') + 'a' - 1
	}
	if len(a) == 2 {
		return 10*(a[0]-'0') + a[1] - '0' + 'a' - 1
	}
	return 0
}

func LetterCombinations(target string) int {
	if len(target) == 0 {
		return 0
	}
	return ProcessLetter(target, 0)
}

// 逆序解法的递归思路
// p[i]:s[...i-1]已经变换，s[i...]还未变换，没有变换的序列变换的组合
// 1. i==len(s), 1
// 2. s[i] == '0', 0
// 3. s[i]\in ['1', '9'], p[i] = p[i+1] （数字字符串，不满足2的一定满足3）
// 4. s[i,i+1] \in ['10', '26'], p[i] += p[i+2]
func ProcessLetter(s string, i int) int {
	if i == len(s) {
		return 1
	}
	if s[i] == '0' {
		return 0
	}
	res := ProcessLetter(s, i+1)                       // s[i] \in ['1', '9']
	if i+1 < len(s) && (s[i]-'0')*10+s[i+1]-'0' < 27 { // s[i, i+1] \in ['10', '26']
		res += ProcessLetter(s, i+2)
	}
	return res
}

func main() {
	a := "11"
	fmt.Println(LetterCombinations3(a))
	fmt.Println(LetterCombinations4(a))
	fmt.Println(LetterCombinations5(a))
	fmt.Println(LetterCombinations6(a))
}
