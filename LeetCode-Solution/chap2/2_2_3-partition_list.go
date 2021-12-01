/*
Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater
than or equal to x.
You should preserve the original relative order of the nodes in each of the two partitions.
For example, Given 1->4->3->2->5->2 and x = 3, return 1->2->2->4->3->5.

* @Author: Yajun
* @Date:   2021/11/30 14:40
*/

package chap2

import "solution/utils"

// time: O(n); space: O(1)
func partitionList(head *utils.ListNode, k int) *utils.ListNode {
	dummyA, dummyB := new(utils.ListNode), new(utils.ListNode)
	preA, preB := dummyA, dummyB

	for ; head != nil; head = head.Next {
		if head.Val.(int) < k {
			preA.Next = head
			preA = preA.Next
		} else {
			preB.Next = head
			preB = preB.Next
		}
	}
	preA.Next = dummyB.Next
	preB.Next = nil
	return dummyA.Next
}
