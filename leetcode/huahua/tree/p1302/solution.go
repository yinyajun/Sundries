/*
* @Author: Yajun
* @Date:   2022/10/3 14:58
 */

package p1302

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层序遍历
func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var res int
	var q []*TreeNode
	q = append(q, root)

	for len(q) > 0 {
		sz := len(q) // 当前层的节点个数
		res = 0
		for _, root = range q {
			res += root.Val

			if root.Left != nil {
				q = append(q, root.Left)
			}
			if root.Right != nil {
				q = append(q, root.Right)
			}
		}
		q = q[sz:] // 出队当前层
	}
	return res
}

func deepestLeavesSum2(root *TreeNode) int {
	var res int
	if root == nil {
		return res
	}

	var maxDepth int
	var recurse func(root *TreeNode, depth int)

	recurse = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level > maxDepth { // 如果遇到更深的叶子节点
			maxDepth = level
			res = 0
		}

		if root.Left == nil && root.Right == nil { // Leaf
			if level == maxDepth { // 只计算最深一层的叶子节点的和
				res += root.Val
			}
			return
		}
		recurse(root.Left, level+1)
		recurse(root.Right, level+1)
	}

	recurse(root, 0)
	return res
}
