/*
Given a sorted linked list, delete all duplicates such that each element appear only once.
For example,
Given 1->1->2, return 1->2.
Given 1->1->2->3->3, return 1->2->3.

* @Author: Yajun
* @Date:   2021/11/30 14:57
*/

package chap2

import (
	"solution/utils"
)

// 不同的时候重新连接，最后需要将尾部置nil
func deleteDuplicate(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}
	var pre, cur *utils.ListNode
	for pre, cur = head, head.Next; cur != nil; cur = cur.Next {
		if cur.Val == pre.Val {
			continue
		}
		pre.Next = cur
		pre = pre.Next
	}
	pre.Next = nil
	return head
}

// 相同的时候重新连接，连接的时候直接跳过当前的相同节点
func deleteDuplicateB(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}

	for pre, cur := head, head.Next; cur != nil; cur = pre.Next {
		if cur.Val == pre.Val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
	}
	return head
}

func deleteDuplicateC(head *utils.ListNode) *utils.ListNode {
	return ddRecurC(head)
}

func ddRecurC(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	s := ddRecurC(head.Next)
	if head.Val == s.Val {
		head.Next = s.Next
	}
	return head
}

func deleteDuplicateD(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ddRecurD(head, head.Next)
	return head
}

func ddRecurD(pre, cur *utils.ListNode) {
	if cur == nil {
		return
	}
	if pre.Val == cur.Val {
		pre.Next = cur.Next // delete cur
	} else {
		pre = cur
	}
	ddRecurD(pre, cur.Next)
}
