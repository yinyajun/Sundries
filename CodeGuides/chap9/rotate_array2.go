package main

// 在旋转数组中寻找一个数

func RotateArray2(arr []int, target int) bool {
	lo, hi := 0, len(arr)-1
	var mid int

	for lo <= hi {
		mid = lo + (hi-lo)/2
		if arr[mid] == target {
			return true
		}

		if arr[lo] != arr[mid] {
			if arr[mid] > arr[lo] { // [lo, mid] increasing
				if arr[lo] <= target && target < arr[mid] { // target 在单调区间上
					hi = mid - 1
				} else { // target 不在单调区间上
					lo = mid + 1
				}
			} else { // arr[mid] < arr[lo], [lo, mid] breaking point
				// [mid, hi] increasing
				if arr[mid] < target && target <= arr[hi] {
					lo = mid + 1
				} else {
					hi = mid - 1
				}
			}
		} else if arr[mid] != arr[hi] {
			if arr[mid] < arr[hi] { // [mid, hi] increasing
				if arr[mid] < target && target <= arr[hi] { // target 在单调区间上
					lo = mid + 1
				} else { // target不在大拿掉区间
					hi = mid - 1
				}
			} else { // arr[mid] > arr[hi] , [mid, hi] breaking point
				// [lo, mid] increasing
				if arr[lo] <= target && target < arr[mid] {
					hi = mid - 1
				} else {
					lo = mid + 1
				}
			}
		} else { //arr[lo] == arr[mid] == arr[hi]
			for lo < mid {
				if arr[lo] != arr[mid] {
					break
				}
				// arr[lo] == mid
				lo++
			}

			if lo == mid {
				lo = mid + 1 // 关键！否则搜索区间将可能永远是有单个元素
			}
		}
	}
	return false

}

//func main() {
//	arr := []int{4, 5, 6, 7, 1, 2, 3}
//	fmt.Println(RotateArray2(arr, 8))
//
//	arr = []int{2, 2, 2, 2, 2, 3, 2, 2}
//	fmt.Println(RotateArray2(arr, 2))
//
//	arr = []int{2, 2, 1, 2, 2, 2, 2, 2}
//	fmt.Println(RotateArray2(arr, 0))
//}
