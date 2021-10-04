/*
Suppose a sorted array is rotated at some pivot unknown to you beforehand.
(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).
You are given a target value to search. If found in the array return its index, otherwise return -1.
You may assume no duplicate exists in the array.

* @Author: Yajun
* @Date:   2021/10/4 10:40
*/

package chap2

// time: O(log n); space: O(1)
func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1 // [lo, hi]
	var mid int
	for lo <= hi {
		mid = lo + (hi-lo)/2

		if target == nums[mid] {
			return mid
		}

		// 不允许有重复元素，可以直接判断单调区间
		if nums[lo] <= nums[mid] { // left part is increasing
			if target >= nums[lo] && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else { // right part is increasing
			if target > nums[mid] && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}

func searchB(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	var mid int
	for lo <= hi {
		mid = lo + (hi-lo)/2

		if target == nums[mid] {
			return mid
		}

		if nums[mid] > nums[hi] { // right part存在断点
			// left part单调
			if target >= nums[lo] && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else if nums[lo] > nums[mid] { // left part存在断点
			// right part单调
			if target > nums[mid] && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else { // nums[lo] <= nums[mid] <= nums[hi]
			// 由于不存在相等元素，[lo, hi]单调，不存在断点
			if target > nums[mid] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}

func searchC(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	var mid int
	for lo <= hi {
		mid = lo + (hi-lo)/2

		if target == nums[mid] {
			return mid
		}

		if nums[mid] > nums[hi] { // right part存在断点
			// left part单调
			if target >= nums[lo] && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else { // left part 可能存在断点
			// right part单调
			if target > nums[mid] && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}

func searchD(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	var mid int
	for lo <= hi {
		mid = lo + (hi-lo)/2

		if target == nums[mid] {
			return mid
		}
		if nums[lo] > nums[mid] { // left part存在断点
			// right part单调
			if target > nums[mid] && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else { // right part可能存在断点
			// left part单调
			if target >= nums[lo] && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
	}
	return -1
}
