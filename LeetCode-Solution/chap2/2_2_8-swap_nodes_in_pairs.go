/*
Given a linked list, swap every two adjacent nodes and return its head.
For example, Given 1->2->3->4, you should return the list as 2->1->4->3.
Your algorithm should use only constant space. You may not modify the values in the list, only nodes
itself can be changed.

* @Author: Yajun
* @Date:   2021/11/30 22:27
*/

package chap2

import "solution/utils"

// todo: again
func swapNodesInPairs(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}
	dummy := new(utils.ListNode)
	dummy.Next = head
	pre := dummy
	//pre每次连接的都是odd，如果存在even节点，将even节点插入
	odd, even := head, head.Next // ensure head != nil
	for even != nil {
		odd.Next = even.Next

		// 头插法
		even.Next = pre.Next
		pre.Next = even

		pre = odd
		odd = pre.Next
		even = nil
		if odd != nil {
			even = odd.Next
		}
	}
	return dummy.Next
}

func swapNodesInPairsB(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd, even := head, head.Next

	odd.Next = swapNodesInPairs(even.Next)
	even.Next = head
	return even
}

// 和迭代法几乎一样
func swapNodesInPairsC(head *utils.ListNode) *utils.ListNode {
	dummy := new(utils.ListNode)
	dummy.Next = head
	snRecur(dummy, head)
	return dummy.Next
}

func snRecur(pre, head *utils.ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	odd, even := head, head.Next

	odd.Next = even.Next

	even.Next = odd
	pre.Next = even

	snRecur(odd, odd.Next)
}
