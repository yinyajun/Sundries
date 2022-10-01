/*
* @Author: Yajun
* @Date:   2022/10/1 16:06
 */

package p590

func postorder2(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}

	var stack []*Node
	var next = make(map[*Node]int) // 下一个需要遍历的是第几个孩子

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			if len(root.Children) == 0 {
				break
			}
			next[root] = 1
			root = root.Children[0]
		}
		// root == nil
		if len(stack) > 0 {
			root = stack[len(stack)-1]
			i := next[root]
			if i < len(root.Children) {
				next[root] = i + 1
				root = root.Children[i]
			} else { // 已经遍历完root的所有子节点
				res = append(res, root.Val)
				stack = stack[:len(stack)-1]
				delete(next, root)
				root = nil
			}
		}
	}

	return res
}
