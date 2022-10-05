/*
* @Author: Yajun
* @Date:   2022/10/5 17:45
 */

package p111

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	var res = 2 << 31

	var recurse func(node *TreeNode, level int)

	recurse = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			if level < res {
				res = level
			}
		}
		recurse(node.Left, level+1)
		recurse(node.Right, level+1)
	}

	recurse(root, 1)

	if res == 2<<31 {
		return 0
	}
	return res
}
