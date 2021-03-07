package main

import (
	"fmt"
	"time"
)

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

// 因为这里是尾递归，类似于迭代
// 没有返回值，那么当前状态下的最优值是随着递归函数的参数传递的
// 这种方式类似于顺序的动态规划，每次state参数都是从第一行到当前行的可插入位置列表（当前状态下的最优值）
func Queens(n int, state []int, ret *[][]int) {
	for pos := 0; pos < n; pos++ {
		if !conflict(state, pos) {
			if len(state) == n-1 { // last line
				*ret = append(*ret, append(state, pos))
			} else {
				// 虽然state是slice是指针传递，相当于全局变量
				// 但是我们这里使用的时候，主要是用append，而append会将生成一个新的slice
				// 此时相当于值传递，原函数中state变量并没有因此改变，因此不需要特意的回溯改变
				// 如果state是一个链表，或者state的slice长度已经固定，只需要修改其元素，就需要回溯改变了
				Queens(n, append(state, pos), ret) // 尾递归
				// 此时的state和递归调用前一毛一样
			}
		}
	}
}

// end代表current line, state是长度为n的slice
func conflict2(state []int, end, pos int) bool {
	for i := 0; i < end; i++ {
		diff := pos - state[i]
		if diff == 0 || diff == end-i || -diff == end-i {
			return true
		}
	}
	return false
}

func Queens2(n, end int, state []int, ret *[][]int) {
	if end == n {
		fmt.Println(state)
		*ret = append(*ret, state)
	}
	for pos := 0; pos < n; pos++ {
		if !conflict2(state, end, pos) { // 只能在前n-1行有效
			// pos位置合法
			state[end] = pos
			Queens2(n, end+1, state, ret) // 这也不用回溯修改state，因为可以直接覆盖end位置的值
		}
	}
}

// 叶子节点为基本情况（确定最后一行可以插入的位置），根据子节点的情况组成父亲节点（不断组成上一行可以插入的位置）
// 这种方式类似于逆序的动态规划，每次返回值的是当前行到最后一行的可插入位置列表（返回值可以认为是当前状态下的最优值）
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

// 上述方法中conflict消耗时间太多，这里使用位运算加速
func NQueens3(n int) int {
	return queens3((1<<n)-1, 0, 0, 0)
}

// uppperLim : 1111(n=4)，可插入值的列，可插入为1
// colLim: 哪些列已经有值了，有值为1
// leftDiaLim：哪些列的左斜对角线上有值，有值为1
// rightDiaLim：哪些列的右斜对角线上有值，有值为1
func queens3(upperLim, colLim, leftDiaLim, rightDiaLim int) int {
	if colLim == upperLim { // 所有列可插入值都满了
		return 1
	}
	pos := upperLim & (^(colLim | leftDiaLim | rightDiaLim)) // 可插入值的列中刨去不能插入值的列
	lowbit := 0                                              // 最低位
	res := 0
	for pos != 0 { // 遍历所有可插入的位置
		lowbit = pos & (-pos)                   //lowbit= x & (~x+1)
		pos -= lowbit                           // 下一个可插入的位置
		res += queens3(upperLim, colLim|lowbit, // 可插入列更新
			(leftDiaLim|lowbit)<<1,  // 左斜线有值列更新
			(rightDiaLim|lowbit)>>1) // 有些线有值列更新
	}
	return res
}

func test(a []int) {
	a[0] = 8
	fmt.Printf("%p\n", a)
}

func main() {
	start := time.Now()
	fmt.Println(NQueens(8))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println(NQueens2(8))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println(NQueens3(8))
	fmt.Println(time.Since(start))

	n := 4
	state := make([]int, n)
	ret := [][]int{}
	Queens2(n, 0, state, &ret)
	fmt.Println(ret)

}
