package chap2

import (
	"fmt"
	"testing"
)

// go test -v -run Test2_1_1
func Test2_1_1(t *testing.T) {
	nums := []int{1, 1, 2, 3, 4, 4, 5, 6}
	idx := removeDuplicatesB(nums)
	fmt.Println(idx, nums[:idx])
}

func Test2_1_2(t *testing.T) {
	nums := []int{5, 5, 5, 5, 6, 6, 6, 7, 8, 9, 9}
	idx := removeDuplicates2c(nums)
	fmt.Println(idx, nums[:idx])
}

func Test2_1_3(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	//nums := []int {0,1,2,3,4,5,6}
	fmt.Println(searchD(nums, 3))
}

func Test2_1_4(t *testing.T) {
	nums := []int{2, 2, 2, 2, 2, 3, 2, 2}
	fmt.Println(search2D(nums, 1))
}

func Test2_1_5(t *testing.T) {
	a := []int{1, 3, 5, 7}
	b := []int{2, 4, 6, 8}
	fmt.Println(findMedianSortedArraysD(a, b))
}

func Test2_1_6(t *testing.T) {
	//nums := []int{100,4,200,1,3,2}
	nums := []int{1, 2, 3, 4, 5, 6}
	//fmt.Println(longestConsecutive(nums))
	fmt.Println(longestConsecutiveB(nums))
	fmt.Println(longestConsecutiveC(nums))
	fmt.Println(longestConsecutiveD(nums))
}

func Test2_1_7(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSumA(nums, target))
}

func Test2_1_8(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	target := 0
	fmt.Println(threeSum(nums, target))
}

func Test2_1_9(t *testing.T) {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}

func Test2_1_10(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(fourSumB(nums, target))
}

func Test2_1_11(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	c := removeElements(nums, target)
	fmt.Println(nums[:c])
}

func Test2_1_12(t *testing.T) {
	nums := []int{1, 2, 7, 4, 3, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}

func Test2_1_13(t *testing.T) {
	//fmt.Println(permutationSequence(3, 3))
	fmt.Println(permutationSequenceB(3, 3))
}

func Test2_1_14(t *testing.T) {
	board1 := [9][9]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	board2 := [9][9]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board1))
	fmt.Println(isValidSudoku(board2))
}

func Test2_1_15(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trappingRainWater(height))
	fmt.Println(trappingRainWaterB(height))
	fmt.Println(trappingRainWaterC(height))
	fmt.Println(trappingRainWaterD(height))
}
