package main

// 判断一个数组是否bst的后序遍历结果

func IsPostArray(arr []int) bool {
	if len(arr) == 0 {
		return false
	}
	return _isPostArr(arr, 0, len(arr)-1)
}

// 后序遍历：左右根
func _isPostArr(arr []int, lo, hi int) bool {
	if lo >= hi {
		return true
	}
	val := arr[hi]
	less, more := lo-1, hi
	// 分界点是最后一个小于val的
	// [lo, less] < ，[more, hi-1] >=
	for i := lo; i < hi; i++ {
		if arr[i] < val {
			less = i
		} else { // arr[i] >= val
			more = i
			break
		}

	}
	for j := more + 1; j < hi; j++ {
		if arr[j] < val {
			return false
		}
	}
	return _isPostArr(arr, lo, less) && _isPostArr(arr, more, hi-1)
}
