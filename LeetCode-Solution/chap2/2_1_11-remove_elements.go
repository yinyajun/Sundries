/*
Given an array and a value, remove all instances of that value in place and return the new length.
The order of elements can be changed. It doesn’t matter what you leave beyond the new length.

* @Author: Yajun
* @Date:   2021/11/21 15:49
*/

package chap2

// time: O(n); space: O(1)
func removeElements(nums []int, target int) int {
	var index int

	for i := 0; i < len(nums); i++ {
		if nums[i] != target {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
