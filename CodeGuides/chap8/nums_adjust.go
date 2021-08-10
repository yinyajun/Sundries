package main

// 奇数下标都是奇数，偶数下标都是偶数

// 用odd，even两个指针标志这数组中奇偶的下标
// 如果数组中最后一个数，是奇数，那么向odd下标发送并交换。反之亦然
// 用位运算判断奇偶
func numAdjust(nums []int) {
	if len(nums) < 2 {
		return
	}

	even, odd := 0, 1
	for even < len(nums) && odd < len(nums) {
		if nums[len(nums)-1]&1 == 0 { // even
			nums[even], nums[len(nums)-1] = nums[len(nums)-1], nums[even]
			even += 2
		} else {
			nums[odd], nums[len(nums)-1] = nums[len(nums)-1], nums[odd]
			odd += 2
		}
	}
}

//func main() {
//	a := []int{1, 8, 3, 2, 4, 6}
//	numAdjust(a)
//	fmt.Println(a)
//}
