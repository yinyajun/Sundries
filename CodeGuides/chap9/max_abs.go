package main

import (
	"CodeGuide/base/fundamentals"
	"CodeGuide/base/utils"
)

// 将数组分为两部分，所有划分策略中，abs(左边最大值 - 右边最大值)的最大值是多少？

// 暴力解法，遍历所有划分策略
func maxAbs(nums []int) int {
	// 将nums切分
	// 分为[0, i] , [i+1, n-1]
	// i \in [0, n-2], 保证每个区间至少有一个元素
	var leftMax, rightMax, res int
	leftMax, rightMax = -1<<31, -1<<31
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j <= i; j++ {
			leftMax = utils.MaxInt(leftMax, nums[j])
		}
		rightMax = 0 // notice，右区间需要移除元素，所以重新统计最大值
		for j := i + 1; j < len(nums); j++ {
			rightMax = utils.MaxInt(rightMax, nums[j])
		}
		res = utils.MaxInt(res, _abs(rightMax-leftMax))
	}
	return res
}

// 每次在区间中找最值都是很费时间的事
// 类似于在滑动区间上找最值的问题，如果将区间中所有备选答案都记录了
// 最佳答案失效的情况下，备选答案可以用O(1)的时间上位为最佳答案
// 使用单调队列来存储所有备选答案（将不可能的答案移除）

// 明确使用单调队列的话，那么最重要的问题来了，确定什么样的单调关系？
// 可以这么思考
// 定性分析：使用单调递增队列，队首的答案为区间的max，所以至少需要单增
// 定量分析：
// top                              tail
// ------------------------------------
//                    |
//              |     |
//              |     |
//              |     |
// ------------------------------------
// i < j , a[i] < a[j], a[j]更大，肯定是备选答案，a[i]还是备选答案吗？
// 不是，因为滑动窗口移动的时候
// 窗口包含i,j， 答案为a[j]
// 窗口移除i，答案为a[j]
// 所以单调栈中不可能是 < 的关系

// 那么，单调栈中，要么是>关系，要么是>=关系
// 这里应该选哪个呢
// 两者区别就是，相同的备选答案是否保留
// 在区间会不断缩小的情况下，需要保留相同的备选答案
// 否则因为区间缩小而移除的最佳答案，而相同的备选答案未存储在数据结构中，那么上位的备选答案将导致错误答案

// 使用单调队列处理一遍数组，O(N)
// 切分的过程中，花费O(N)，期间可以用O(1)的时间求得左区间和右区间的最大值
// 总时间为O(N)
// 空间复杂度为O(N)
func maxAbs2(nums []int) int {
	var leftMax, rightMax, res int
	leftMax = -1 << 31

	queue := fundamentals.NewDeque()
	// [1, n-1]
	for i := 1; i < len(nums); i++ { // 单调不减队列
		for !queue.IsEmpty() && utils.Less(queue.Tail(), nums[i]) {
			queue.PopBack()
		}
		queue.PushBack(nums[i]) // q is empty || cur <= q.tail
	}

	for i := 0; i < len(nums)-1; i++ {
		// left add i, right remove i+1
		// left part
		if nums[i] > leftMax {
			leftMax = nums[i]
		}
		// right part
		rightMax = queue.Top().(int) // max of [i+1, n-1]
		if rightMax == nums[i+1] {   // 左边界正好是区间的最大值，移除之
			queue.PopFront() // answers of [i+2, n-1]
		}
		res = utils.MaxInt(res, _abs(rightMax-leftMax))
	}
	return res
}

// 从上面的解法中可以发现，左区间的最大值求起来很简单，因为左区间一直是扩张，没有缩减。
// 而右区间有缩减，所以上一个方法，使用mono queue这种数据结构来存储备选答案，
// 当缩减的时候，移除的元素正好是答案时，数据结构中备选答案可以用O(1)的时间上位为答案，从而可以O(1)的时间获得右区间的最大值

// 但是左右区间，其实地位是相等的，只不过习惯从左到右观看，才觉得左右区间不一样
// 如果从右到左看，右区间也是不断扩张而不缩减的。
// 那么可以通过预处理
// 先从左到右遍历一次，先求出所有可能的左区间的最大值
// 然后从右到左遍历一次，求出所有可能的右区间的最大值
// 然后再遍历一次，分别求得划分的左右区间的最大值的差的绝对值

// 时间复杂度为O(N), 遍历3次；空间复杂度O(N)
// 虽然这个方法和上一个方法的复杂度是同一个数量级的，但是上一个方法有着常数级别的优势
func maxAbs3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// left[i]: nums[...i]的max; left[i] = max(left[i-1], nums[i])
	// right[i]: nums[i...]的max; right[i] = max(right[i+1], nums[i])
	left, right := make([]int, len(nums)), make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			left[i] = nums[i]
			continue
		}
		left[i] = utils.MaxInt(left[i-1], nums[i])
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			right[i] = nums[i]
			continue
		}
		right[i] = utils.MaxInt(right[i+1], nums[i])
	}

	var res int
	// [0, i] [i+1, n-1]   i in [0, n-2]
	for i := 0; i <= len(nums)-2; i++ {
		res = utils.MaxInt(res, _abs(right[i+1]-left[i]))
	}
	return res
}

// 这种方式比较trick，时间复杂度不能压缩，但是空间复杂度能够进一步压缩到O(1)
// 首先遍历一遍数组，获取全局max
// 划分中，max要么在左区间，要么在右区间
// 假设max = nums[k]在左区间
// 那么只需要求得右区间的的最小值，
// min(
//		max(nums[k+1, N-1]),
//		max(nums[k+2, N-1]),
//		max(nums[k+3, N-1]),
//      ...
//		max(nums[N-2, N-1]),
//		max(nums[N-1, N-1]),
//		)
//
// 如果区间出现了较大的值，那么外面的min肯定会使其忽略，最后的出min max(...) = nums[N-1]
// res = max - nums[n-1]

// 同理max落在左区间 res = max - nums[0]

func maxAbs4(nums []int) int {
	max := -1 << 31
	for _, v := range nums {
		max = utils.MaxInt(max, v)
	}
	return max - utils.MinInt(nums[len(nums)-1], nums[0])
}

func _abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

//func main() {
//	nums := []int{2, 7, 4, 1, 1}
//	fmt.Println(maxAbs(nums))
//	fmt.Println(maxAbs2(nums))
//	fmt.Println(maxAbs3(nums))
//	fmt.Println(maxAbs4(nums))
//}
