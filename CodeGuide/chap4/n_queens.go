package main

import (
	"fmt"
	"time"
)

//func NQueen(n int) int {
//
//}

func conflict(state []int, nextX int) bool {
	nextY := len(state)
	for i := 0; i < nextY; i++ {
		if nextX-state[i] == 0 || nextX-state[i] == nextY-i || state[i]-nextX == nextY-i { // 同一列或者在同一斜线中
			return true
		}
	}
	return false
}

// 递归、回溯、动态规划、树遍历
// 这种方式类似于顺序的动态规划，关注于从开始到当前的量（state的长度）
// 因为这里是尾递归，类似于迭代
func Queens(n int, state []int, ret *[][]int) {
	for pos := 0; pos < n; pos++ {
		if !conflict(state, pos) {
			if len(state) == n-1 { // last line
				*ret = append(*ret, append(state, pos))
			} else {
				Queens(n, append(state, pos), ret) // 尾递归
			}
		}
	}
}

// 这种方式类似于逆序的动态规划，关注于当前到最后的量（state下面的行）
// def queens(num=8, state=()):
//     for pos in range(num):
//         if not conflict(state, pos):
//             if len(state) == num - 1:
//                 yield (pos,)
//             else:
//                 for result in queens(num, state + (pos,)):
//                     yield (pos,) + result

func NQueens(n int) int {
	ret := 0
	queens(n, []int{}, &ret)
	return ret
}

func queens(n int, state []int, num *int) {
	for pos := 0; pos < n; pos++ {
		if !conflict(state, pos) {
			if len(state) == n-1 { // last line
				*num = *num + 1
			} else {
				queens(n, append(state, pos), num)
			}
		}
	}
}

func NQueens2(n int) int {
	return queens2(n, []int{})
}

func queens2(n int, state []int) int {
	if len(state) == n {
		return 1
	}
	res := 0
	for pos := 0; pos < n; pos++ {
		if !conflict(state, pos) {
			res += queens2(n, append(state, pos))
		}
	}
	return res
}

func NQueens3(n int) int {
	return queens3((1<<n)-1, 0, 0, 0)
}

func queens3(upperLim, colLim, leftDiaLim, rightDiaLim int) int {
	if colLim == upperLim {
		return 1
	}
	pos := upperLim & (^(colLim | leftDiaLim | rightDiaLim))
	lowbit := 0
	res := 0
	for pos != 0 {
		lowbit = pos & (-pos) //lowbit= x & (~x+1)
		pos -= lowbit
		res += queens3(upperLim, colLim|lowbit,
			(leftDiaLim|lowbit)<<1,
			(rightDiaLim|lowbit)>>1)
	}
	return res
}

func main() {
	start := time.Now()
	fmt.Println(NQueens(8))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println(NQueens2(8))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println(NQueens3(16))
	fmt.Println(time.Since(start))
}
