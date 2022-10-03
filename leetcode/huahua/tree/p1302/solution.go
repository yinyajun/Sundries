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
