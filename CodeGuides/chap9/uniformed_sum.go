package main

import (
	"CodeGuide/base/utils"
	"fmt"
	"sort"
)

// 正整数数组的最小不可组成和

func uniformSum1(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	var min, sum int
	min = 1 << 31
	for i := 0; i < len(arr); i++ {
		min = utils.MinInt(min, arr[i])
		sum += arr[i]
	}

	memo := dpProcess3(arr)
	for k := min + 1; k < sum; k++ { // 如何判断子集能组成k，使用暴力法，将所有的子集计算出来
		if !memo[k] {
			return k
		}
	}
	return sum + 1
}

var c int

func indent2(c int) {
	for i := 0; i < c; i++ {
		fmt.Print("        ")
	}
}

// 计算[idx...N-1]的所有子集的和，并且[0...idx)的子集和为sum
// 逆序解法，已知初始状态idx=0，sum([0,0))=0
func process1(arr []int, idx, sum int, memo map[int]struct{}) {

	indent2(c)
	c++
	fmt.Println(arr[idx:], sum)

	if idx == len(arr) {
		memo[sum] = struct{}{}

		c--
		indent2(c)
		fmt.Println("return", sum)
		return
	}

	process1(arr, idx+1, sum+arr[idx], memo) // including arr[idx]
	process1(arr, idx+1, sum, memo)          // excluding arr[idx]

	c--
	indent2(c)
	fmt.Println("return", sum)
}

var sum int

// 使用回溯的方式
func process11(arr []int, idx int, memo map[int]struct{}) {
	indent2(c)
	c++
	fmt.Println(arr[idx:], sum)

	if idx == len(arr) {
		memo[sum] = struct{}{}
		c--
		indent2(c)
		fmt.Println("return", sum)
		return
	}

	// 选择列表 [加上arr[idx], 不加]
	// choose adding arr[idx]
	sum += arr[idx]
	process11(arr, idx+1, memo)
	sum -= arr[idx] // withdraw choose
	process11(arr, idx+1, memo)

	c--
	indent2(c)
	fmt.Println("return", sum)
}

// 计算[0...idx]的所有子集的和，并且(idx, N-1]的子集和为sum
// 顺序解法，已知终止状态idx=N-1, sum((N-1, N-1])=0
func process2(arr []int, idx, sum int, memo map[int]struct{}) {

	indent2(c)
	c++
	fmt.Println(arr[:idx+1], sum)

	if idx == -1 {
		memo[sum] = struct{}{}

		c--
		indent2(c)
		fmt.Println("return1", sum)

		return
	}

	process2(arr, idx-1, sum+arr[idx], memo)
	process2(arr, idx-1, sum, memo)

	c--
	indent2(c)
	fmt.Println("return", sum)
}

//// dp[i] = {k+arr[i] |k in dp[i+1]} + {k | k in dp[i+1]}
//// dp[i] : [i...N-1]的所有子集的和
//// *****错误******这个方案看起来没有问题，但是有个严重问题，在遍历map的时候，改变map结构，这会影响遍历
//func dpProcess(arr []int) {
//	memo := make(map[int]struct{})
//
//	// dp[N] = {0}
//	memo[0] = struct{}{}
//	// dp[i] = {k+arr[i] |k in dp[i+1]} + {k | k in dp[i+1]}
//	for i := len(arr) - 1; i >= 0; i-- {
//		for k := range memo {
//			memo[k+arr[i]] = struct{}{}
//		}
//		fmt.Println(arr[i], memo)
//	}
//	fmt.Println(memo)
//}

// 自顶向下实现
func dpProcess(arr []int) []bool {
	var sum int
	for _, v := range arr {
		sum += v
	}
	memo := make([]bool, sum+1)

	_process(arr, 0, memo)
	fmt.Println(memo)
	return memo
}

// dp[i] = {k+arr[i] |k in dp[i+1]} + {k | k in dp[i+1]}
// dp[i] : [i...N-1]的所有子集的和
// !!!! 注意这里k的遍历要从大到小，因为新产生的数字将会比k更大，如果从小到大遍历，会影响遍历结果
// 时间复杂度O(N*sum), 空间O(sum)
func _process(arr []int, idx int, memo []bool) {
	if idx == len(arr) {
		memo[0] = true
		return
	}

	_process(arr, idx+1, memo)

	for k := len(memo) - 1; k >= 0; k-- {
		if memo[k] == true {
			memo[k+arr[idx]] = true
		}
	}
}

// 自下向上实现逆序解法
// dp[i] = {k+arr[i] |k in dp[i+1]} + {k | k in dp[i+1]}
// dp[i] : [i...N-1]的所有子集的和
func dpProcess2(arr []int) []bool {
	var sum int
	for _, v := range arr {
		sum += v
	}
	memo := make([]bool, sum+1)

	// dp[N] = 0
	memo[0] = true
	// dp[i] = {k+arr[i] |k in dp[i+1]} + {k | k in dp[i+1]}
	for i := len(arr) - 1; i >= 0; i-- {
		for k := len(memo) - 1; k >= 0; k-- {
			if memo[k] == true {
				memo[k+arr[i]] = true
			}
		}
	}
	fmt.Println(memo)
	return memo
}

// 自下向上实现顺序解法
// dp[i]:[0....i]的所有子集的的和
// dp[i] = {k+arr[i] |k in dp[i-1]} + {k | k in dp[i-1]}
func dpProcess3(arr []int) []bool {
	var sum int
	for _, v := range arr {
		sum += v
	}
	memo := make([]bool, sum+1)

	// dp[-1] = 0
	memo[0] = true
	// dp[i] = {k+arr[i] |k in dp[i-1]} + {k | k in dp[i-1]}
	for i := 0; i < len(arr); i++ {
		for k := len(memo) - 1; k >= 0; k-- {
			if memo[k] == true {
				memo[k+arr[i]] = true
			}
		}
	}
	fmt.Println(memo)
	return memo
}

// 进阶问题： 已知正数数组中有1，能否更快地得到最小不可组成和？
// 先排序，升序
// arr[0] = 1, range=[1,1]
// if arr[i] > range + 1，那么 range +1 该数就是最小不可组成和，因为后面的数字只会更大
// if arr[i] <= range + 1，那么 range可以扩展到[1, range+arr[i]]

// 排序O(NlogN), 后面只需要使用O(N)时间，空间为O(1)
func uniformSum2(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	r := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] > r+1 {
			return r + 1
		} else {
			r += arr[i]
		}
	}
	return r + 1
}

//func main() {
//	arr := []int{3, 2, 5}
//
//	//memo := make(map[int]struct{})
//	//process1(arr, 0, 0, memo)
//	//fmt.Println(memo)
//	//
//	//memo = make(map[int]struct{})
//	//process11(arr, 0, memo)
//	//fmt.Println(memo)
//	//
//	//memo = make(map[int]struct{})
//	//process2(arr, len(arr)-1, 0, memo)
//	//fmt.Println(memo)
//
//	dpProcess(arr)
//	dpProcess2(arr)
//	dpProcess3(arr)
//
//	fmt.Println(uniformSum1(arr))
//	fmt.Println(uniformSum1([]int{1, 2, 4}))
//	fmt.Println(uniformSum2([]int{1, 2, 4}))
//}
