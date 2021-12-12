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

func reverseBetweenC(head *utils.ListNode, m, n int) *utils.ListNode {
	return reverseBetweenRecursive(head, m, n)
}

func reverseBetweenRecursive(head *utils.ListNode, m, n int) *utils.ListNode {
	if m == 1 { // 经过了m-1次
		// 反转链表前n个
		return reverseKList(head, n)
	}
	head.Next = reverseBetweenRecursive(head.Next, m-1, n-1)
	return head
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
		pre.Next = cur.Next // note: 这里不仅有缓存cur.next的作用，pre作为反转后的链表尾部，将其后继节点连接到当前节点之后（保证最后tail.next=nil）

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
	// 由于每次操作仅仅反向连接，没有处理尾节点
	dummy.Next.Next = nil // 这里直接写dummy = nil并不能改变head.next
	return pre            // cur is nil
}

// 迭代解法：C解法的更好写法
func reverseListD(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}
	pre, cur := head, head.Next

	for cur != nil {
		head.Next = cur.Next // 不仅缓存cur.next，而且处理了尾节点

		cur.Next = pre
		pre, cur = cur, head.Next
	}
	return pre // cur is nil
}

// 迭代解法：C解法的更好写法
func reverseListF(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}

	var pre, cur *utils.ListNode
	pre, cur = nil, head

	for cur != nil {
		post := cur.Next
		cur.Next = pre
		pre, cur = cur, post
	}
	return pre
}

// 和迭代解法一模一样
func reverseListE(head *utils.ListNode) *utils.ListNode {
	dummy := new(utils.ListNode)
	dummy.Next = head
	if head == nil {
		return head
	}
	reverseRecur(dummy, head, head.Next)
	return dummy.Next
}

func reverseRecur(dummy, head, cur *utils.ListNode) {
	if cur == nil {
		return
	}
	head.Next = cur.Next

	cur.Next = dummy.Next
	dummy.Next = cur

	reverseRecur(dummy, head, head.Next)
}

/*
反转整个链表中的前k个
*/

// if k <=1, 直接返回head，不需要反转
// else, 需要反转，反转以head为首的链表的前k个，返回新的链表头，链表尾部链接后续不需要反转的链表
func reverseKList(head *utils.ListNode, k int) *utils.ListNode {
	if k <= 1 || head == nil || head.Next == nil {
		return head
	}
	newHead := reverseKList(head.Next, k-1) // ensure head != nil
	follow := head.Next.Next                // ensure head.Next != nil
	head.Next.Next = head
	head.Next = follow
	// 也可以这样简略，但是不好理解
	//head.Next.Next, head.Next = head, head.Next.Next
	return newHead
}

func reverseKListC(head *utils.ListNode, k int) *utils.ListNode {
	dummy := new(utils.ListNode)
	dummy.Next = head
	if head == nil {
		return head
	}

	reverseKRecur(dummy, head, head.Next, k)
	return dummy.Next
}

func reverseKRecur(dummy, head, cur *utils.ListNode, k int) {
	if cur == nil || k <= 1 {
		return
	}

	head.Next = cur.Next

	cur.Next = dummy.Next
	dummy.Next = cur
	reverseKRecur(dummy, head, head.Next, k-1)
}

// 使用头插法
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
