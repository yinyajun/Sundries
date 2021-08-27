package main

import (
	"math"
)

// 柱形图中最大矩形
// 矩形=长*宽

// 暴力解法：遍历的过程中，枚举长或者宽

// 暴力解法1：枚举宽，高度为宽度内的柱子最小高度
// 妥妥的O(N^2)时间，主要是每次寻找最小高度花费O(N)的时间，太耗时了
func MaxRectHist(heights []int) int {
	var ans int
	var area int
	for left := 0; left < len(heights); left++ {
		minHeight := int(math.MaxInt32)
		for right := left; right < len(heights); right++ {
			if heights[right] < minHeight {
				minHeight = heights[right]
			}
			area = (right - left + 1) * minHeight
			if area > ans {
				ans = area
			}
		}
	}
	return ans
}

// 暴力解法2：枚举高度
// 使用一重循环枚举某根柱子高度，然后从这根柱子左右延伸，知道遇到高度比起小的柱子，就确定了这个高度的左右边界。
// 这里每次确定左右边界，花费O(N)的时间，太耗时了
// 如果将其优化为O(1)的时间，那么复杂度就降下来了
func MaxRectHist2(heights []int) int {
	var ans int
	var area int
	var left, right int

	for k := 0; k < len(heights); k++ {
		h := heights[k]

		// 确定左边界
		for left = k; (left-1) >= 0 && heights[left-1] >= h; left-- {
		}

		// 确定右边界
		for right = k; right+1 < len(heights) && heights[right+1] >= h; right++ {
		}
		area = h * (right - left + 1)
		if area > ans {
			ans = area
		}
	}
	return ans
}

// 确定左右边界，就是寻找最近的小于其高度的元素
// 如何用O(1)的时间去找到呢？
// 首先在遍历的时候，每次仅仅维护一个最近小于其高度的元素是不够的，因为下次遍历到一个很小的元素时，该答案会失效，需要额外花费O(N)的时间再次寻找答案
// 所以需要有个数据结构，存储一些备选答案（移除哪些不可能的答案，遍历过程中，i<j && height[i] >= height[j], heights[i]不可能是备选答案），在答案失效的时候，备选答案可以用O(1)的时间转正。
// 那么此时发现备选答案满足单调性质：height[j0] < height[j1] < ...
// 而最近加入数据结构中的备选答案就是当前答案，那么这样LIFO性质，使得需要栈来存储
// 现在这样的数据结构就是单调栈（单调递减，从出栈角度）

// 有个细节问题：这里的单调栈中需要严格单调吗？需要！如果没有严格单调，就会出现相同元素，那么栈顶可能就是和当前元素相等的元素，不是正确的左右边界
func MaxRectHist3(heights []int) int {
	var ans int
	var area int
	left, right := make([]int, len(heights)), make([]int, len(heights))

	// 使用单调栈确定左边界
	monoStack := []int{}
	for i := 0; i < len(heights); i++ {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		// stack is empyt || cur > stack.top
		if len(monoStack) == 0 {
			left[i] = -1 // 当左边没有柱子的时候，使用-1作为哨兵
		} else { // cur > stack.top
			left[i] = monoStack[len(monoStack)-1]
		}
		// push cur into stack
		monoStack = append(monoStack, i)
	}
	// 使用单调栈确定右边界
	monoStack = []int{}
	for j := len(heights) - 1; j >= 0; j-- {
		for len(monoStack) > 0 && heights[j] <= heights[monoStack[len(monoStack)-1]] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		// stack is empty || cur > stack.top
		if len(monoStack) == 0 {
			right[j] = len(heights)
		} else {
			right[j] = monoStack[len(monoStack)-1]
		}
		// push cur into stack
		monoStack = append(monoStack, j)
	}
	for i := 0; i < len(heights); i++ {
		area = heights[i] * (right[i] - left[i] - 1)
		if area > ans {
			ans = area
		}
	}
	return ans
}

// 上面的方法已经将时间复杂度降到了O(N)，因为每次遍历只能确定左侧边界或者右侧边界，所以遍历了三次
// 有什么办法在一次遍历的时候完成吗？遍历到i时，左边界在入栈前确定，由于单调栈中一定需要符合<, 所以左边界 < h[i]
// 出栈的时候，已经遍历到cur, h[i] >= h[cur]，cur可能是将要出栈的元素i的右边界
// 但是呢，由于是>=关系，不能保证右边界是正确的，但是重复柱子中最右边那个一定是正确的
// 因此可以在出栈的时候确定右边界，这样就不需要遍历两次了
func MaxRectHist4(heights []int) int {
	var ans int
	var area int
	monoStack := Stack{}
	left, right := make([]int, len(heights)), make([]int, len(heights))
	for i := 0; i < len(right); i++ {
		right[i] = len(heights) // 默认都是右侧哨兵
	}

	for i := 0; i < len(heights); i++ {
		for !monoStack.empty() && heights[i] <= heights[monoStack.top()] {
			right[monoStack.top()] = i
			monoStack.pop()
		}
		// stack is empty || h[i] > stack.top
		if monoStack.empty() {
			left[i] = -1 // 左侧哨兵
		} else {
			left[i] = monoStack.top()
		}
		// push into stack
		monoStack.push(i)
	}
	for i := 0; i < len(heights); i++ {
		area = heights[i] * (right[i] - left[i] - 1)
		if area > ans {
			ans = area
		}
	}
	return ans
}

type Stack []int

func (s *Stack) push(e int) { *s = append(*s, e) }

func (s *Stack) pop() int {
	e := s.top()
	*s = (*s)[:len(*s)-1]
	return e
}

func (s *Stack) top() int {
	if len(*s) == 0 {
		panic("underflow")
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) empty() bool { return len(*s) == 0 }

//func main() {
//	heights := []int{2, 1, 5, 6, 2, 3}
//	fmt.Println(MaxRectHist(heights))
//	fmt.Println(MaxRectHist2(heights))
//	fmt.Println(MaxRectHist3(heights))
//	fmt.Println(MaxRectHist4(heights))
//	//s := Stack{}
//	//s.push(1)
//	//s.push(2)
//	//s.push(3)
//	//s.pop()
//	//fmt.Println(s)
//	//fmt.Println(s.top())
//}
