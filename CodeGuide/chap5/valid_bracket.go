package main

import "fmt"

func ValidBracket(a string) bool {
	leftNum := 0
	for i := range a {
		if a[i] != '(' && a[i] != ')' {
			return false
		}
		if a[i] == '(' {
			leftNum++
		} else {
			leftNum--
			if leftNum < 0 {
				return false
			}
		}
	}
	// 注意此时leftNum没有被消耗完也是错误的
	if leftNum > 0 {
		return false
	}
	return true
}

// dp[i] = 以a[i]结尾的最大括号数
// if a[i] == '(', dp[i] = 0
// if a[i] == ')', dp[i] = a[i-dp[i-1]-1]=='(' ? dp[i-1] + 2, 0 // 可以和之前的str[i-1]结尾的后面再增加
// 注意这里容易漏掉一项，之前的步骤将[i-dp[i]-1...i]处理过了，可能这里是一个全新的以'('开头的
// 如果之前也有合法的，那么可以连起来，所以还需要加上dp[i-dp[i]-1-2]
func ValidBracket2(a string) {
	cond := func(expr bool, a, b int) int {
		if expr {
			return a
		}
		return b
	}
	n := len(a)
	dp := make([]int, n)
	for i := range a {
		if a[i] != ')' {
			dp[i] = 0
		} else {
			dp[i] = cond(a[i-dp[i-1]-1] == '(', dp[i-1]+2, 0)
			if dp[i] > 0 && i-dp[i-1]-2 >= 0 {
				dp[i] += dp[i-dp[i-1]-2]
			}
		}
	}
	fmt.Println(dp)
}

func ValidBracket3(a string) {
	n := len(a)
	dp := make([]int, n)
	// 单个元素肯定invalid
	for i := 1; i < n; i++ {
		if a[i] == ')' { // 才有机会valid
			pre := i - dp[i-1] - 1 // i>=1的隐含要求
			if pre >= 0 && a[pre] == '(' {
				dp[i] = dp[i-1] + 2
				if pre > 0 {
					dp[i] += dp[pre-1]
				}
			}
		}
	}
	fmt.Println(dp)
}

func main() {
	fmt.Println(ValidBracket("()"))
	fmt.Println(ValidBracket("()("))
	fmt.Println(ValidBracket("())"))
	fmt.Println(ValidBracket("()a()"))
	ValidBracket2("(()())")
	ValidBracket3("(()())")
}
