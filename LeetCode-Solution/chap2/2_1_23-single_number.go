/*
Given an array of integers, every element appears twice except for one. Find that single one.
Note: Your algorithm should have a linear runtime complexity. Could you implement it without using
extra memory?

* @Author: Yajun
* @Date:   2021/11/27 20:38
*/

package chap2

import "solution/utils"

// time : O(n); space: O(1)
// a ^ a ^ b = b
func singleNumber(nums []int) int {
	utils.Assert(len(nums) > 0)
	var res = nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
