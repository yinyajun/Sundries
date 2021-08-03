package main

import (
	"CodeGuide/base/utils"
	"fmt"
	"time"
)

// O(2^N)
func f1(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return f1(n-1) + f1(n-2)
}

// 利用压缩空间的动态规划，O(N)
func f2(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	dp_0, dp_1 := 1, 1
	for i := 3; i <= n; i++ {
		dp_1, dp_0 = dp_1+dp_0, dp_1
	}
	return dp_1
}

// 二阶递推式可以视为线性变换
// [f(n), f(n-1)] = [f(2), f(1)] * A^(n-2)
// 那么将问题转化为矩阵的n次方的问题，可以有O(log N)的解法
func f3(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	A := [][]int{{1, 1}, {1, 0}}
	res := MatrixPower(A, n-2)
	return res[0][0] + res[0][1]
}

// 方阵的p次方
// 原理：二进制思想。将一个十进制数转化为二进制数。二进制数的位数为logN。任何小于N的数总可以用若干个2^k的数之和表示
// 举例：75:(1001011)2=64+8+2+1=2^6+2^3+2^1+2^0
func MatrixPower(m [][]int, p int) [][]int {
	row, col := len(m), len(m[0])
	utils.Assert(row == col)

	res := make([][]int, row)
	// 将res初始化为单位矩阵
	for i, _ := range res {
		res[i] = make([]int, col)
		res[i][i] = 1
	}
	tmp := m
	for ; p != 0; p >>= 1 { // 将p视为二进制，二进制位数为log2(p)
		if (p & 1) != 0 {
			res = MatMul(res, tmp)
		}

		tmp = MatMul(tmp, tmp)
	}
	return res
}

// 整数矩阵相乘
func MatMul(m1, m2 [][]int) [][]int {
	m1Row, m1Col := len(m1), len(m1[0])
	m2Row, m2Col := len(m2), len(m2[0])
	utils.Assert(m1Col == m2Row)

	// res : m1Row * m2Col
	res := make([][]int, m1Row)
	for i, _ := range res {
		res[i] = make([]int, m2Col)
	}

	for i := 0; i < m1Row; i++ {
		for j := 0; j < m2Col; j++ {
			for k := 0; k < m1Col; k++ {
				res[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return res
}

func main() {
	n := 46
	s := time.Now()
	fmt.Println(f1(n), time.Since(s))
	s = time.Now()
	fmt.Println(f2(n), time.Since(s))
	s = time.Now()
	fmt.Println(f3(n), time.Since(s))

}
