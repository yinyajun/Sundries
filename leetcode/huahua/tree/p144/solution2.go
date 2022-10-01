/*
* @Author: Yajun
* @Date:   2022/9/29 20:01
 */

package p144


// 非递归实现
// 前序遍历，访问完节点后需要能索引到左右孩子，简单粗暴直接放入到stack中，注意顺序
// time: O(n) space:O(h)
func preorderTraversal2(root *TreeNode) []int {
	var res []int

	if root == nil{
		return res
	}

	var stack []*TreeNode

	stack = append(stack, root)

	// 栈中保存都是孩子节点（将需要遍历的节点提前压入栈中）
	for len(stack)>0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)

		// 按照右左孩子入栈，出栈时才是左右孩子
		if root.Right !=nil{
			stack = append(stack, root.Right)
		}

		if root.Left!=nil{
			stack = append(stack, root.Left)
		}
	}
	return res
}
