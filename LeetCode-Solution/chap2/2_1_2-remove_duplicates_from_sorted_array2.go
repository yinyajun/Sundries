/*
Follow up for ”Remove Duplicates”: What if duplicates are allowed at most twice?
For example, Given sorted array A = [1,1,1,2,2,3],
Your function should return length = 5, and A is now [1,1,2,2,3]

* @Author: Yajun
* @Date:   2021/10/3 15:18
*/

package chap2

// time: O(n); space: O(1)
// 注意nums[i]和nums[idx-1]比较，而不是nums[i-2]比较, 因为nums[i-2]可能已经被修改
func removeDuplicates2(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	var idx = 1 // [0, idx]中元素最多重复两次

	for i := 2; i < len(nums); i++ {
		if nums[i] == nums[idx-1] {
			continue
		}
		idx++
		nums[idx] = nums[i]
	}
	return idx + 1
}

// time: O(n); space: O(1)
// 使用nums[i]和nums[i-k]比较是危险的，因为nums[i-k]可能已经修改了
// 但是当k=1时候是安全的，因为idx<i，还没有来得及修改nums[i-1]
func removeDuplicates2b(nums []int) int {
	var idx int // [0, idx)中元素最多重复两次

	for i := 0; i < len(nums); i++ {
		if i > 0 && i < len(nums)-1 && nums[i] == nums[i-1] && nums[i] == nums[i+1] { // 连续相同元素只保留两端
			continue
		}
		nums[idx] = nums[i]
		idx++
	}
	return idx
}

// time: O(n); space: O(1)
// 由于有序，使用cnt变量记录当前元素的重复次数
// 扩展特别方便
func removeDuplicates2c(nums []int) int {
	var idx, cnt int

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			cnt++
		} else { // i==0 || nums[i] != nums[i-1]
			cnt = 1
		}
		if cnt <= 2 {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}
