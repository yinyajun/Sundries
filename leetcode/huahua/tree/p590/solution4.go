/*
* @Author: Yajun
* @Date:   2022/10/1 16:48
 */

package p590

func postorder4(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}

	var stack []*Node
	var visited = make(map[*Node]bool)
	stack = append(stack, root)

	for len(stack) > 0 {
		root = stack[len(stack)-1]

		if len(root.Children) == 0 || visited[root] {
			res = append(res, root.Val)
			stack = stack[:len(stack)-1]
			continue
		}

		for i := len(root.Children) - 1; i >= 0; i-- {
			stack = append(stack, root.Children[i])
		}
		visited[root] = true
	}

	return res
}
