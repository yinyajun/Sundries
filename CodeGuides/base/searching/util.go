/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2021/1/8 20:47
 */

package searching

import (
	"CodeGuide/base/abstract"
	"CodeGuide/base/fundamentals"
	"container/list"
	"fmt"
)

// preorder建树
func CreateTreeFromQueue(queue abstract.Queue) *abstract.TreeNode {
	if queue.IsEmpty() {
		return nil
	}
	val := queue.Dequeue().(string)
	if val == "#" {
		return nil
	}
	node := abstract.NewTreeNode(val, nil)
	node.Left = CreateTreeFromQueue(queue)
	node.Right = CreateTreeFromQueue(queue)
	return node
}

func CreateTreeFromArray(array []string) *abstract.TreeNode {
	queue := fundamentals.NewLinkedQueue()
	for i := range array {
		queue.Enqueue(array[i])
	}
	return CreateTreeFromQueue(queue)
}

func CreateTreeFromArray2(array []string) *abstract.TreeNode {
	if len(array) == 0 {
		return nil
	}
	var root, cur *abstract.TreeNode
	q := new(list.List)
	var i int

	getNode := func(str string) *abstract.TreeNode {
		if str == "#" {
			return nil
		}
		return abstract.NewTreeNode(str, nil)
	}

	root = getNode(array[i])
	i++
	if root != nil {
		q.PushBack(root)
	}

	for q.Len() > 0 {
		cur = q.Remove(q.Front()).(*abstract.TreeNode)
		if i < len(array) {
			cur.Left = getNode(array[i])
			i++
			if cur.Left != nil {
				q.PushBack(cur.Left)
			}
		}
		if i < len(array) {
			cur.Right = getNode(array[i])
			i++
			if cur.Right != nil {
				q.PushBack(cur.Right)
			}
		}
	}
	return root
}

func createTreeFromArray(idx *int, array []string) *abstract.TreeNode {
	if *idx >= len(array) {
		return nil
	}
	input := array[*idx]
	*idx = (*idx) + 1
	if input == "#" {
		return nil
	}
	node := abstract.NewTreeNode(input, nil)
	node.Left = createTreeFromArray(idx, array)
	node.Right = createTreeFromArray(idx, array)
	return node
}

func CreateTree() *abstract.TreeNode {
	var input string
	fmt.Scanln(&input)
	if input == "#" { // empty tree
		return nil
	}
	node := abstract.NewTreeNode(input, nil)
	node.Left = CreateTree()
	node.Right = CreateTree()
	return node
}

// 为什么要用二阶指针？因为形参是值传递
func CreateTree2(node **abstract.TreeNode) {
	var input string
	fmt.Scanln(&input)
	if input == "#" { // empty tree
		*node = nil
		return
	}
	*node = abstract.NewTreeNode(input, nil)
	CreateTree2(&((*node).Left))
	CreateTree2(&((*node).Right))
}

func CloneTree(root *abstract.TreeNode) *abstract.TreeNode {
	if root == nil {
		return root
	}
	node := abstract.NewTreeNode(root.Key, nil)
	node.Left = CloneTree(root.Left)
	node.Right = CloneTree(root.Right)
	return node
}
