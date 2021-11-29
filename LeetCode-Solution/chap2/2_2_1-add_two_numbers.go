/*
You are given two linked lists representing two non-negative numbers. The digits are stored in reverse
order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8

* @Author: Yajun
* @Date:   2021/11/27 22:02
*/

package chap2

import (
	"solution/utils"
)

// time: O(m+n); space: O(1)
// 好久没写链表了，忘了在链表上移动指针
func addTwoNumbers(l1, l2 *utils.ListNode) *utils.ListNode {
	var (
		dummyHead        = new(utils.ListNode)
		pre              *utils.ListNode
		a, b, val, carry int
	)

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// l1 != nil && l2 != nil
	for pre = dummyHead; l1 != nil || l2 != nil; pre = pre.Next {
		// !note: 由于自己写的三目表达式，没有短路功能，所以nil节点，调取其val成员会panic
		// a = utils.If(l1 == nil, 0, l1.Val).(int)
		// b = utils.If(l2 == nil, 0, l2.Val).(int)
		if l1 == nil {
			a = 0
		} else {
			a = l1.Val.(int)
		}
		if l2 == nil {
			b = 0
		} else {
			b = l2.Val.(int)
		}

		val = (a + b + carry) % 10
		carry = (a + b + carry) / 10
		pre.Next = utils.NewListNode(val)
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		pre.Next = utils.NewListNode(carry)
	}
	return dummyHead.Next
}
