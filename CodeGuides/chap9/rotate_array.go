package main

import (
	"CodeGuide/base/utils"
)

// 将有序数组旋转，找到旋转数组中的最小值
// 数组元素可能重复

// 如何使用二分法来压缩时间？

// arr[lo...hi]
// 1. arr[lo] < arr[hi]：符合单调，那么arr[lo]是min
// 2. arr[mid] < arr[lo], [lo...mid]之间不符合单调顺序，存在断点
// 3. arr[hi] < arr[mid], [mid...hi]之间不符合单调顺序，存在端点
// 上面的分类并不是完备的，还有情况就是
// arr[lo] == arr[mid] == arr[hi]
// 此时不能通过mid来缩减搜索区间

// 此时只能遍历搜索了吗？
// i = lo, 向右遍历
// 1. arr[i] < arr[lo], 出现降序，必为断点
// 2. arr[i] > arr[lo], arr[i] > arr[mid]，出现降序，[i, mid]之间存在断点
// 3. arr[i] == arr[lo]，只能继续遍历下去，遍历到mid-1，此时[lo,mid)上不可能有断点，断点在[mid, hi]上

func rotateArray(arr []int) int {
	if len(arr) < 2 { // 至少需要两个元素，才能判断相对关系
		panic("invalid array length")
	}
	// 搜索区间[lo, hi]
	lo, hi := 0, len(arr)-1
	var mid int

	for lo < hi { // 搜索区间至少包含两个元素
		if hi-lo == 1 { // 超级关键！否则无限循环，因为[lo,hi]至少有两个元素，
			// 而下面的分支缩减的区间也是包含mid，当区间只有两个元素的时候，区间不会进一步缩小
			break
		}

		// 已经单调，直接返回最小
		if arr[lo] < arr[hi] {
			return arr[lo]
		}

		mid = lo + (hi-lo)/2
		if arr[mid] < arr[lo] {
			hi = mid // [lo, mid]
		} else if arr[mid] > arr[hi] {
			lo = mid // [mid, hi]
		} else { // arr[lo] == arr[mid] == arr[hi]
			for lo < mid {
				if arr[lo] == arr[mid] {
					lo++
				} else if arr[lo] < arr[mid] {
					return arr[lo]
				} else {
					hi = mid // [lo, mid]间存在断点
					break    // 关键！！
				}
			}
		}
	}
	return utils.MinInt(arr[lo], arr[hi])
}

//func main() {
//	arr := []int{1, 2, 3, 4, 5}
//	fmt.Println(rotateArray(arr))
//
//	arr = []int{1, 2, 3, 4, 5, 0}
//	fmt.Println(rotateArray(arr))
//
//	arr = []int{2, 2, 2, 2, 2}
//	fmt.Println(rotateArray(arr))
//}
