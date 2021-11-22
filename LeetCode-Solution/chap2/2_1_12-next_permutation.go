/*
Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.
If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).
The replacement must be in-place, do not allocate extra memory.
Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the
right-hand column.
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1


* @Author: Yajun
* @Date:   2021/11/21 15:55
*/

package chap2

/*
1 2 3
1 3 2
2 1 3
2 3 1
3 1 2
3 2 1

1 2 3 -> 1 3 2

https://www.cnblogs.com/grandyang/p/4428207.html
1 2 7 4 3 1 -> 1 3 1 2 4 7

规律：
1. 从末尾往前看，数字逐渐变大，到了2时才减小的
2. 从后往前寻找第一个比2大的数字，是3
3. 交换2和3，然后将3后面的数字反转一下

直觉：
1. 初始状态是全升序，末尾状态是全降序
2. 寻找子序列是末尾状态，在子序列中寻找大于节点的值交换作为新节点，然后reverse子序列，变成初始状态
*/

// time: O(n); space: O(1)
func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}

	var (
		i, j int
	)
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	for j = len(nums) - 1; j > i; j-- {
		if nums[j] > nums[i] {
			break
		}
	}
	nums[i], nums[j] = nums[j], nums[i]
	reverse(nums, i+1, len(nums)-1)
}

func reverse(nums []int, lo, hi int) {
	for lo < hi {
		nums[lo], nums[hi] = nums[hi], nums[lo]
		lo++
		hi--
	}
}
