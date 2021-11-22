/*
Given an array of integers, find two numbers such that they add up to a specific target number.
The function twoSum should return indices of the two numbers such that they add up to the target, where
index1 must be less than index2. Please note that your returned answers (both index1 and index2) are not
zero-based.

* @Author: Yajun
* @Date:   2021/10/19 10:10
*/

package chap2

func twoSumA(nums []int, target int) [2]int {
	meets := make(map[int]int)
	result := [2]int{}

	for idx, n := range nums {
		if index, ok := meets[target-n]; ok {
			result[0] = index + 1
			result[1] = idx + 1
			return result
		}
		meets[n] = idx
	}
	return result
}

// 先排序，然后左右夹逼，时间复杂度为O(NlogN)。但是此题需要返回的是下标，而排序会打乱下标
