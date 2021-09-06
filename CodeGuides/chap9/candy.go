package main

import "fmt"

// 1. 每个孩子至少有1个糖果
// 2. 相邻两个孩子，得分较多的必须多拿糖果
// 求最少需要的糖果，时间和空间复杂度要求为O(N)

// 思路，找到得分数组的上升沿和下降沿

// 从start开始寻找下降沿的结束位置（上升沿的起始位置）
func nextMin(arr []int, start int) int {
	for i := start; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			continue
		}
		//arr[i] <= arr[i+1]，开始不能下降了
		return i
	}
	return len(arr) - 1
}

// 因为下降沿可以找到结束位置，用来计算下降沿的分配糖果数目
func regionSum(left, right int) int {
	n := right - left + 1
	return (1 + n) * n / 2
}

func minCandies(arr []int) int {
	var i int //遍历arr的index
	var res int
	var lSize, rSize int
	var next int

	// 找到首个上升沿
	i = nextMin(arr, 0)
	res += regionSum(0, i)
	lSize = 1
	i++

	for i < len(arr) {
		if arr[i] > arr[i-1] { // 仍然处于上升沿
			lSize++
			res += lSize
			i++
		} else if arr[i] < arr[i-1] { // 处于下降沿
			next = nextMin(arr, i-1)
			rSize = next - (i - 1) + 1
			if lSize > rSize {
				res += regionSum(i-1, next) - rSize
			} else {
				res += regionSum(i-1, next) - lSize
			}
			lSize = 1
			i = next + 1
		} else {
			res += 1
			lSize = 1 // import, 相等计算上升沿也算下降沿
			i++
		}
	}
	return res
}

// 进阶：相邻的孩子之间得分一样，糖果数必须相同
func minCandies2(arr []int) int {
	var i int //遍历arr的index
	var res int
	var lSize, rSize int
	var candies int
	var next int
	var same int // 相等元素的个数

	// 找到首个上升沿
	i = NextMin2(arr, 0)
	res, _ = RightCandiesAndSize(arr, 0, i)
	lSize = 1
	same = 1
	i++

	for i < len(arr) {
		if arr[i] > arr[i-1] { // 仍然处于上升沿
			lSize++
			res += lSize
			same = 1
			i++
		} else if arr[i] < arr[i-1] { // 处于下降沿
			next = NextMin2(arr, i-1)
			candies, rSize = RightCandiesAndSize(arr, i-1, next)
			if lSize >= rSize {
				res += candies - rSize
			} else {
				res += candies - rSize - same*lSize + same*rSize
			}
			lSize = 1
			i = next + 1
			same = 1
		} else {
			res += lSize
			same++
			i++
		}
	}
	return res
}

// 从start开始寻找下降沿的结束位置（下降沿：只要没有上升）
func NextMin2(arr []int, start int) int {
	for i := start; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			continue
		}
		// arr[i] < arr[i+1], 只有开始上升，才说明下降沿结束
		return i
	}
	return len(arr) - 1
}

func RightCandiesAndSize(arr []int, left, right int) (int, int) {
	var size, res int
	res = 1
	size = 1

	for i := right - 1; i >= left; i-- {
		if arr[i] > arr[i+1] {
			size++
			res += size
		} else { // arr[i] == arr[i+1]，没有<情况
			res += size
		}
	}
	return res, size
}

func main() {
	arr := []int{0, 1, 2, 3, 3, 3, 2, 2, 2, 2, 2, 1, 1}
	fmt.Println(minCandies(arr))
	fmt.Println(minCandies2(arr))
}
