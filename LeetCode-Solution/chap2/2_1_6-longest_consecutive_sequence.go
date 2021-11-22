/*
Given an unsorted array of integers, find the length of the longest consecutive elements sequence.
For example, Given [100, 4, 200, 1, 3, 2], The longest consecutive elements sequence is [1,
2, 3, 4]. Return its length: 4.
Your algorithm should run in O(n) complexity.

* @Author: Yajun
* @Date:   2021/10/5 15:40
*/

package chap2

import (
	"solution/utils"
	"sort"
)

// time: O(nlogn); space: O(logn) 取决于排序
// 时间超出要求，不符合题意
func longestConsecutive(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	var cnt int
	var ans int

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1]+1 {
			cnt++
		} else { // i ==0 || nums[i] not consecutive
			cnt = 1
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return ans
}

// time: O(n); space: O(n)
func longestConsecutiveB(nums []int) int {
	var ans int
	starts := make(map[int]int) // key: 区间开头， value: 区间末尾
	ends := make(map[int]int)   // key: 区间末尾， value：区间开头

	union := func(x, y int) int { // union [a, x][y, b], return length of new region [a, b]
		starts[ends[x]] = starts[y]
		ends[starts[y]] = ends[x]
		ans := starts[y] - ends[x] + 1
		delete(ends, x)
		delete(starts, y)
		return ans
	}
	for _, n := range nums {
		if _, ok := starts[n]; ok { //
			continue
		}
		if _, ok := ends[n]; ok {
			continue
		}
		starts[n] = n
		ends[n] = n
		if _, ok := ends[n-1]; ok { // union [s, n-1] [n, e]
			ans = utils.MaxInt(ans, union(n-1, n))
		}
		if _, ok := starts[n+1]; ok { // union [s, n] [n+1, e]
			ans = utils.MaxInt(ans, union(n, n+1))
		}
	}
	return ans
}

// time: O(n); space: O(n)
// B中使用两个map来维护同一个区间，其实可以规约到同一个map中，区间的某一个边界 + 区间长度 可以确定一个区间。
func longestConsecutiveC(nums []int) int {
	var ans int
	region := make(map[int]int) // key：boundary(left or right), value: region length

	union := func(x, y int) int { // union [a, x][y, b], return length of new region [a, b]
		upper := y + region[y] - 1
		lower := x - region[x] + 1
		ans := upper - lower + 1
		region[upper] = ans
		region[lower] = ans
		return ans
	}

	for _, n := range nums {
		if _, ok := region[n]; ok {
			continue
		}
		region[n] = 1
		if _, ok := region[n-1]; ok {
			ans = utils.MaxInt(ans, union(n-1, n))
		}
		if _, ok := region[n+1]; ok {
			ans = utils.MaxInt(ans, union(n, n+1))
		}
	}
	return ans
}

// time: O(n^2); space: O(n)
// solution上给出time为O(n)，但是对于连续数组，最坏情况应该是O(n^2)
func longestConsecutiveD(nums []int) int {
	used := make(map[int]bool)
	var ans int

	for _, n := range nums {
		if used[n] {
			continue
		}

		used[n] = true
		length := 1

		for j := n + 1; ; j++ {
			if _, ok := used[j]; !ok {
				break
			}
			length++
		}

		for j := n - 1; ; j-- {
			if _, ok := used[j]; !ok {
				break
			}
			length++
		}
		ans = utils.MaxInt(ans, length)
	}
	return ans
}
