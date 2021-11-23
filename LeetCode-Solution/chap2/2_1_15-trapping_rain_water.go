/*
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute
how much water it is able to trap after raining.
For example, Given [0,1,0,2,1,0,1,3,2,1,2,1], return 6.


* @Author: Yajun
* @Date:   2021/11/21 23:51
*/

package chap2

import (
	"container/list"
	"solution/utils"
)

// time: O(n^2); space: O(1)
// 暴力解法，对于每个下标的柱子，能够存储的水的数量= min(leftMax_i, rightMax_i) - height[i]
// leftMax_i表示i柱子左侧最高的值，rightMax_i表示i柱子右侧最高的值
// 对于每个柱子，使用O(n)时间向两边扫描获取最大高度，然后花O(n)时间遍历每根柱子。
func trappingRainWater(height []int) int {
	var (
		leftMax, rightMax int
		res               int
	)

	for i := range height {
		leftMax = height[i]
		rightMax = height[i]
		for k := i - 1; k > 0; k-- {
			if height[k] > leftMax {
				leftMax = height[k]
			}
		}
		for k := i + 1; k < len(height); k++ {
			if height[k] > rightMax {
				rightMax = height[k]
			}
		}
		res += utils.MinInt(leftMax, rightMax) - height[i]
	}
	return res
}

/*
暴力解法：时间复杂度高在，每次遍历都需要花费O(n)的时间扫描
思路1：离线处理-使用O(n)时间预处理，然后每次可以花费O(1)的时间取出最大值
思路2：在线处理-在遍历过程中，每次都能用O(1)的时间取出最值
https://leetcode-cn.com/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode-solution-tuvc/
*/

// time: O(n); space: O(n)
// 创建leftMax，rightMax数组，leftMax[i]表示下标i及其左边位置中的最高位置，rightMax[i]亦然。
// leftMax[0] = h[0], leftMax[i] = max(leftMax[i-1], h[i])
// rightMax[n-1] = h[n-1], rightMax[i] = max(rightMax[i+1], h[i])
// 在预处理中，利用已求的最值来更新新的最值
func trappingRainWaterB(height []int) int {
	var (
		n        = len(height)
		leftMax  = make([]int, n)
		rightMax = make([]int, n)
		res      int
	)

	// preprocess
	leftMax[0] = height[0]
	rightMax[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		leftMax[i] = utils.MaxInt(leftMax[i-1], height[i])
	}
	for i := n - 2; i > 0; i-- {
		rightMax[i] = utils.MaxInt(rightMax[i+1], height[i])
	}

	for i := 0; i < n; i++ {
		res += utils.MinInt(leftMax[i], rightMax[i]) - height[i]
	}
	return res
}

// time: O(n); space: O(1)
// 在上述方法中，如何将空间复杂度降低到O(1)?
// i柱子的水量只和min(leftMax_i,rightMax_i)有关，并不需要精确知道两者的每一个
// 具体而言，假如右侧不可能是最小值，那么只需要知道leftMax_i的值即可

// 定义双指针left, right; 定义两个变量 leftMax，rightMax分别记录当前i柱子的左右边界最值
// 1. left =0 , right = n-1; leftMax=rightMax =0
// 2. leftMax = max(leftMax, h[left]); rightMax = max(rightMax, h[right])
// 3. left向右移动，right向左移动，其中高者保持不动，低者向内更新；leftMax，rightMax分别记录当前遍历的最大值
//	 	* h[left] < h[right] => leftMax < rightMax, left处水量= leftMax-h[left], left++
//	 	* h[right] < h[left] => rightMax < leftMax, right处水量= rightMax-h[right], right--
func trappingRainWaterC(height []int) int {
	var (
		n                 = len(height)
		left, right       = 0, n - 1
		leftMax, rightMax int
		res               int
	)
	for left < right {
		leftMax = utils.MaxInt(leftMax, height[left])
		rightMax = utils.MaxInt(rightMax, height[right])
		if height[left] < height[right] {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}
	return res
}

/*
 单调栈，递增栈：栈顶到栈底，递增=>处理next greater问题

 遍历时，
 * h[i] <= top, 入栈 （栈中是可能的左边界）
 * h[i] > top, 出栈 （i是右边界）

 每次出栈的时候，确定了右边界，可以计算以i为右边界的所有雨水量
 width = i - prev(top) - 1
 height = min(h[prev(top)], h[i]) - h[top]

*/
// time: O(n); space: O(1)
func trappingRainWaterD(height []int) int {
	var (
		monoStack = new(list.List)
		w, h      int
		left, top int
		res       int
	)

	for i := range height {
		for monoStack.Len() > 0 && height[i] > height[monoStack.Back().Value.(int)] {
			top = monoStack.Remove(monoStack.Back()).(int) // 每次求top位置的积水
			if monoStack.Len() == 0 {
				break
			}
			left = monoStack.Back().Value.(int)

			w = i - left - 1
			h = utils.MinInt(height[left], height[i]) - height[top]
			res += w * h
		}
		monoStack.PushBack(i)
	}
	return res
}

// time: O(n); space: O(1)
// 类似于方法3，遍历一遍找到最高柱子，将整个height分为两个区间，[lo, max][max, hi]
// 对于左区间，右边界已经确定，只要不断找左边界即可
// 对于右区间，左边界已经确定，只要不断找右边界即可
func trappingRainWaterE(height []int) int {
	var (
		n   = len(height)
		max int
		top int
		res int
	)

	for i, h := range height {
		if h > height[max] {
			max = i
		}
	}

	top = 0
	for i := 0; i < max; i++ {
		if height[i] > top { // top代表左边界
			top = height[i]
		} else {
			res += top - height[i]
		}
	}

	top = 0
	for j := n - 1; j > max; j-- {
		if height[j] > top { // top代表右边界
			top = height[j]
		} else {
			res += top - height[j]
		}
	}
	return res
}
