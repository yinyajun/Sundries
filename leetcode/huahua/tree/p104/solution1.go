/*
* @Author: Yajun
* @Date:   2022/10/5 12:21
 */

package p104

// 递归方法是DFS，这里用BFS的方式
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var depth int
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		sz := len(queue)
		for _, root := range queue {
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
		}
		queue = queue[sz:]
		depth++
	}
	return depth
}
