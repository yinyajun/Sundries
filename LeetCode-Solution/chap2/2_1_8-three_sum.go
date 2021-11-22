/*
Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? Find all unique
triplets in the array which gives the sum of zero.
Note:
• Elements in a triplet (a, b, c) must be in non-descending order. (ie, a ≤ b ≤ c) • The solution set must not contain duplicate triplets.
For example, given array S = {-1 0 1 2 -1 -4}.
A solution set is:
(-1, 0, 1)
(-1, -1, 2)

* @Author: Yajun
* @Date:   2021/10/19 10:22
*/

package chap2

import (
	"sort"
)

// time: O(n^2), space: O(1)
// 先排序，可以保证结果是non-descending order
func threeSum(nums []int, target int) [][3]int {
	res := make([][3]int, 0)
	if len(nums) < 3 {
		return res
	}

	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	var (
		a, b, c int
		sum     int
	)

	for a = 0; a < len(nums)-2; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		b, c = a+1, len(nums)-1
		for b < c {
			sum = nums[a] + nums[b] + nums[c]
			if sum < target {
				b++
			} else if sum > target {
				c--
			} else {
				if b == a+1 || nums[b] != nums[b-1] {
					res = append(res, [3]int{nums[a], nums[b], nums[c]})
				}
				b++
				c--
			}
		}
	}
	return res
}
