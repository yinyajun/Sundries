/*
* @Author: Yajun
* @Date:   2022/10/1 16:56
 */

package p145

// 和前序几乎一样，仅仅是多维护了一个visited，判断root节点的所有孩子是否已经压入栈中
func postorderTraversal6(root *TreeNode) []int {
	var res []int

	if root == nil{return res}

	var stack []*TreeNode
	var visited = make(map[*TreeNode]bool)

	stack = append(stack, root)

	for len(stack)> 0{
		root = stack[len(stack)-1]

		if visited[root]{
			res = append(res, root.Val)
			stack = stack[:len(stack)-1]
			continue
		}

		if root.Right!=nil{
			stack = append(stack, root.Right)
		}
		if root.Left !=nil{
			stack = append(stack, root.Left)
		}
		visited[root] = true
	}

	return res
}