package main

import (
	"CodeGuide/base/abstract"
	"CodeGuide/base/searching"
	"fmt"
)

// 找到两个节点的最近公共祖先

// 先要找到两个节点，然后寻找其最近公共祖先
// 当前节点，cur == nil || cur == o1 || cur == o2 -> return cur
// 后序遍历得到左右子树
// 左右子树为空，return nil
// 左右子树都不为空，o1和o2分别在两个不同的子树上，cur是其最近的祖先，return cur
// 左右子树有一个不为空，根据之前的return值，可以知道非空子树要么是O1、O2中的一个，要么就是公共祖先 -> return 非空

func LowestAncestor(head, o1, o2 *abstract.TreeNode) *abstract.TreeNode {
	if head == nil || head == o1 || head == o2 {
		return head
	}

	left := LowestAncestor(head.Left, o1, o2)
	right := LowestAncestor(head.Right, o1, o2)

	// 左右子树都不为空
	if left != nil && right != nil {
		return head
	}

	// 有一个不为空或者两个都为空
	if left != nil {
		return left
	} else {
		return right
	}
}

// 想要通过预处理，降低每次query的时间

// 第一种方式，（类似于quickunion）建立每个节点的父节点
// 这样每个节点通过父链接，可以形成父节点路径
// 然后转为两个父链接路径中有没有公共点
type Record1 struct {
	m map[*abstract.TreeNode]*abstract.TreeNode
}

func NewRecord1(root *abstract.TreeNode) *Record1 {
	r := &Record1{make(map[*abstract.TreeNode]*abstract.TreeNode)}
	if root != nil { // 根节点的父节点是nil
		r.m[root] = nil
	}
	r.setMap(root)
	return r
}

// 前序遍历，将root节点的子节点注册其父节点
// 时间复杂度是O(N)，空间复杂度是O(N)
func (r *Record1) setMap(root *abstract.TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		r.m[root.Left] = root
	}
	if root.Right != nil {
		r.m[root.Right] = root
	}
	r.setMap(root.Left)
	r.setMap(root.Right)
}

// 时间复杂度为O(h)
func (r *Record1) query(o1, o2 *abstract.TreeNode) *abstract.TreeNode {
	path := map[*abstract.TreeNode]struct{}{}
	// 形成o1的父节点路径path

	for parent, ok := r.m[o1]; ok; parent, ok = r.m[o1] {
		path[o1] = struct{}{}
		o1 = r.m[parent]
	}
	// 不断查看o2的父节点路径中的节点是否在path中
	for _, ok := path[o2]; !ok; _, ok = path[o2] {
		o2 = r.m[o2]
	}
	return o2
}

type tuple struct {
	_1 *abstract.TreeNode
	_2 *abstract.TreeNode
}

// 直接建立树中任意两个节点之间的最近公共祖先的query
// 也不是直接穷举
// 对于每个节点h，h和所有后代节点的la都是h
// h左子树的每个节点和h右子树的每个节点的公共祖先都是h
type Record2 struct {
	m map[tuple]*abstract.TreeNode
}

func NewRecord2(root *abstract.TreeNode) *Record2 {
	r := &Record2{make(map[tuple]*abstract.TreeNode)}
	r.setMap(root)
	return r
}

// h和h左孩子，h和h右孩子，h左和h右
func (r *Record2) setMap(node *abstract.TreeNode) {
	if node == nil {
		return
	}
	r.headRecord(node.Left, node)
	r.headRecord(node.Right, node)
	r.subRecord(node)
	r.setMap(node.Left)
	r.setMap(node.Right)
}

// n作为h的子树，将n和h的pair设置公共祖先为h
func (r *Record2) headRecord(n, h *abstract.TreeNode) {
	if n == nil { // h != nil
		return
	}

	r.m[tuple{n, h}] = h
	r.headRecord(n.Left, h)
	r.headRecord(n.Right, h)
}

func (r *Record2) subRecord(h *abstract.TreeNode) {
	if h == nil {
		return
	}
	r.preLeft(h.Left, h.Right, h) // 左子树中的节点和右子树的节点组成的对
	//r.subRecord(h.Left)
	//r.subRecord(h.Right)
}

// 移动左子树节点，每次调用preRight
func (r *Record2) preLeft(ll, rr, h *abstract.TreeNode) {
	if ll == nil {
		return
	}
	r.preRight(ll, rr, h) // 允许右子树为空
	r.preLeft(ll.Left, rr, h)
	r.preLeft(ll.Right, rr, h)
}

//确定左子树中节点，移动右子树节点(保证左子树非空)
func (r *Record2) preRight(ll, rr, h *abstract.TreeNode) {
	if rr == nil {
		return
	}
	r.m[tuple{ll, rr}] = h
	r.preRight(ll, rr.Left, h)
	r.preRight(ll, rr.Right, h)
}

func (r *Record2) query(o1, o2 *abstract.TreeNode) *abstract.TreeNode {
	t1 := tuple{o1, o2}
	if node, ok := r.m[t1]; ok {
		return node
	}

	t2 := tuple{o2, o1}
	if node, ok := r.m[t2]; ok {
		return node
	}
	return nil
}

func main() {
	root := searching.CreateTreeFromArray([]string{"1", "2", "4", "#", "#", "5", "#", "#", "3", "6", "#", "#", "7", "8", "#", "#", "#"})
	r1 := NewRecord1(root)
	r2 := NewRecord2(root)
	o1 := root.Left.Right
	o2 := root.Right.Right.Left
	fmt.Println(o1.Key, o2.Key, r1.query(o1, o2).Key)
	fmt.Println(o1.Key, o2.Key, r2.query(o1, o2).Key)
}
