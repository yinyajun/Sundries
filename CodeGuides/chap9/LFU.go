package main

// LRU的核心就是哈希链表，相对比较简单
// 而LFU 算法的淘汰策略是 Least Frequently Used，也就是每次淘汰那些使用次数最少的数据。

// 实现get和put方法
// 1. 任何操作都会使freq增加
// 2. 当缓存达到容量的时候看，移除freq最小的pair；如果对应了多个pair，移除最旧的那个

// 1. KV: map[key]value , O(1)时间获取val
// 2. KF: map[key]freq, O(1)时间获取freq
// 3. 删除freq最小的key，首先要维护minFreq变量，可以O(1)时间知道最小频次；
//   其次，多个pair都有相同freq，也就是说freq->pair是一对多的关系，且pair之间保留插入顺序
//	 map[freq]list可以吗？ 乍一看是可以的
//   假如操作过某个key，对应的pair需要从freq对应的list删除，并加入到freq对应的list中，这些操作仅用list是无法保证O(1)的
//   具体而言，在list中删除任意元素，不是O(1)的，而linkedHashMap可以做到
//   所以FK: map[freq]linkedHashMap

type LinkedHashSet struct {
	list *DoublyLinkedList
	hash map[interface{}]*Node
}

func NewLinkedHashSet() *LinkedHashSet {
	return &LinkedHashSet{
		list: NewDoublyLinkedList(),
		hash: make(map[interface{}]*Node),
	}
}

func (m *LinkedHashSet) Add(key interface{}) {
	node, ok := m.hash[key]
	if ok {
		m.list.Remove(node)
		m.list.PushBack(node)
		return
	}
	node = &Node{Key: key}
	m.list.PushBack(node)
	m.hash[key] = node
}

func (m *LinkedHashSet) Remove(key interface{}) {
	node, ok := m.hash[key]
	if !ok {
		return
	}
	m.list.Remove(node)
	delete(m.hash, key)
}

type LFU struct {
	KeyToVal   map[interface{}]interface{}
	KeyToFreq  map[interface{}]int
	FreqToKeys map[int]*LinkedHashSet
	cap        int
	minFreq    int
}

func NewLFU(cap int) *LFU {
	return &LFU{
		make(map[interface{}]interface{}),
		make(map[interface{}]int),
		make(map[int]*LinkedHashSet),
		cap,
		0,
	}
}

func (c *LFU) Get(key interface{}) (interface{}, bool) {
	val, ok := c.KeyToVal[key]
	if !ok {
		return nil, ok
	}
	c.increaseFreq(key) // add freq
	return val, ok
}

func (c *LFU) Put(key, val interface{}) {
	_, ok := c.KeyToVal[key]
	if ok {
		c.KeyToVal[key] = val
		c.increaseFreq(key) // add freq
		return
	}
	//	no exist
	if len(c.KeyToVal) == c.cap {
		c.removeMinFreq()
	}

	c.KeyToVal[key] = val
	c.KeyToFreq[key] = 1
	if _, ok := c.FreqToKeys[1]; !ok {
		c.FreqToKeys[1] = NewLinkedHashSet()
	}
	c.FreqToKeys[1].Add(key)
	c.minFreq = 1 // 插入新key后，minFreq肯定为1
}

func (c *LFU) increaseFreq(key interface{}) {
	freq := c.KeyToFreq[key]
	// update FK
	// remove from FK[freq]
	oldKeys := c.FreqToKeys[freq]
	oldKeys.Remove(key)

	if oldKeys.list.Size == 0 {
		delete(c.FreqToKeys, freq)
		if freq == c.minFreq {
			c.minFreq++
		}
	}
	// add in FK[freq+1]
	if _, ok := c.FreqToKeys[freq+1]; !ok {
		c.FreqToKeys[freq+1] = NewLinkedHashSet()
	}
	c.FreqToKeys[freq+1].Add(key)

	// update KF
	c.KeyToFreq[key] += 1
}

func (c *LFU) removeMinFreq() {
	keys := c.FreqToKeys[c.minFreq]
	toDel := keys.list.Root.Next.Key

	// update FK
	keys.Remove(toDel)
	if keys.list.Size == 0 {
		delete(c.FreqToKeys, c.minFreq)
		// update minFreq
		// 需要更新吗？更新的时间复杂度肯定不止O(1)，但是不需要更新，因为该函数仅在put中使用，put中插入新元素，必然会更新minFreq
	}
	// update KF
	delete(c.KeyToFreq, toDel)

	// update KV
	delete(c.KeyToVal, toDel)
}

//func main() {
//	cache := NewLFU(2)
//
//	cache.Put(9, 10)
//	cache.Put(8, 20)
//
//	fmt.Println(cache.Get(9))
//
//	cache.Put(7, 30)
//
//	fmt.Println(cache.Get(8))
//}
