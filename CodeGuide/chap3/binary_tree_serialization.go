package main

import (
	"CodeGuide/base/abstract"
	"CodeGuide/base/fundamentals"
	"CodeGuide/base/searching"
	"fmt"
	"strings"
)

func PreSerialization(root *abstract.TreeNode) string {
	if root == nil {
		return "#!"
	}
	res := strings.Builder{}
	res.WriteString(root.Key.(string) + "!")
	res.WriteString(PreSerialization(root.Left))
	res.WriteString(PreSerialization(root.Right))
	return res.String()
}

func PreDeserialization(str string) *abstract.TreeNode {
	queue := fundamentals.NewLinkedQueue()
	strs := strings.Split(str, "!")
	for i := range strs {
		queue.Enqueue(strs[i])
	}
	return deserialization(queue)
}

func deserialization(queue abstract.Queue) *abstract.TreeNode {
	if queue.IsEmpty() {
		return nil
	}
	val := queue.Dequeue().(string)
	if val == "#" {
		return nil
	}
	node := abstract.NewTreeNode(val, nil)
	node.Left = deserialization(queue)
	node.Right = deserialization(queue)
	return node
}

// 入队前访问，这样能访问到空节点
// 出队时访问，智能访问到非空节点
func levelSerialization(root *abstract.TreeNode) string {
	if root == nil {
		return "#!"
	}
	queue := fundamentals.NewLinkedQueue()
	res := strings.Builder{}
	res.WriteString(root.Key.(string) + "!")
	queue.Enqueue(root)
	for !queue.IsEmpty() {
		node := queue.Dequeue().(*abstract.TreeNode)
		if node.Left != nil {
			res.WriteString(node.Left.Key.(string) + "!")
			queue.Enqueue(node.Left)
		} else {
			res.WriteString("#!")
		}
		if node.Right != nil {
			res.WriteString(node.Right.Key.(string) + "!")
			queue.Enqueue(node.Right)
		} else {
			res.WriteString("#!")
		}
	}
	return res.String()
}

func levelDeserialization(str string) *abstract.TreeNode {
	values := strings.Split(str, "!")
	valuesQueue := fundamentals.NewLinkedQueue()
	for i := range values {
		valuesQueue.Enqueue(values[i])
	}
	queue := fundamentals.NewLinkedQueue()
	root := generateNodeByQueue(valuesQueue)
	if root == nil {
		return nil
	}
	queue.Enqueue(root)
	for !queue.IsEmpty() {
		node := queue.Dequeue().(*abstract.TreeNode)
		node.Left = generateNodeByQueue(valuesQueue)
		node.Right = generateNodeByQueue(valuesQueue)
		if node.Left != nil {
			queue.Enqueue(node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
	}
	return root
}

func generateNodeByQueue(queue abstract.Queue) *abstract.TreeNode {
	if queue.IsEmpty() {
		return nil
	}
	val := queue.Dequeue().(string)
	if val == "#" {
		return nil
	}
	return abstract.NewTreeNode(val, nil)
}

func main() {
	root := searching.CreateTreeFromArray([]string{"12", "3", "#", "#", "#"})
	res := PreSerialization(root)
	fmt.Println(res)
	root = PreDeserialization(res)
	fmt.Println(PreSerialization(root))
	fmt.Println(levelSerialization(root))
	root = levelDeserialization(levelSerialization(root))
	fmt.Println(levelSerialization(root))
}
