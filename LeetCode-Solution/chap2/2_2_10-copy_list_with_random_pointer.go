/*
A linked list is given such that each node contains an additional random pointer which could point to
any node in the list or null.
Return a deep copy of the list.

* @Author: Yajun
* @Date:   2021/12/1 17:03
*/

package chap2

type RandomNode struct {
	val          interface{}
	next, random *RandomNode
}

func copyRandomList(head *RandomNode) *RandomNode {
	for cur := head; cur != nil; {
		node := &RandomNode{val: cur.val}
		node.next = cur.next
		cur.next = node
		cur = node.next
	}

	for cur := head; cur != nil; {
		node := cur.next

		if cur.random != nil {
			node.random = cur.random.next
		}
		cur = node.next
	}

	dummy := new(RandomNode)
	pre := dummy

	for cur := head; cur != nil; {
		node := cur.next
		pre.next = node
		pre = node

		cur.next = node.next
		cur = cur.next
	}
	return dummy.next
}
