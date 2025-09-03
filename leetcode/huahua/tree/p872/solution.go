/*
* @Author: Yajun
* @Date:   2022/10/8 21:12
 */

package p872

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var res []int
	var recurse func(node *TreeNode)

	recurse = func(n *TreeNode) {
		if n == nil {
			return
		}
		if n.Left == nil && n.Right == nil { // leaf
			res = append(res, n.Val)
			return
		}
		recurse(n.Left)
		recurse(n.Right)
	}

	recurse(root1)
	var res2 = make([]int, len(res))
	for i, n := range res {
		res2[i] = n
	}
	res = res[:0]
	recurse(root2)

	if len(res) != len(res2) {
		return false
	}
	for i, n := range res {
		if n != res2[i] {
			return false
		}
	}
	return true
}

func leafSimilar1(root1 *TreeNode, root2 *TreeNode) bool {
	var res1, res2 []int

	traverse := func(node *TreeNode, res *[]int) {
		if node == nil {
			return
		}

		var q []*TreeNode
		q = append(q, node)

		for len(q) > 0 {
			sz := len(q)
			for _, n := range q {
				if n.Left == nil && n.Right ==nil{
					*res = append(*res, n.Val)
				}
				if n.Left != nil {
					q = append(q, n.Left)
				}

				if n.Right != nil {
					q = append(q, n.Right)
				}
			}
			q = q[sz:]
		}
	}

	traverse(root1, &res1)
	traverse(root2, &res2)

	if len(res1) != len(res2) {
		return false
	}

	for i, n := range res1 {
		if res2[i] != n {
			return false
		}
	}
	return true
}
