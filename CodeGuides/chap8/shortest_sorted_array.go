package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 需要排序的最短子数组的长度
// [1,5,3,4,2,6,7] ret=4    [5,3,4,2]

// 寻找 左边界  和  右边界
// 分别从左到右、从右到左 寻找第一个逆序对
// 寻找所有逆序对，是归并排序中merge方法
// 而寻找第一个逆序对的话，只要按指定方向，遍历一遍即可

// O(N)
func getMinLength(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}
	// [0,i) 有序  [i, length) unknown
	left := -1
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			left = i - 1
			break
		}
	}
	if left == -1 { // sorted
		return 0
	}

	right := -1
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			right = i + 1
			break
		}
	}
	fmt.Println(left, right)
	return right - left + 1
}

// ！！！！！！！！
// 这里有个细节搞错了
// 初始化的时候left=0，那么会混淆两种情况
// 1. (0,1)两个数组形成逆序对，此时left=0
// 2. 如果没有任何逆序对，此时left=0

// 同理，right=len(arr)-1，这样的初始化也会有类似的问题
// 所以left，right的初始化应该避开 [0, length-1]的值
// 可以是，left=length， right=-1

// 为什么找完left后要提前中止？因为如果有序，那么此时left和right都是初始值，[left, right]将不会是一个有效区间

// O(2N)
func getMinLength2(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}

	// (i, length] 有序
	min := arr[len(arr)-1]
	left := -1
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > min {
			left = i
		} else {
			min = utils.MinInt(min, arr[i])
		}
	}
	if left == -1 {
		return 0
	}

	// [0, i) 有序
	right := -1
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < max {
			right = i
		} else {
			max = utils.MaxInt(max, arr[i])
		}
	}
	fmt.Println(left, right)
	return right - left + 1
}

//func main() {
//	arr := []int{1, 5, 3, 4, 2, 6, 7}
//	fmt.Println(getMinLength(arr))
//	fmt.Println(getMinLength2(arr))
//
//	arr = []int{5, 4}
//	fmt.Println(getMinLength(arr))
//	fmt.Println(getMinLength2(arr))
//
//	arr = []int{1, 2, 3}
//	fmt.Println(getMinLength(arr))
//	fmt.Println(getMinLength2(arr))
//}
