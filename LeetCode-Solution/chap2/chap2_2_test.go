/*
* @Author: Yajun
* @Date:   2021/11/27 22:42
 */

package chap2

import (
	"fmt"
	"solution/utils"
	"testing"
)

func Test2_2_1(t *testing.T) {
	l1 := utils.NewListFromSlice([]int{2, 4, 3})
	l2 := utils.NewListFromSlice([]int{5, 6, 4, 1})
	res := addTwoNumbers(l1, l2)
	utils.ShowList(res)
}

func Test2_2_2(t *testing.T) {
	c := []int{1, 2, 3, 4, 5}
	l := utils.NewListFromSlice(c)
	utils.ShowList(reverseBetween(l, 2, 4))

	l = utils.NewListFromSlice(c)
	utils.ShowList(reverseBetweenB(l, 2, 4))

	l = utils.NewListFromSlice(c)
	utils.ShowList(reverseBetweenC(l, 2, 4))

	a := []int{1, 2, 3, 4, 5}
	l = utils.NewListFromSlice(a)
	utils.ShowList(reverseList(l))

	l = utils.NewListFromSlice(a)
	utils.ShowList(reverseListB(l))

	l = utils.NewListFromSlice(a)
	utils.ShowList(reverseListC(l))

	l = utils.NewListFromSlice(a)
	utils.ShowList(reverseListD(l))

	l = utils.NewListFromSlice(a)
	utils.ShowList(reverseListE(l))

	l = utils.NewListFromSlice(a)
	utils.ShowList(reverseListF(l))

	b := []int{1, 2, 3, 4, 5}
	l = utils.NewListFromSlice(b)
	utils.ShowList(reverseKList(l, 3))

	l = utils.NewListFromSlice(b)
	utils.ShowList(reverseKListB(l, 3))

	l = utils.NewListFromSlice(b)
	utils.ShowList(reverseKListC(l, 3))
}

func Test2_2_3(t *testing.T) {
	l := utils.NewListFromSlice([]int{1, 4, 3, 2, 5, 2})
	utils.ShowList(partitionList(l, 3))
}

func Test2_2_4(t *testing.T) {
	l := utils.NewListFromSlice([]int{1, 1, 2, 3, 3})
	utils.ShowList(deleteDuplicate(l))

	l = utils.NewListFromSlice([]int{1, 1, 2, 3, 3})
	utils.ShowList(deleteDuplicateB(l))

	l = utils.NewListFromSlice([]int{1, 1, 2, 3, 3})
	utils.ShowList(deleteDuplicateC(l))

	l = utils.NewListFromSlice([]int{1, 1, 2, 3, 3})
	utils.ShowList(deleteDuplicateD(l))
}

func Test2_2_5(t *testing.T) {
	l := utils.NewListFromSlice([]int{1, 1, 1, 2, 3, 4, 4, 4})
	utils.ShowList(deleteDuplicates2(l))

	l = utils.NewListFromSlice([]int{1, 1, 1, 2, 3, 4, 4, 4})
	utils.ShowList(deleteDuplicates2B(l))
}

func Test2_2_6(t *testing.T) {
	l := utils.NewListFromSlice([]int{1, 2, 3, 4, 5})
	utils.ShowList(rotateRight(l, 2))
}

func Test2_2_7(t *testing.T) {
	l := utils.NewListFromSlice([]int{1, 2, 3, 4, 5, 6})
	utils.ShowList(removeNthNode(l, 3))
}

func Test2_2_8(t *testing.T) {
	l := utils.NewListFromSlice([]int{1, 2, 3, 4, 5, 6})
	utils.ShowList(swapNodesInPairs(l))

	l = utils.NewListFromSlice([]int{1, 2, 3, 4, 5, 6})
	utils.ShowList(swapNodesInPairsB(l))

	l = utils.NewListFromSlice([]int{1, 2, 3, 4, 5, 6})
	utils.ShowList(swapNodesInPairsC(l))
}

func Test2_2_9(t *testing.T) {
	l := utils.NewListFromSlice([]int{1, 2, 3, 4, 5, 6, 7})
	utils.ShowList(reverseNodesInGroup(l, 3))

	l = utils.NewListFromSlice([]int{1, 2, 3, 4, 5, 6, 7})
	utils.ShowList(reverseNodesInGroupB(l, 3))
}

func Test2_2_10(t *testing.T) {
	n1 := &RandomNode{val: 1}
	n2 := &RandomNode{val: 2}
	n3 := &RandomNode{val: 3}
	n4 := &RandomNode{val: 4}

	n1.next = n2
	n2.next = n3
	n3.next = n4
	n1.random = n4
	n3.random = n2

	for cur := n1; cur != nil; cur = cur.next {
		fmt.Print(cur)
	}
	fmt.Println()

	c := copyRandomList(n1)

	for cur := c; cur != nil; cur = cur.next {
		fmt.Print(cur)
	}
	fmt.Println()
}
