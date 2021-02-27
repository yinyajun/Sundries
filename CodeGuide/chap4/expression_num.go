package main

import "fmt"

// 表达式长度为奇数
// 偶数位为0或1
// 奇数位为符号
func ValidExpression(a string) bool {
	length := len(a)
	if length%2 == 0 {
		return false
	}
	for i := 0; i < length; i++ {
		if i%2 == 0 && a[i] != '0' && a[i] != '1' {
			return false
		}
		if i%2 == 1 && a[i] != '&' && a[i] != '|' && a[i] != '^' {
			return false
		}
	}
	return true
}

// 将express根据每一个逻辑符号划分为左右两个部分

// split=="^", desired="true"
// num = left为true的num * right为false的num + left为false的num * right为true的num

// split=="^", desired="false"
// num = left为true的num * right为true的num + left为false的num * right为false的num

// split=="|", desired="true"
// num = left为true的num * right为false的num + left为false的num * right为true的num + left为true的num * right为true的num

// split=="|", desired="false"
// num = left为false的num * right为false的num

// split=="&", desired="true"
// num = left为true的num * right为true的num

// split=="&", desired="false"
// num = left为true的num * right为false的num + left为false的num * right为true的num + left为false的num * right为false的num

func ExpressionNum(express string, desired bool) int {
	if !ValidExpression(express) {
		return 0
	}
	return expressionNum(express, desired, 0, len(express)-1)
}

// 时间复杂度：N!
func expressionNum(express string, desired bool, lo, hi int) int {
	if lo == hi {
		if (desired && express[lo] == '1') || (!desired && express[lo] == '0') {
			return 1
		} else {
			return 0
		}
	}

	num := 0
	for i := lo + 1; i <= hi-1; i += 2 {
		symbol := express[i]
		if desired {
			switch symbol {
			case '^':
				num += expressionNum(express, desired, lo, i-1)*expressionNum(express, !desired, i+1, hi) +
					expressionNum(express, !desired, lo, i-1)*expressionNum(express, desired, i+1, hi)
			case '&':
				num += expressionNum(express, desired, lo, i-1) * expressionNum(express, desired, i+1, hi)
			case '|':
				num += expressionNum(express, desired, lo, i-1)*expressionNum(express, desired, i+1, hi) +
					expressionNum(express, desired, lo, i-1)*expressionNum(express, !desired, i+1, hi) +
					expressionNum(express, !desired, lo, i-1)*expressionNum(express, desired, i+1, hi)
			}
		} else { // desired=false
			switch symbol {
			case '^':
				num += expressionNum(express, desired, lo, i-1)*expressionNum(express, desired, i+1, hi) +
					expressionNum(express, !desired, lo, i-1)*expressionNum(express, !desired, i+1, hi)
			case '&':
				num += expressionNum(express, desired, lo, i-1)*expressionNum(express, desired, i+1, hi) +
					expressionNum(express, desired, lo, i-1)*expressionNum(express, !desired, i+1, hi) +
					expressionNum(express, !desired, lo, i-1)*expressionNum(express, desired, i+1, hi)
			case '|':
				num += expressionNum(express, desired, lo, i-1) * expressionNum(express, desired, i+1, hi)
			}
		}
	}
	return num
}

// 动态规划法
// 之所以上面递归方法的时间复杂度高，是因为没有记录中间结果

// 首先什么是起始状态？express[1:1], express[2:2], ..., express[n:n]这些只有单个元素的表达式是起始状态
// 从所有单元素，然后二元素，... ，直到n元素的表达式。

// 怎么定义状态？要保证能够表达 k元素
// 很容易想到定义 express[i...j]为状态，这也是有关字符串的dp问题的一种常用状态定义方式。那么需要用矩阵来存储状态值。
// t[i][j] = express[i...j]的desired为true的种数
// f[i][j] = express[i...j]的desired为false的种数

// init：t[i][i], f[i][i]
// 可以发现，i<=j才有意义，所以状态矩阵t和f只有上三角部分才有意义。
// 而初值是对角线上的所有元素，目标是右上角的元素。

// 那么可以有两种方式迭代
// 1. 斜向上迭代
// 2. 从左向右，从下往上。

// 斜向上迭代
// 初始斜线位置：对角线，i=0，j=0
// 斜线结束条件：j<= N
func ExpressionNum2(express string, desired bool) int {
	N := len(express)
	t := make([][]int, N)
	for i := range t {
		t[i] = make([]int, N)
	}
	f := make([][]int, N)
	for i := range f {
		f[i] = make([]int, N)
	}
	cond := func(condition bool, a, b int) int {
		if condition {
			return a
		}
		return b
	}
	// init
	for i := 0; i < N; i += 2 { // 所有位置为布尔值的赋上初值
		t[i][i] = cond(express[i] == '1', 1, 0)
		f[i][i] = cond(express[i] == '0', 1, 0)
	}
	// iterate
	row, col := 0, 2 // 初始斜线位置
	for col < N {
		i, j := row, col
		for i < N-col {
			// 更新express[i:j]位置的状态，根据其中所有symbol来划分左右
			// i和j的位置都是bool值，i+1的位置是symbol
			// 划分点k, 划分为[i,k] k+1 [k+2, j],k的取值范围为[i:2:j)
			for k := i; k < j; k += 2 {
				if express[k+1] == '&' {
					t[i][j] += t[i][k] * t[k+2][j]
					f[i][j] += f[i][k]*f[k+2][j] + f[i][k]*t[k+2][j] + t[i][k]*f[k+2][j]
				} else if express[k+1] == '|' {
					t[i][j] += t[i][k]*t[k+2][j] + f[i][k]*t[k+2][j] + t[i][k]*f[k+2][j]
					f[i][j] += f[i][k] * f[k+2][j]
				} else { // ^
					t[i][j] += t[i][k]*f[k+2][j] + f[i][k]*t[k+2][j]
					f[i][j] += f[i][k]*f[k+2][j] + t[i][k]*t[k+2][j]
				}
			}
			// 斜线中更新下一个元素位置
			i += 2
			j += 2
		}
		// 更新斜线初始位置
		col += 2
	}
	return cond(desired, t[0][N-1], f[0][N-1])
}

// 从左向右，从下往上迭代
func ExpressionNum3(express string, desired bool) int {
	N := len(express)
	t := make([][]int, N)
	for i := range t {
		t[i] = make([]int, N)
	}
	f := make([][]int, N)
	for i := range f {
		f[i] = make([]int, N)
	}
	cond := func(condition bool, a, b int) int {
		if condition {
			return a
		}
		return b
	}
	// iterate
	for j := 0; j < N; j += 2 { // 先列
		t[j][j] = cond(express[j] == '1', 1, 0)
		f[j][j] = cond(express[j] == '0', 1, 0)
		for i := j - 2; i >= 0; i -= 2 { // 后行
			// [i, k], k+1 [k+2, j]
			for k := i; k < j; k += 2 {
				if express[k+1] == '&' {
					t[i][j] += t[i][k] * t[k+2][j]
					f[i][j] += f[i][k]*f[k+2][j] + f[i][k]*t[k+2][j] + t[i][k]*f[k+2][j]
				} else if express[k+1] == '|' {
					t[i][j] += t[i][k]*t[k+2][j] + f[i][k]*t[k+2][j] + t[i][k]*f[k+2][j]
					f[i][j] += f[i][k] * f[k+2][j]
				} else { // ^
					t[i][j] += t[i][k]*f[k+2][j] + f[i][k]*t[k+2][j]
					f[i][j] += f[i][k]*f[k+2][j] + t[i][k]*t[k+2][j]
				}
			}
		}
	}
	return cond(desired, t[0][N-1], f[0][N-1])
}

// 动态规划方法，空间复杂度为O(N^2)
// 时间复杂度为O(N^3)

func main() {
	fmt.Println(ExpressionNum("1^0|0|1", false))
	fmt.Println(ExpressionNum2("1^0|0|1", false))
	fmt.Println(ExpressionNum3("1^0|0|1", false))
}
