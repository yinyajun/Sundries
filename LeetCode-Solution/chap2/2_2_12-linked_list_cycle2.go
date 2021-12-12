/*
Given a linked list, return the node where the cycle begins. If there is no cycle, return null.
Follow up: Can you solve it without using extra space?

* @Author: Yajun
* @Date:   2021/12/12 12:11
*/

package chap2

import "solution/utils"

// 设环长为r，链表长度为L，环入口点和相遇点距离为a，起点到环入口点为x
// 设slow走了s步，fast走了2s步
// 2s = s + nr
// s = nr
// x + a = s = nr = (n-1)r + (L-x)
// x = (n-1)r + (L-x-a)
// (x+a)为相遇点，L-x-a为相遇点到环入口的距离
// 也就是说，x = 相遇点到环入口的距离 + (n-1)环
// 再来个slow2，从head出发。slow和slow2同步前进，再次相遇的地方就是环入口

func detectCycle(head *utils.ListNode) *utils.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow2 := head
			for slow2 != slow {
				slow = slow.Next
				slow2 = slow2.Next
			}
			return slow2
		}
	}
	return nil
}
