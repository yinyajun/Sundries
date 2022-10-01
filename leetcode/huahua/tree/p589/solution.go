/*
* @Author: Yajun
* @Date:   2022/10/1 15:03
 */

package p589

type Node struct {
	Val      int
	Children []*Node
}

// 递归解法（mutable）
func preorder(root *Node) []int {
	var res []int

	var recursive func(node *Node)

	recursive = func(node *Node) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		for _, child := range node.Children {
			recursive(child)
		}
	}

	recursive(root)

	return res
}


// 递归解法（immutable）
func preorder1(root *Node) []int{
	var res []int
	if root == nil{
		return res
	}

	res= append(res, root.Val)
	for _, child := range root.Children{
		res = append(res, preorder1(child)...)
	}
	return res
}