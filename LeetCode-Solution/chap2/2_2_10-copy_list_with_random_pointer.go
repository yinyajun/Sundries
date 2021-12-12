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
	// 在现有节点后面连上一个新节点（头插法）
	for cur := head; cur != nil; {
		node := &RandomNode{val: cur.val}
		node.next = cur.next
		cur.next = node
		cur = node.next
	}
	// 将新节点的random指针指向正确位置
	for cur := head; cur != nil; {
		node := cur.next
		if cur.random != nil {
			node.random = cur.random.next
		}
		cur = node.next
	}

	// 分离出新链表
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
