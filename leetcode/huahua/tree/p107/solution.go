/*
* @Author: Yajun
* @Date:   2022/10/8 20:37
 */

package p107

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		sz := len(queue)
		res = append(res, make([]int, 0, sz))
		for _, root := range queue {
			res[len(res)-1] = append(res[len(res)-1], root.Val)
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}

		}
		queue = queue[sz:]
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

func levelOrderBottom1(root *TreeNode) [][]int {
	var res [][]int
	var recurse func(node *TreeNode, level int)

	recurse = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if level >= len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], node.Val)

		recurse(node.Left, level+1)
		recurse(node.Right, level+1)
	}

	recurse(root, 0)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
