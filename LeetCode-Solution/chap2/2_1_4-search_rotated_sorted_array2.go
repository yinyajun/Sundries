/*
Follow up for ”Search in Rotated Sorted Array”: What if duplicates are allowed?
Would this affect the run-time complexity? How and why?
Write a function to determine if a given target is in the array.

* @Author: Yajun
* @Date:   2021/10/4 13:10
*/

package chap2

// time: O(n), space: O(1)
func search2(nums []int, target int) bool {
	lo, hi := 0, len(nums)-1 // [lo, hi]
	var mid int

	for lo <= hi {
		mid = (hi-lo)/2 + lo
		if nums[mid] == target {
			return true
		}

		if nums[lo] < nums[mid] { // [lo, mid] increasing
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else if nums[mid] < nums[hi] { // [mid, hi] increasing
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else { // nums[lo]>=nums[mid]>=nums[hi]
			lo++ // skip duplicate one
		}
	}
	return false
}

func search2B(nums []int, target int) bool {
	lo, hi := 0, len(nums)-1 // [lo, hi]
	var mid int

	for lo <= hi {
		mid = (hi-lo)/2 + lo
		if nums[mid] == target {
			return true
		}
		if nums[lo] < nums[mid] { // [lo, mid] increasing
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else if nums[lo] > nums[mid] { // 断点在[lo, mid], [mid, hi] increasing
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else { // nums[lo]==nums[mid]
			lo++ // skip duplicate one
		}
	}
	return false
}

func search2C(nums []int, target int) bool {
	lo, hi := 0, len(nums)-1 // [lo, hi]
	var mid int

	for lo <= hi {
		mid = (hi-lo)/2 + lo
		if nums[mid] == target {
			return true
		}
		if nums[mid] > nums[hi] { // 断点在[mid, hi], [lo, mid] increasing
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else if nums[lo] > nums[mid] { // 断点在[lo, mid], [mid, hi] increasing
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else { // nums[lo] <= nums[mid] <= nums[hi]
			if nums[lo] < nums[hi] { // 未旋转
				if target > nums[mid] {
					lo = mid + 1
				} else {
					hi = mid - 1
				}
			} else { // nums[lo] == nums[mid] == nums[hi]
				lo++ // skip duplicate one
			}
		}
	}
	return false
}

// 在nums[lo] == nums[mid]的情况下，尽可能再使用二分
func search2D(nums []int, target int) bool {
	lo, hi := 0, len(nums)-1 // [lo, hi]
	var mid int

	for lo <= hi {
		mid = (hi-lo)/2 + lo
		if nums[mid] == target {
			return true
		}
		if nums[lo] < nums[mid] { // [lo, mid] increasing
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else if nums[lo] > nums[mid] { // 断点在[lo, mid], [mid, hi] increasing
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else { // nums[lo]==nums[mid]
			// for lo' = lo; lo'< mid; lo'++
			// 直接用lo代替lo'，nums[mid]可以代替nums[lo]
			for lo <= mid { // 关键的等号，不然搜索区间可能永远会是[mid, mid]
				if nums[lo] == nums[mid] {
					lo++
				} else {
					//nums[lo] < nums[mid]  ==> [lo', mid-1] increasing
					//nums[lo] > nums[mid]， [lo', mid-1]出现断点
					hi = mid - 1
					break
				}
			}
		}
	}
	return false
}

func search2E(nums []int, target int) bool {
	lo, hi := 0, len(nums)-1 // [lo, hi]
	var mid int

	for lo <= hi {
		mid = (hi-lo)/2 + lo
		if nums[mid] == target {
			return true
		}

		if nums[lo] < nums[mid] { // [lo, mid] increasing
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else if nums[mid] < nums[hi] { // [mid, hi] increasing
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else { // nums[lo]>=nums[mid]>=nums[hi]
			for lo <= mid {
				if nums[lo] == nums[mid] {
					lo++
				} else {
					// nums[lo] != nums[mid], [lo, mid-1]递增或者出现断点
					hi = mid - 1
					break
				}
			}
		}
	}
	return false
}
