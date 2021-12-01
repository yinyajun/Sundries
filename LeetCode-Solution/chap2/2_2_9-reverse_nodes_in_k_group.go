/*
Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.
You may not alter the values in the nodes, only nodes itself may be changed.
Only constant memory is allowed.
For example, Given this linked list: 1->2->3->4->5
For k = 2, you should return: 2->1->4->3->5
For k = 3, you should return: 3->2->1->4->5


* @Author: Yajun
* @Date:   2021/12/1 11:09
*/

package chap2

import (
	"solution/utils"
)

func reverseNodesInGroup(head *utils.ListNode, k int) *utils.ListNode {
	dummy := new(utils.ListNode)
	dummy.Next = head
	pre := dummy

	hasK := func(node *utils.ListNode) bool {
		for i := 0; i < k; i++ {
			if node == nil {
				return false
			}
			node = node.Next
		}
		return true
	}

	reverseKList := func(head *utils.ListNode, k int) *utils.ListNode {
		if head == nil {
			return head
		}
		dummy := new(utils.ListNode)
		dummy.Next = head

		idx := 1
		pre, cur := head, head.Next // head != nil

		for cur != nil && idx < k {
			pre.Next = cur.Next

			cur.Next = dummy.Next
			dummy.Next = cur

			cur = pre.Next
			idx++
		}
		return dummy.Next
	}

	for hasK(pre.Next) {
		start := reverseKList(pre.Next, k)
		end := pre.Next

		pre.Next = start
		pre = end // pre.next is end
	}

	return dummy.Next
}

func reverseNodesInGroupB(head *utils.ListNode, k int) *utils.ListNode {
	if head == nil {
		return head
	}
	group := head
	for i := 0; i < k; i++ {
		if group == nil {
			return head
		}
		group = group.Next
	}

	newGroupHead := reverseNodesInGroupB(group, k)

	var prev *utils.ListNode
	cur := head

	for cur != group {
		next := cur.Next

		cur.Next = utils.If(prev == nil, newGroupHead, prev).(*utils.ListNode)
		prev, cur = cur, next
	}
	return prev
}
