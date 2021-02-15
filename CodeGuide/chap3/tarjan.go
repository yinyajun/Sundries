package main

import (
	"CodeGuide/base/abstract"
	"CodeGuide/base/fundamentals"
	"CodeGuide/base/searching"
	"fmt"
)

// 二叉树的最近公共祖先的问题
// 对于大小为N的二叉树，给定M个query，求出每队query的最近公共祖先，要求时间复杂度为O(N+M)

// 思路就是二叉树的遍历，二叉树的递归遍历有5个阶段
// 根，左，根，右，根
// 这里使用<左，根，右，根>的方式
// 情形1：两个节点在公共祖先节点的两个子树中，而通过回溯到根节点，提供了进入另一个子树的入口，所以该根节点一定是公共祖先
// 情形2：两个节点是祖孙关系，那么通过回溯到根节点，即公共祖先节点
// 已经遍历过的树结构的根节点就是公共祖先节点，需要记录已经遍历过的树结构及其根节点

// 如果在遍历中，快速确认另一个节点是否出现过，需要使用map（类似于twoSum）
// 而要记录树结构及其对应的根节点。采用<集合，node>的map方式、
// 同时考虑到有大量find和union操作，使用并查集

// 二叉树上的并查集，原来的unionFind使用数组，是因为key用索引代替
type UnionFind struct {
	// 二叉树上的并查集要用map
	parent map[*abstract.TreeNode]*abstract.TreeNode
	rank   map[*abstract.TreeNode]int
	count  int
}

func NewUnionFind() *UnionFind {
	uf := &UnionFind{
		parent: make(map[*abstract.TreeNode]*abstract.TreeNode),
		rank:   make(map[*abstract.TreeNode]int),
	}
	return uf
}

func (uf *UnionFind) Init(root *abstract.TreeNode) {
	uf.count = uf.init(root)
}

func (uf *UnionFind) init(root *abstract.TreeNode) int {
	if root == nil {
		return 0
	}
	uf.parent[root] = root
	uf.rank[root] = 0
	return uf.init(root.Left) + uf.init(root.Right) + 1
}

func (uf *UnionFind) Find(node *abstract.TreeNode) *abstract.TreeNode {
	return uf.find(node)
}

func (uf *UnionFind) find(node *abstract.TreeNode) *abstract.TreeNode {
	if node != uf.parent[node] {
		uf.parent[node] = uf.find(uf.parent[node])
	}
	return uf.parent[node]
}

func (uf *UnionFind) Union(a, b *abstract.TreeNode) {
	if a == nil || b == nil {
		return
	}
	aFather := uf.Find(a)
	bFather := uf.Find(b)

	if aFather == bFather {
		return
	}

	aRank, bRank := uf.rank[aFather], uf.rank[bFather]
	if aRank < bRank {
		uf.parent[aFather] = bFather
	} else if aRank > bRank {
		uf.parent[bFather] = aFather
	} else {
		uf.parent[aFather] = bFather
		uf.rank[bFather]++
	}
	uf.count--
}

type Query struct {
	o1, o2 *abstract.TreeNode
}

type Tarjan struct {
	queries   map[*abstract.TreeNode]*fundamentals.LinkedQueue
	indexes   map[*abstract.TreeNode]*fundamentals.LinkedQueue
	ancestors map[*abstract.TreeNode]*abstract.TreeNode
	set       *UnionFind
}

func NewTarjan() *Tarjan {
	return &Tarjan{
		make(map[*abstract.TreeNode]*fundamentals.LinkedQueue),
		make(map[*abstract.TreeNode]*fundamentals.LinkedQueue),
		make(map[*abstract.TreeNode]*abstract.TreeNode),
		NewUnionFind(),
	}
}

func (t *Tarjan) Query(head *abstract.TreeNode, ques []*Query) []*abstract.TreeNode {
	ans := make([]*abstract.TreeNode, len(ques))
	t.SetQueries(ques, ans)
	t.set.Init(head)
	t.SetAnswers(head, ans)
	return ans
}

func (t *Tarjan) SetQueries(ques []*Query, ans []*abstract.TreeNode) {
	for idx, q := range ques {
		if q.o1 == q.o2 {
			ans[idx] = q.o1
		}
		if q.o1 == nil {
			ans[idx] = q.o2
		}
		if q.o2 == nil {
			ans[idx] = q.o1
		}

		if _, ok := t.queries[q.o1]; !ok {
			t.queries[q.o1] = fundamentals.NewLinkedQueue()
			t.indexes[q.o1] = fundamentals.NewLinkedQueue()
		}
		if _, ok := t.queries[q.o2]; !ok {
			t.queries[q.o2] = fundamentals.NewLinkedQueue()
			t.indexes[q.o2] = fundamentals.NewLinkedQueue()
		}
		t.queries[q.o1].Enqueue(q.o2)
		t.indexes[q.o1].Enqueue(idx)
		t.queries[q.o2].Enqueue(q.o1)
		t.indexes[q.o2].Enqueue(idx)
	}
}

// 处理和head节点相关的query，存储到ans数组中
// 遍历过的节点，组成set，其最小公共祖先为head，记录到ancestors中。然后在处理head相关的query
func (t *Tarjan) SetAnswers(head *abstract.TreeNode, ans []*abstract.TreeNode) {
	if head == nil {
		return
	}

	t.SetAnswers(head.Left, ans)
	t.set.Union(head, head.Left)
	t.ancestors[t.set.Find(head)] = head

	t.SetAnswers(head.Right, ans)
	t.set.Union(head, head.Right)
	t.ancestors[t.set.Find(head)] = head

	nodes := t.queries[head]
	indexes := t.indexes[head]
	for nodes != nil && !nodes.IsEmpty() {
		node := nodes.Dequeue().(*abstract.TreeNode) // head和node组成query
		index := indexes.Dequeue().(int)             // 结果放到ans[index]
		if ancestor, ok := t.ancestors[t.set.Find(node)]; ok {
			ans[index] = ancestor
		}
	}
}

func main() {
	root := searching.CreateTreeFromArray([]string{"1", "2", "4", "#", "#", "5", "7", "#", "#", "8", "#", "#", "3", "#", "6", "9", "#", "#", "#"})
	queries := []*Query{
		&Query{root.Left.Left, root.Left.Right.Left},
		&Query{root.Left.Right.Right, root.Left.Right.Left},
		&Query{root.Left.Right.Right, root.Right.Right.Left},
		&Query{root.Right.Right.Left, root.Right},
		&Query{root.Right.Right, root.Right.Right},
		&Query{nil, root.Left.Right},
		&Query{nil, nil},
	}
	t := NewTarjan()
	ans := t.Query(root, queries)
	for i := range ans {
		fmt.Println(queries[i].o1, queries[i].o2, ans[i])
	}
}
