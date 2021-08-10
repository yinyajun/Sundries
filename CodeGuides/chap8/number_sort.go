package main

func numberSort1(nums []int) {
	// 每次遍历，保证i位置上的元素正确
	for i := 0; i < len(nums); i++ {
		tmp := nums[i]
		for nums[i] != i+1 {
			tmp, nums[tmp-1] = nums[tmp-1], tmp // 这里本质上可以认为是临时变量tmp和nums[tmp]的交换
			//next := nums[tmp-1]
			//nums[tmp-1] = tmp
			//tmp = next
		}
	}
}

func numberSort2(nums []int) {
	// 每次遍历，都保证nums[i]位置上的元素正确
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
}

//func main() {
//	a := []int{1, 2, 6, 3, 5, 4}
//	numberSort2(a)
//	fmt.Println(a)
//}
