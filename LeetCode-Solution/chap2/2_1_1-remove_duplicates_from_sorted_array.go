/*

Given a sorted array, remove the duplicates in place such that each element appear only once
and return the new length.
Do not allocate extra space for another array, you must do this in place with constant memory.
For example, Given input array A = [1,1,2],
Your function should return length = 2, and A is now [1,2].

* @Author: Yajun
* @Date:   2021/10/3 14:46
*/

package chap2

// time: O(n); space: O(1)
// 注意nums[i]和nums[idx]比较，而不是nums[i-1]比较，尽管这里不影响
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var idx int // [0, idx]不重复，至少有1个元素不重复
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[idx] {
			continue
		}
		idx++
		nums[idx] = nums[i]
	}
	return idx + 1
}

func removeDuplicatesB(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var idx = 1 // [0, idx) no duplicate
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[idx-1] {
			continue
		}
		nums[idx] = nums[i]
		idx++
	}
	return idx
}
