/*
Given a linked list, determine if it has a cycle in it.
Follow up: Can you solve it without using extra space?

* @Author: Yajun
* @Date:   2021/12/12 11:47
*/

package chap2

import "solution/utils"

// 快慢指针，fast指针进入环后会一直在环内跑圈，等slow指针进入环后，fast指针会套圈追上slow指针
func hasCycle(head *utils.ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
