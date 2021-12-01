/*
Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers
from the original list.
For example,
Given 1->2->3->3->4->4->5, return 1->2->5.
Given 1->1->1->2->3, return 2->3.


* @Author: Yajun
* @Date:   2021/11/30 18:05
*/

package chap2

import (
	"solution/utils"
)

func deleteDuplicates2(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	post := head.Next
	if post.Val == head.Val {
		for post != nil && post.Val == head.Val {
			post = post.Next
		}
		// post == nil || post.val != head.val
		return deleteDuplicates2(post)
	} else {
		head.Next = deleteDuplicates2(head.Next)
		return head
	}
}

// todo: again
func deleteDuplicates2B(head *utils.ListNode) *utils.ListNode {
	dummy := new(utils.ListNode)
	dummy.Next = head
	pre, cur := dummy, head
	var duplicated bool

	for cur != nil {
		duplicated = false
		for cur.Next != nil && cur.Val == cur.Next.Val {
			duplicated = true
			cur = cur.Next
		}
		// cur.next == nil || cur.val != cur.next.val
		if duplicated {
			cur = cur.Next
			continue
		}
		pre.Next = cur
		pre, cur = pre.Next, cur.Next
	}
	pre.Next = nil // 只有遇到不重复的才重新连接，最后一次pre没有链接到nil
	return dummy.Next
}
