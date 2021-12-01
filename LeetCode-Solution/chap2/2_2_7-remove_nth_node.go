/*
Given a linked list, remove the nth node from the end of list and return its head.
For example, Given linked list: 1->2->3->4->5, and n = 2.
After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:
• Given n will always be valid.
• Try to do this in one pass.

* @Author: Yajun
* @Date:   2021/11/30 21:04
*/

package chap2

import "solution/utils"

// 快慢指针，寻找到倒数第n+1个节点
func removeNthNode(head *utils.ListNode, n int) *utils.ListNode {
	dummy := new(utils.ListNode)
	dummy.Next = head
	p, q := dummy, dummy

	for i := 0; i < n; i++ {
		q = q.Next
	}

	for q.Next != nil {
		p = p.Next
		q = q.Next
	}
	// p is n+1 th node
	p.Next = p.Next.Next
	return dummy.Next
}
