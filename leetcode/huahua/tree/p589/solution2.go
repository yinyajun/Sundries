/*
* @Author: Yajun
* @Date:   2022/10/1 15:08
 */

package p589

// 迭代解法
func preorder2(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}

	var stack []*Node
	stack = append(stack, root)

	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, root.Val)
		for i := len(root.Children) - 1; i >= 0; i-- {
			stack = append(stack, root.Children[i])
		}
	}
	return res
}

func preorder3(root *Node) []int {
	var res []int

	if root == nil {
		return res
	}

	var stack []*Node
	var next = make(map[*Node]int)

	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			if len(root.Children) >= 1 {
				next[root] = 1
				root = root.Children[0]
			} else {
				root = nil
			}
		}
		// root == nil
		if len(stack) > 0 {
			root = stack[len(stack)-1]  // 需要索引完所有子节点，才能出栈
			i := next[root]             // 已经遍历到root的第几个孩子啦
			if i < len(root.Children) { // 子节点都遍历过了吗
				next[root] = i + 1
				root = root.Children[i]
			} else {
				stack = stack[:len(stack)-1] // root的孩子都遍历完了，root可以出栈了
				delete(next, root)
				root = nil // 已经无法索引到下一个节点了，root置为nil，从stack中出栈下一个节点
			}
		}
	}
	return res
}
