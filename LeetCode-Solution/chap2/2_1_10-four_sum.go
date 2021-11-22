/*
Given an array S of n integers, are there elements a, b, c, and d in S such that a + b + c + d = target?
Find all unique quadruplets in the array which gives the sum of target.
Note:
• Elements in a quadruplet (a, b, c, d) must be in non-descending order. (ie, a ≤ b ≤ c ≤ d) • The solution set must not contain duplicate quadruplets.
For example, given array S = {1 0 -1 0 -2 2}, and target = 0.
A solution set is:
(-1, 0, 0, 1)
(-2, -1, 1, 2)
(-2, 0, 0, 2)

* @Author: Yajun
* @Date:   2021/11/21 12:40
*/

package chap2

import (
	"sort"
)

// time: O(n^3), space: O(1)
func fourSum(nums []int, target int) [][4]int {
	res := make([][4]int, 0)
	if len(nums) < 4 {
		return res
	}

	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	var (
		a, b, c, d int
		sum        int
	)

	for a = 0; a < len(nums)-3; a++ {
		if a > 0 && nums[a-1] == nums[a] {
			continue
		}
		for b = a + 1; b < len(nums)-2; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			c, d = b+1, len(nums)-1
			for c < d {
				sum = nums[a] + nums[b] + nums[c] + nums[d]
				if sum < target {
					c++
				} else if sum > target {
					d--
				} else {
					if c == b+1 || nums[c-1] != nums[c] {
						res = append(res, [4]int{nums[a], nums[b], nums[c], nums[d]})
					}
					c++
					d--
				}
			}
		}
	}
	return res
}

// time: avg O(n^2), worst O(n^4); space: O(n^2)
func fourSumB(nums []int, target int) [][4]int {
	res := make([][4]int, 0)
	if len(nums) < 4 {
		return res
	}

	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	var (
		cache = make(map[int][][2]int)
		a, b  int
		c, d  int
		sum   int
	)

	for a = 0; a < len(nums)-2; a++ {
		for b = a + 1; b < len(nums)-1; b++ {
			sum = nums[a] + nums[b]
			if _, ok := cache[sum]; !ok {
				cache[sum] = [][2]int{{a, b}}
			} else {
				cache[sum] = append(cache[sum], [2]int{a, b})
			}
		}
	}

	// 遍历所有c，d可能的组合（需要针对c，d去重）
	for c = 2; c < len(nums)-1; c++ {
		if c > 2 && nums[c] == nums[c-1] {
			continue
		}
		for d = c + 1; d < len(nums); d++ {
			if d > c+1 && nums[d] == nums[d-1] { // ！此时d仍然需要去重（遍历所有cd情况）
				continue
			}
			sum = nums[c] + nums[d]
			if t, ok := cache[target-sum]; !ok {
				continue
			} else {
				for _, tt := range t {
					if c < tt[1] {
						continue
					}
					// c >= second
					res = append(res, [4]int{nums[tt[0]], nums[tt[1]], nums[c], nums[d]})
				}
			}
		}
	}
	return res
}
