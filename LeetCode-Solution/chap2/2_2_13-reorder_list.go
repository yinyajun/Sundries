/*
Given a singly linked list L : L0 → L1 → · · · → Ln 1 → Ln, reorder it to: L0 → Ln → L1 → Ln 1 → L2 → Ln 2 → · · ·
You must do this in-place without altering the nodes’ values.
For example, Given {1,2,3,4}, reorder it to {1,4,2,3}.

* @Author: Yajun
* @Date:   2021/12/12 13:11
*/

package chap2

import "solution/utils"

// 找到中间节点，断开，将后半段链表反转，然后合并
// 1 2 3 4 5 6 7
// 1 7 2 6 3 5 4 （1 2 3 4）（7 6 5）
func reorderList(head *utils.ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head
	var prev *utils.ListNode // 前段的尾节点，需要将其next置为nil，来剪断两个链表

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil // cut at middle

	reverse := func(head *utils.ListNode) *utils.ListNode {
		var pre, cur, post *utils.ListNode
		cur = head

		for cur != nil {
			post = cur.Next
			cur.Next = pre
			pre, cur = cur, post
		}
		// cur == nil
		return pre
	}
	slow = reverse(slow)

	// merge
	cur := head

	for cur.Next != nil { // ensure cur!=nil
		tmp := cur.Next
		cur.Next = slow
		slow = slow.Next
		cur.Next.Next = tmp
		cur = tmp
	}
	// cur.next == nil
	cur.Next = slow

}
