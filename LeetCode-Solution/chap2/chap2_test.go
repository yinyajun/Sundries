package chap2

import (
	"fmt"
	"testing"
)

// go test -v -run Test2_1_1
func Test2_1_1(t *testing.T) {
	nums := []int{1, 1, 2, 3, 4, 4, 5, 6}
	idx := removeDuplicates(nums)
	fmt.Println(idx, nums[:idx])
}

func Test2_1_2(t *testing.T) {
	nums := []int{5, 5, 5, 5, 6, 6, 6, 7, 8, 9, 9}
	idx := removeDuplicates2c(nums)
	fmt.Println(idx, nums[:idx])
}


func Test2_1_3(t *testing.T){
	nums := []int {4,5,6,7,0,1,2}
	//nums := []int {0,1,2,3,4,5,6}
	fmt.Println(searchD(nums, 3))
}

func Test2_1_4(t *testing.T){
	nums := []int {2, 2, 2, 2, 2, 3, 2, 2}
	fmt.Println(search2D(nums, 1))
}