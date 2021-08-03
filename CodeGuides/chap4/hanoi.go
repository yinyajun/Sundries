package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 1. 将前n-1个盘子从左塔移动到中塔
// 2. 将第n个盘子从左塔移动到右塔
// 3. 将n-1个盘子从中塔移动到右塔

func Hanoi(n int) {
	utils.Assert(n > 0)
	move(n, "left", "mid", "right")
}

// 将n个盘子从from塔移动到to塔，借助于mid塔
func move(n int, from, mid, to string) {
	if n == 1 {
		fmt.Println("move from", from, "to", to)
		return // 竟然漏了
	}
	move(n-1, from, to, mid)
	move(1, from, mid, to)
	move(n-1, mid, from, to)
}

// 求递归的时间复杂度。
// 将n个盘子从from移动到to需要的最少步数
// s(n)=step1+1+step3=s(n-1)+1+s(n-1)=2s(n-1)+1
// s(1) =1
// s(n) + 1 = 2[s(n-1)+1], 对于{s(n-1)+1}而言，是等比数列
// s(n)+1 = 2^n
// s(n)=2^n-1

// arr[n-1]表示第n个盘子在哪个柱子上
// 对于1~n个圆盘从from到to的步数
// 1. 圆盘n在from上，继续考察1~n-1从from到mid
// 2. 圆盘n在to上，至少走了2^(n-1)-1+1，继续考察1~n-1从mid到to
// 3. 圆盘n在mid上，不可能，直接返回-1

func Step(arr []int) int {
	// 0~n个圆盘从from到to
	return step(arr, len(arr)-1, 1, 2, 3)
}

// 这里i是index，相比于圆盘数小1
// 时间复杂度计算。线性递归，每次递归中至多调用一次递归，除去递归部分的时间复杂度为O(1)
// 递归最多调用n次，所以时间复杂度为O(n)，空间复杂度为O(n)
func step(arr []int, i, from, mid, to int) int {
	if i == -1 {
		return 0
	}
	if arr[i] != from && arr[i] != to {
		return -1
	}

	if arr[i] == from {
		return step(arr, i-1, from, to, mid)
	} else { // arr[i] = to
		rest := step(arr, i-1, mid, from, to)
		if rest == -1 {
			return -1
		}
		return rest + (1 << i)
	}
}

// 为了降低空间复杂度，需要改为非递归的写法
// dp[i] = dp[i-1], if arr[i]==from
//       = dp[i-1] + 1<<i, if arr[i]==to
// 这里类似于尾递归
// 尾递归之所以可以比较容易的转化为迭代，就是因为尾递归中不需要额外保留上次递归函数的函数栈
// 所以只要在迭代中将所使用到的局部变量变化即可
func step2(arr []int) int {
	from, mid, to := 1, 2, 3
	i := len(arr) - 1
	res := 0

	for i >= 0 {
		if arr[i] != from && arr[i] != to {
			return -1
		}
		if arr[i] == from {
			// switch to step(arr, i-1, from, to, mid)
			to, mid = mid, to
		} else { // arr[i] = to
			// switch to step(arr, i-1, mid, from , to)
			res += 1 << i
			from, mid = mid, from
		}
		i--
	}
	return res
}

func main() {
	Hanoi(2)
}
