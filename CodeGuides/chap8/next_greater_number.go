package main

import (
	"CodeGuide/base/fundamentals"
	"CodeGuide/base/utils"
)

// 单调递增栈：从栈顶到栈底递增的栈，栈顶小，栈底大
// 新元素 < top, push
// 新元素 >= top, pop(一直弹，直到新元素可以push)
// 分析：从右往左遍历，如果有个数据结构存储了右边所见到元素，且摒弃了不可能的元素（违反了<规则，这样的元素不可能被看到）
// 当前元素进入数据结构时，同时也要遵循规则，对于不符合规则的元素予以删除（这种后进先出的LIFO->栈）

//  ----------------------------
//	                      |   |
//                   |    |   |
//              |    |    |   |
//  ----------------------------
func NextGreaterNumber(nums []int) []int {
	res := make([]int, len(nums))
	stack := fundamentals.NewLinkedStack()

	for i := len(nums) - 1; i >= 0; i-- {
		for !stack.IsEmpty() && !utils.Less(nums[i], stack.Peek()) {
			stack.Pop()
		}
		// nums[i] < top || s is empty
		if stack.IsEmpty() {
			res[i] = -1
		} else {
			res[i] = stack.Peek().(int)
		}
		// push current element into stack
		stack.Push(nums[i])
	}
	return res
}

// 同样还是next greater number问题，不过这里不是直接问next greater number是多少，而是当前元素距离next greater number的距离
// 所以数据结构中，需要存入元素的索引
func dailyTemperatures(nums []int) []int {
	res := make([]int, len(nums))
	stack := fundamentals.NewLinkedStack()

	for i := len(nums) - 1; i >= 0; i-- {
		for !stack.IsEmpty() && !utils.Less(nums[i], nums[stack.Peek().(int)]) {
			stack.Pop()
		}
		// s is empty || nums[i]  < s.peek
		if stack.IsEmpty() {
			res[i] = 0
		} else {
			res[i] = stack.Peek().(int) - i
		}
		// push current ele into stack
		stack.Push(i)
	}
	return res
}

// 如果数组是循环数组，那么最简单的方式就是将数组翻倍，假设真有一个数组在后面
func nextGreaterNumber2(nums []int) []int {
	res := make([]int, len(nums))
	stack := fundamentals.NewLinkedStack()

	for i := 2*len(nums) - 1; i >= 0; i-- {
		for !stack.IsEmpty() && !utils.Less(nums[i%len(nums)], stack.Peek()) {
			stack.Pop()
		}
		// s is empty || cur < s.Peek()
		if stack.IsEmpty() {
			res[i%len(nums)] = -1
		} else {
			res[i%len(nums)] = stack.Peek().(int)
		}
		stack.Push(nums[i%len(nums)])
	}
	return res
}

//func main() {
//	nums := []int{2, 1, 2, 4, 3}
//	fmt.Println(NextGreaterNumber(nums))
//	nums = []int{73, 74, 75, 71, 69, 76}
//	fmt.Println(dailyTemperatures(nums))
//	nums = []int{2, 1, 2, 4, 3}
//	fmt.Println(nextGreaterNumber2(nums))
//}
