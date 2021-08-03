package main

import (
	"CodeGuide/base/abstract"
	"CodeGuide/base/fundamentals"
	"CodeGuide/base/utils"
)

// 判断一个二叉树是否是BST，改写中序遍历，在遍历的过程中，查看节点值是否递增即可

// 使用morris中序遍历（morris的思路是，在向左的过程中建立线索，然后回溯的时候取消线索）
func isBSTNode(root *abstract.TreeNode) bool {
	if root == nil {
		return true
	}

	isValid := func(prev, cur *abstract.TreeNode) bool {
		if prev != nil && utils.Less(cur.Key, prev.Key) {
			return false
		}
		return true
	}

	res := true
	var prev, cur *abstract.TreeNode
	cur = root
	for cur != nil {
		if cur.Left != nil { // 在左子树中寻找前续节点(左子树最右边的节点)
			pre := cur.Left
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}
			if pre.Right == nil { // 建立线索
				pre.Right = cur
				cur = cur.Left
			} else { // pre.right == cur， 取消线索
				pre.Right = nil
				res = isValid(prev, cur)
				prev = cur
				cur = cur.Right
			}
		} else { // cur.left == nil
			res = isValid(prev, cur)
			prev = cur
			cur = cur.Right
		}
	}
	return res
}

// 有右孩子没左孩子，非法；遍历到一定节点后，后续节点应该全是叶子节点
// 左0右1：invalid
// 左0右0：leaf
// 左1右0：leaf
// 左1右1：not leaf
func IsCBT(root *abstract.TreeNode) bool {
	if root == nil {
		return true
	}

	queue := fundamentals.NewLinkedQueue()
	queue.Enqueue(root)
	shouldBeLeaf := false
	for !queue.IsEmpty() {
		node := queue.Dequeue().(*abstract.TreeNode)
		l, r := node.Left, node.Right
		if (shouldBeLeaf && (l != nil || r != nil)) || (r != nil && l == nil) {
			return false
		}
		if l != nil {
			queue.Enqueue(l)
		}
		if r != nil {
			queue.Enqueue(r)
		} else { // r == nil
			shouldBeLeaf = true
		}
	}
	return true
}
