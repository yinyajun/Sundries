package main

import (
	"CodeGuide/base/abstract"
	"CodeGuide/base/fundamentals"
	"CodeGuide/base/searching"
	"CodeGuide/base/utils"
	"fmt"
)

func PrintBTBoundary1(root *abstract.TreeNode) {
	if root == nil {
		return
	}
	// 获取树的高度
	height := GetHeight1(root)
	// 初始化每层的边界节点map
	edgeMap := make([][]*abstract.TreeNode, height)
	for i := range edgeMap {
		edgeMap[i] = make([]*abstract.TreeNode, 2)
	}
	// 设置边界节点
	SetEdgeMap1(root, 0, edgeMap)
	// 打印左边界
	for i := range edgeMap {
		fmt.Println(edgeMap[i][0].Key)
	}
	// 打印非边界的叶子节点
	PrintLeafNotInMap(root, 0, edgeMap)
	// 打印右边界（排除既是左边界又是右边界的）
	for i := len(edgeMap) - 1; i >= 0; i-- {
		if edgeMap[i][0] != edgeMap[i][1] {
			fmt.Println(edgeMap[i][1].Key)
		}
	}
}

func GetHeight1(root *abstract.TreeNode) int {
	if root == nil {
		return 0
	}
	return utils.MaxInt(GetHeight1(root.Left), GetHeight1(root.Right)) + 1
}

func GetHeight2(root *abstract.TreeNode, l int) int {
	if root == nil {
		return l
	}
	return utils.MaxInt(GetHeight2(root.Left, l+1), GetHeight2(root.Right, l+1))
}

func SetEdgeMap1(root *abstract.TreeNode, l int, edgeMap [][]*abstract.TreeNode) {
	if root == nil {
		return
	}
	if edgeMap[l][0] == nil {
		edgeMap[l][0] = root
	}
	edgeMap[l][1] = root
	SetEdgeMap1(root.Left, l+1, edgeMap)
	SetEdgeMap1(root.Right, l+1, edgeMap)
}

func SetEdgeMap2(root *abstract.TreeNode, l int, edgeMap [][]*abstract.TreeNode) {
	if root == nil {
		return
	}
	if edgeMap[l][0] == nil {
		edgeMap[l][0] = root
	}
	SetEdgeMap2(root.Left, l+1, edgeMap)
	SetEdgeMap2(root.Right, l+1, edgeMap)
	edgeMap[l][1] = root
}

func PrintLeafNotInMap(root *abstract.TreeNode, l int, m [][]*abstract.TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil && root != m[l][0] && root != m[l][1] {
		fmt.Println(root.Key)
	}
	PrintLeafNotInMap(root.Left, l+1, m)
	PrintLeafNotInMap(root.Right, l+1, m)
}

func levelTraverse(root *abstract.TreeNode) {
	if root == nil {
		return
	}
	var level int
	queue := fundamentals.NewLinkedQueue()
	queue.Enqueue(root)
	for !queue.IsEmpty() {
		size := queue.Size()
		for i := 0; i < size; i++ {
			node := queue.Dequeue().(*abstract.TreeNode)
			fmt.Println(level, size, node.Key)
			if node.Left != nil {
				queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}
		level += 1
	}
}

func postTraverse(h *abstract.TreeNode) {
	if h == nil {
		return
	}
	stack := fundamentals.NewLinkedStack()
	stack.Push(h)
	var c *abstract.TreeNode
	for !stack.IsEmpty() {
		c = stack.Peek().(*abstract.TreeNode)
		if c.Left != nil && h != c.Left && h != c.Right { // 当前节点的左右子树均未遍历过，因为可能从左右子树节点回溯
			stack.Push(c.Left)
		} else if c.Right != nil && h != c.Right {
			stack.Push(c.Right)
		} else {
			fmt.Println(stack.Pop().(*abstract.TreeNode).Key)
			h = c
		}
	}
}

// 头节点、叶节点；树左边界延伸下去的路径；树右边界延伸下去的路径；
// 先找到第一个既有左孩子又有右孩子的节点
func PrintBTBoundary2(root *abstract.TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Key)
	if root.Left != nil && root.Right != nil {
		PrintLeftEdge(root.Left, true)
		PrintRightEdge(root.Right, true)
	} else {
		if root.Left == nil {
			PrintBTBoundary2(root.Right)
		} else {
			PrintBTBoundary2(root.Left)
		}
	}
}

func isLeaf(root *abstract.TreeNode) bool {
	if root == nil {
		panic("root is nil")
	}
	return root.Left == nil && root.Right == nil
}

// 前序遍历，从上到下遍历
func PrintLeftEdge(root *abstract.TreeNode, print bool) {
	if root == nil {
		return
	}
	if print || isLeaf(root) {
		fmt.Println(root.Key)
	}
	PrintLeftEdge(root.Left, print)
	PrintLeftEdge(root.Right, print && (root.Left == nil))
}

// 后序遍历，从下到上遍历
func PrintRightEdge(root *abstract.TreeNode, print bool) {
	if root == nil {
		return
	}
	PrintRightEdge(root.Left, print && (root.Right == nil))
	PrintRightEdge(root.Right, print)
	if print || isLeaf(root) {
		fmt.Println(root.Key)
	}
}

func main() {
	root := searching.CreateTreeFromArray([]string{
		"1", "2", "#", "4", "7", "#", "#", "8", "#", "11", "13", "#", "#", "14", "#", "#",
		"3", "5", "9", "12", "15", "#", "#", "16", "#", "#", "#", "10", "#", "#", "6", "#", "#"})
	PrintBTBoundary2(root)
}
