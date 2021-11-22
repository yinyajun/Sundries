/*
Given an array S of n integers, find three integers in S such that the sum is closest to a given number,
target. Return the sum of the three integers. You may assume that each input would have exactly one solution.
For example, given array S = {-1 2 1 -4}, and target = 1.
The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

* @Author: Yajun
* @Date:   2021/11/21 12:18
*/

package chap2

import (
	"solution/utils"
	"sort"
)

func threeSumClosest(nums []int, target int) (res int) {
	if len(nums) < 3 {
		panic("invalid nums length")
	}

	minGap := 1 << 32
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	var (
		b, c int
		sum  int
		gap  int
	)

	for a := 0; a < len(nums)-2; a++ {
		b, c = a+1, len(nums)-1
		for b < c {
			sum = nums[a] + nums[b] + nums[c]
			gap = utils.AbsInt(sum - target)

			if gap < minGap {
				minGap = gap
				res = sum
			}

			if sum < target {
				b++
			} else {
				c--
			}
		}
	}
	return res
}
