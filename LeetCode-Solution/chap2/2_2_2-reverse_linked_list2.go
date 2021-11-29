/*
Reverse a linked list from position m to n. Do it in-place and in one-pass.
For example: Given 1->2->3->4->5->nullptr, m = 2 and n = 4,
return 1->4->3->2->5->nullptr.
Note: Given m, n satisfy the following condition: 1 ≤ m ≤ n ≤ length of list.

* @Author: Yajun
* @Date:   2021/11/28 09:51
*/

package chap2

import (
	"solution/utils"
)

// time: O(n); space： O(1)
func reverseBetween(head *utils.ListNode, m, n int) *utils.ListNode {
	// suppose m, n is valid
	var (
		idx                                  int
		startPre, startNext, endPre, endNext *utils.ListNode
		next                                 *utils.ListNode
	)

	dummy := new(utils.ListNode)
	dummy.Next = head
	pre, cur := dummy, head

	for idx = 1; cur != nil; idx++ {
		if idx <= m {
			if idx == m {
				startPre, endPre = pre, cur
			}
			pre, cur = pre.Next, cur.Next
		} else if idx <= n {
			next = cur.Next
			if idx == n {
				startNext, endNext = cur, next
			}
			cur.Next = pre // reverse
			pre, cur = cur, next
		} else {
			pre, cur = pre.Next, cur.Next
		}
	}
	if startPre != nil {
		startPre.Next = startNext
	}
	if endPre != nil {
		endPre.Next = endNext
	}
	return dummy.Next
}

// 利用头插法
func reverseBetweenB(head *utils.ListNode, m, n int) *utils.ListNode {
	dummy := new(utils.ListNode)
	dummy.Next = head
	pre := dummy

	for i := 1; i < m; i++ {
		pre = pre.Next
	}

	head2 := pre // pre is m-1(相当于反转的这段链表的dummyHead)

	pre = head2.Next
	cur := pre.Next
	for i := m; i < n; i++ {
		pre.Next = cur.Next

		// 头插法，将当前cur插入到head2中
		cur.Next = head2.Next
		head2.Next = cur

		cur = pre.Next
	}
	return dummy.Next
}

/*
反转整个链表
*/

// time: O(n); space: O(n)
// 递归解法
func reverseList(head *utils.ListNode) *utils.ListNode {
	return reverseRecursive(head)
}

func reverseRecursive(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil // 否则整个链表将会没有终止
	return newHead
}

// 迭代解法：利用头插法，每次都将当前元素作为反转链表的表头元素
func reverseListB(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}

	dummy := new(utils.ListNode)
	dummy.Next = head
	pre, cur := head, head.Next // ensure head != nil

	for cur != nil {
		pre.Next = cur.Next

		// 头插法，将cur插入到首个位置
		cur.Next = dummy.Next
		dummy.Next = cur

		cur = pre.Next
	}
	return dummy.Next
}

// 迭代解法：当前遍历元素链接上previous元素，最后返回尾节点
func reverseListC(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}
	dummy := new(utils.ListNode)
	dummy.Next = head

	pre, cur := dummy, head
	var post *utils.ListNode

	for cur != nil {
		post = cur.Next
		cur.Next = pre
		pre, cur = cur, post
	}
	// dummy.next -> head, ensure head != nil
	dummy.Next.Next = nil // 这里直接写dummy = nil并不能改变head.next
	return pre            // cur is nil
}

/*
反转整个链表中的前k个
*/
func reverseKList(head *utils.ListNode, k int) *utils.ListNode {
	return reverseKRecursive(head, k)
}

// if k <=1, 直接返回head，不需要反转
// else, 需要反转，反转以head为首的链表的前k个，返回新的链表头，链表尾部链接后续不需要反转的链表
func reverseKRecursive(head *utils.ListNode, k int) *utils.ListNode {
	if k <= 1 { // todo: check head
		return head
	}
	newHead := reverseKRecursive(head.Next, k-1)
	follow := head.Next.Next
	head.Next.Next = head
	head.Next = follow
	// 也可以这样简略，但是不好理解
	//head.Next.Next, head.Next = head, head.Next.Next
	return newHead
}

func reverseKListB(head *utils.ListNode, k int) *utils.ListNode {
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
