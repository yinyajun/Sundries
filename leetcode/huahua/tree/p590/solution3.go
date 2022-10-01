/*
* @Author: Yajun
* @Date:   2022/10/1 16:30
 */

package p590

// 求逆后序，然后reverse
func postorder3(root *Node) []int {
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

		for _, child := range root.Children {
			stack = append(stack, child)
		}
	}

	length := len(res)
	for i := 0; i < length/2; i++ {
		res[i], res[length-1-i] = res[length-1-i], res[i]
	}
	return res
}
