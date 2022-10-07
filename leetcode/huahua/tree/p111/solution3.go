/*
* @Author: Yajun
* @Date:   2022/10/7 12:14
 */

package p111

// BFS方式，只要遇到第一个叶子节点
func minDepth4(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue []*TreeNode
	var res int
	queue = append(queue, root)

	for len(queue) > 0 {
		res++
		sz := len(queue)

		for _, root = range queue {
			if root.Left == nil && root.Right == nil { // leaf
				return res
			}

			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
		}

		queue = queue[sz:]
	}
	return res
}
