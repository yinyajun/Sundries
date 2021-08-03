package main

type bitmap struct {
	size uint
	cap  uint
	m    []uint
}

func NewBitmap(capacity uint) *bitmap {
	// 可以表征0~capacity-1总共capacity个数
	// cap用uint表示，最多表示42亿
	// 需要(capacity/32+1)个int
	return &bitmap{size: 0, cap: capacity, m: make([]uint, capacity/32+1)}
}

// 插入，先找到word所在的index和offset，然后将word和1<<offset按位或
func (b *bitmap) Add(id uint) {
	if id > b.cap-1 {
		panic("id is too large")
	}
	idx, offset := id/32, id%32
	b.m[idx] |= 1 << offset
	b.size++
}

// 删除，将word & (~(1<<offset))，对应位置置零
func (b *bitmap) Remove(id uint) {
	idx, offset := id/32, id%32
	b.m[idx] &= ^(1 << offset)
	b.size--
}

func (b *bitmap) Find(id uint) bool {
	idx, offset := id/32, id%32
	return (b.m[idx] & (1 << offset)) > 0
}

//func main() {
//	b := NewBitmap(96)
//	fmt.Println(b.m)
//	b.Add(63)
//	fmt.Println(b.m)
//	fmt.Println(b.Find(63))
//}
