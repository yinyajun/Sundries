package main

import (
	"fmt"
	"math/rand"
)

type node struct {
	val  int
	next *node
}

type MessageBox struct {
	headMap  map[int]*node
	tailMap  map[int]*node
	printNum int
}

func (b *MessageBox) Receive(v int) {
	fmt.Println("receive", v)
	cur := &node{
		val:  v,
		next: nil,
	}

	b.Merge(cur)
	b.Print()
}

func (b *MessageBox) Merge(cur *node) {
	b.headMap[cur.val] = cur
	b.tailMap[cur.val] = cur

	// merge tail
	if n, ok := b.tailMap[cur.val-1]; ok {
		n.next = cur
		delete(b.tailMap, cur.val-1)
		delete(b.headMap, cur.val)
	}

	// merge head
	if n, ok := b.headMap[cur.val+1]; ok {
		cur.next = n
		delete(b.headMap, cur.val+1)
		delete(b.tailMap, cur.val)
	}
}

func (b *MessageBox) Print() {
	n, ok := b.headMap[b.printNum+1]
	if ok {
		begin := n.val
		for n != nil {
			fmt.Print(n.val, " ")
			b.printNum = n.val
			n = n.next
		}
		end := b.printNum
		delete(b.headMap, begin)
		delete(b.tailMap, end)
		fmt.Println()
	}
}

func KnuthShuffle(arr []int) {
	var r int
	for i := len(arr) - 1; i >= 0; i-- {
		r = rand.Intn(i + 1)
		arr[r], arr[i] = arr[i], arr[r]
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	KnuthShuffle(arr)
	fmt.Println(arr)

	box := &MessageBox{
		headMap:  make(map[int]*node),
		tailMap:  make(map[int]*node),
		printNum: 0,
	}

	for _, v := range arr {
		box.Receive(v)
	}
}
