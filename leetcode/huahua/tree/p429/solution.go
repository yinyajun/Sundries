/*
* @Author: Yajun
* @Date:   2022/9/30 11:43
 */

package p429

type Node struct {
	Val      int
	Children []*Node
}

//输入：root = [1,null,3,2,4,null,5,6]
//输出：[[1],[3,2,4],[5,6]]

// 层序遍历
// time: O(n) space:O(h)
func levelOrder(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var queue []*Node
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue) // 当前层的节点个数
		var vals = make([]int, size)
		for i := 0; i < size; i++ {
			vals[i] = queue[i].Val
			for k := range queue[i].Children { // 将其孩子节点入队
				queue = append(queue, queue[i].Children[k])
			}
		}
		res = append(res, vals)
		queue = queue[size:] // 出队当前层的
	}

	return res
}

func levelOrder2(root *Node) [][]int {
	var res [][]int
	var preorder func(node *Node, level int)

	preorder = func(n *Node, level int) {
		if n == nil {
			return
		}
		if level >= len(res) {
			res = append(res, []int{})
		}

		res[level] = append(res[level], n.Val)
		for _, child := range n.Children {
			preorder(child, level+1)
		}
	}

	preorder(root, 0)

	return res
}
