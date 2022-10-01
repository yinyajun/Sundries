/*
* @Author: Yajun
* @Date:   2022/10/1 16:02
 */

package p590


type Node struct {
	Val      int
	Children []*Node
}


// 递归解法（mutable）
func postorder(root *Node) []int {
	var res []int

	var recurse func(n *Node)

	recurse= func(n *Node) {
		if n == nil{
			return
		}

		for _, child := range n.Children{
			recurse(child)
		}
		res = append(res, n.Val)
	}

	recurse(root)
	return res
}


// 递归解法（immutable）
func postorder1(root *Node) []int {
	var res []int
	if root ==nil{
		return res
	}

	for _, child := range root.Children{
		res = append(res, postorder1(child)...)
	}
	res = append(res, root.Val)

	return res
}