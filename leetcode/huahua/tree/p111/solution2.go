/*
* @Author: Yajun
* @Date:   2022/10/7 11:47
 */

package p111

type pair struct {
	node  *TreeNode
	depth int
}

// 迭代解法：用stack模拟自顶向下递归
func minDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var stack []pair
	var res = 2 << 31
	stack = append(stack, pair{root, 1})

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if p.node.Left == nil && p.node.Right == nil { // leaf
			res = min(res, p.depth)
		}

		if p.node.Left != nil {
			stack = append(stack, pair{p.node.Left, p.depth + 1})
		}

		if p.node.Right != nil {
			stack = append(stack, pair{p.node.Right, p.depth + 1})
		}
	}
	return res
}
