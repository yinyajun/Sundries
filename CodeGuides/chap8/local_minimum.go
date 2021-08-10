package main

// 暴力解法，通过遍历以O(N)的复杂度找出local minimum
func LocalMinimum(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}

	// len(nums) >= 2
	for i := 0; i < len(nums); i++ {
		if i == 1 && nums[1] > nums[0] {
			return 0
		}
		if i == len(nums)-1 && nums[i] < nums[i-1] { // len(nums)>=2, nums[i-1] is valid
			return i
		}
		if i-1 >= 0 && i+1 < len(nums) && nums[i] < nums[i-1] && nums[i] < nums[i+1] { // 确保i+1和i-1合法
			return i
		}
	}
	return -1
}

// 某侧只有全都比mid大的时候，才不会有local minimum
// 可以发现，二分查找不一定需要数组有序，只要能确定二分两侧的某一侧一定有要查找的内容，就可以用二分查找。
func LocalMinimum3(nums []int) int {
	lo, hi := 0, len(nums)-1
	var mid int
	for lo < hi { // [lo, hi]，搜索区间至少有两个元素
		mid = lo + (hi-lo)/2
		if mid-1 >= 0 && nums[mid] > nums[mid-1] { // 左侧必然有local minimum
			hi = mid - 1
		} else if mid+1 < len(nums) && nums[mid] > nums[mid+1] { // 右侧必然有local minimum
			lo = mid + 1
		} else { // 两端 或者是  nums[mid]比两侧小
			return mid
		}
	}
	// lo == hi, 搜索区间只有一个元素
	return hi
}

//func main() {
//	nums := []int{3, 2, 5, 4, 6}
//	//nums := []int{1,2,3,4,5}
//	//nums := []int{5,4,3,2,1}
//	index := LocalMinimum3(nums)
//	fmt.Println(nums[index])
//}
