/*
* @Author: Yajun
* @Date:   2021/11/27 22:42
 */

package chap2

import (
	"solution/utils"
	"testing"
)

func Test2_2_1(t *testing.T) {
	l1 := utils.NewListFromSlice([]int{2, 4, 3})
	l2 := utils.NewListFromSlice([]int{5, 6, 4, 1})
	res := addTwoNumbers(l1, l2)
	utils.ShowListNode(res)
}

func Test2_2_2(t *testing.T) {
	l := utils.NewListFromSlice([]int{1, 2, 3, 4, 5})
	utils.ShowListNode(reverseBetween(l, 2, 4))
	l = utils.NewListFromSlice([]int{1, 2, 3, 4, 5})
	utils.ShowListNode(reverseBetweenB(l, 2, 4))

	a := []int{1, 2, 3, 4, 5}
	l = utils.NewListFromSlice(a)
	utils.ShowListNode(reverseList(l))

	l = utils.NewListFromSlice(a)
	utils.ShowListNode(reverseListB(l))

	l = utils.NewListFromSlice(a)
	utils.ShowListNode(reverseListC(l))

	b := []int{1, 2, 3, 4, 5}
	l = utils.NewListFromSlice(b)
	utils.ShowListNode(reverseKList(l, 3))

	l = utils.NewListFromSlice(b)
	utils.ShowListNode(reverseKListB(l, 3))
}
