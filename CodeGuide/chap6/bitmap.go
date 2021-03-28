package main

import "fmt"

type bitmap struct {
	size int32
	cap  int32
	m    []int32
}

func NewBitmap(capacity int32) *bitmap {
	// 可以表征0~capacity总共capacity+1个数
	// cap用int表示，最多表示21亿
	// 需要(capacity/32+1)个int
	return &bitmap{size: 0, cap: capacity, m: make([]int32, capacity/32+1)}
}

// 插入，先找到word所在的index和offset，然后将word和1<<offset按位或
func (b *bitmap) Add(id int32) {
	if id >= b.cap {
		panic("id is too large")
	}
	idx, offset := id/32, id%32
	b.m[idx] |= 1 << offset
	b.size++
}

// 删除，将word & (~(1<<offset))，对应位置置零
func (b *bitmap) Remove(id int32) {
	idx, offset := id/32, id%32
	b.m[idx] &= ^(1 << offset)
	b.size--
}

func (b *bitmap) Find(id int32) bool {
	idx, offset := id/32, id%32
	return (b.m[idx] & (1 << offset)) > 0
}

func main() {
	b := NewBitmap(10)
	b.Add(1)
	b.Add(0)
	fmt.Println(b.m)
	fmt.Println(b.Find(4))
	fmt.Println(b.Find(1))
	b.Remove(1)
	fmt.Println(b.m)
	fmt.Println(b.Find(1))
	fmt.Println(b.Find(4))
	fmt.Println(b.Find(0))
}
