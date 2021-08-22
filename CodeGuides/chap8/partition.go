package main

// 给定有序数组arr，调整使其左半部分没有重复元素且升序，不用管右部分
// 时间O(N)，空间O(1)

// 错误：判断条件错误，应该判断arr[i] == arr[j]，而不是前一个元素
//func partitionAdjust(arr []int) []int {
//	if len(arr) == 0 {
//		return arr
//	}
//	// len(arr)>=1
//	j, i := 1, 1
//	for ; i < len(arr); i++ {
//		if arr[i] != arr[i-1]{
//			arr[j], arr[i] = arr[i], arr[j]
//			j++
//		}
//	}
//	return arr
//}

// 类似于快排partition的过程
// a[1...j)不重复有序，[j,i)未处理
func partitionAdjust(arr []int) {
	if len(arr) < 2 {
		return
	}
	// len(arr)>=1
	j, i := 1, 1
	for ; i < len(arr); i++ {
		if arr[i] != arr[j-1] {
			arr[j], arr[i] = arr[i], arr[j]
			j++
		}
	}
}

// 补充问题，arr其中只含有0，1，2三个值，实现arr排序

// [0, lt)==0  [lt, i) ==1  (gt, hi] ==2
func sortPartition(arr []int) {
	lt, gt := 0, len(arr)-1
	i := 0

	for i <= gt {
		if arr[i] == 0 {
			arr[i], arr[lt] = arr[lt], arr[i]
			lt++
			i++
		} else if arr[i] == 2 {
			arr[i], arr[gt] = arr[gt], arr[i]
			gt--
		} else { // arr[i] ==1
			i++
		}
	}
}

//func main() {
//	arr := []int{1, 2, 2, 2, 3, 3, 4, 5, 6, 6, 7, 7, 8, 8, 8, 9, 9}
//	partitionAdjust(arr)
//	fmt.Println(arr)
//
//
//	arr = []int{0,2,1,2,1,2,1,0}
//	sortPartition(arr)
//	fmt.Println(arr)
//}
