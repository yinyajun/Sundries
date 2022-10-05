/*
* @Author: Yajun
* @Date:   2022/10/5 12:11
 */

package p104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归解法，类似于后序遍历
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lDepth := maxDepth(root.Left)
	rDepth := maxDepth(root.Right)

	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}

// 递归解法，类似于前序遍历
func maxDepth1(root *TreeNode) int {
	var res = -1
	var recurse func(*TreeNode, int)

	recurse = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level > res {
			res = level
		}

		recurse(node.Left, level+1)
		recurse(node.Right, level+1)
	}

	recurse(root, 0)
	return res + 1
}
