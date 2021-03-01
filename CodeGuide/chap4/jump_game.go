package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 顺序解法
// dp[i]: 从小于i的位置跳到i位置的最小跳跃次数
// dp[i]=min{
//		dp[i-1] + 1, if arr[i-1] == 1  && i-1>=0,
//		dp[i-2] + 1, if arr[i-2] == 2  && i-2>=0,
//	 	...
//		dp[i-k] + 1, if arr[i-k] == k && i-k>=0,
//	}
// init
// dp[0] = 0
// 时间复杂度 平方，空间复杂度线性；不满足要求
func JumpGame(arr []int) int {
	dp := make([]int, len(arr))
	for i := range dp {
		dp[i] = i // 填充一个比较大的树，最多就是i次
	}
	dp[0] = 0

	for i := 1; i < len(arr); i++ {
		for k := 0; k <= i; k++ {
			if arr[i-k] >= k {
				dp[i] = utils.MinInt(dp[i], dp[i-k]+1)
			}
		}
	}
	return dp[len(arr)-1]
}

// 为了要符合题目要求的复杂度，要换一种方式的dp，所求的答案不一定要是dp数组存放的最优值，也可以是dp数组的index
// 这里以跳的次数作为状态，注意arr的索引i并不是状态，通过遍历i，发现不符合条件的才更新s
// dp[s]:跳s次最远可以到的位置
// dp[0] = 0

// 状态s更新：if (dp[s]<i) s++ else pass // 如果第s跳够不到i，那么多跳一次，变为s+1跳，否则不变
// dp[s+1] = max(dp[s+1], i+arr[i])
// 每次遍历到新的arr[i]，用来更新下一跳的最远距离，因为arr[i]是通过当前跳来的，从arr[i]跳走，只能是从下一跳
// 下一跳是否保证一定可以够到i呢？一定可以，因为i按1步更新，而每一跳至少为1
// 对于一个本来满足条件的跳数，由于i++，而不满足条件，多增加一跳即可使其满足
func JumpGame2(arr []int) int {
	dp := make([]int, len(arr)+1) // 存储0-N跳

	s := 0
	for i := 0; i < len(arr); i++ { // 遍历所有arr[i]
		if dp[s] < i { // 跳s次够不到i
			s++
		}
		// 此时可以够到i，更新下一跳
		dp[s+1] = utils.MaxInt(dp[s+1], i+utils.MaxInt(1, arr[i]))
		//if dp[s]>=len(arr)-1{ // 如果第s步已经到达最后一位
		//	return s
		//}
	}
	return s
}

func JumpGame3(arr []int) int {
	dp := make([]int, len(arr)+1)

	// init
	s := 1 // 下一跳
	dp[1] = utils.MaxInt(1, arr[0])

	for i := 1; i < len(arr); i++ { // 遍历所有arr[i]
		if dp[s-1] < i { // 跳s-1次够不到i，当前够不到i
			s++
		}
		// 用arr[i]更新第s跳
		dp[s] = utils.MaxInt(dp[s], i+utils.MaxInt(1, arr[i]))
	}
	return s - 1
}

func JumpGame4(arr []int) int {
	cur, next := 0, 0
	step := 0

	for i := 0; i < len(arr); i++ {
		if cur < i { // 当前跳无法够到i
			step++
			cur = next // 进入下一跳
		}
		next = utils.MaxInt(next, i+utils.MaxInt(1, arr[i]))
	}
	return step
}

func main() {
	arr := []int{1, 2, 3, 1, 1, 4}
	fmt.Println(JumpGame(arr))
	fmt.Println(JumpGame2(arr))
	fmt.Println(JumpGame3(arr))
	fmt.Println(JumpGame4(arr))
}
