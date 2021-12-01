/*
Given a list, rotate the list to the right by k places, where k is non-negative.
For example: Given 1->2->3->4->5->nullptr and k = 2, return 4->5->1->2->3->nullptr.

* @Author: Yajun
* @Date:   2021/11/30 20:40
*/

package chap2

import (
	"solution/utils"
)

// 数组的旋转，可以通过多个reverse做到
// [1,2,3,4,5]
// [3,2,1,5,4]
// [4,5,1,2,3]

// 遍历两次，第一次遍历求出len，然后首尾连接形成环，接着往后跑len-k步，然后断开即可
func rotateRight(head *utils.ListNode, k int) *utils.ListNode {
	if head == nil {
		return head
	}
	var length = 1
	var cur *utils.ListNode

	for cur = head; cur.Next != nil; cur = cur.Next {
		length++
	}
	// cur.next == nil
	k = k % length
	if k == 0 {
		return head
	}
	cur.Next = head

	for i := 0; i < length-k; i++ {
		cur = cur.Next
	}
	head = cur.Next
	cur.Next = nil
	return head
}
