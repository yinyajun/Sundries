package main

import (
	"CodeGuide/base/utils"
)

//* insert, 不重复加入结构
//* delete
//* random, 等概返回结构中任意一个key
//所有操作时间为O(1)

// insert和delete 做到O(1)很容易
// 使用map，同时也能做到不重复

// 这样的话，random操作就比较困难了，从所有数中，通过随机数来获取index
// 而引入index，最简单的就是数组了(类似于索引优先队列)

// 为了考虑拓展性，需要数组可以自动扩缩容
// 为了简单起见，使用golang中的slice来代替动态数组

type pool struct {
	hash map[interface{}]int // element -> idx
	data []interface{}       // idx -> element
	size int
}

func NewPool() *pool {
	return &pool{make(map[interface{}]int), make([]interface{}, 0), 0}
}

func (p *pool) Insert(key interface{}) {
	_, ok := p.hash[key]
	if ok {
		return
	}
	p.hash[key] = p.size
	p.data = append(p.data, key)
	p.size++
}

func (p *pool) Delete(key interface{}) {
	idx, ok := p.hash[key]
	if !ok {
		return
	}
	delete(p.hash, key)
	// 将该元素交换到最后一位
	p.data[idx], p.data[p.size-1] = p.data[p.size-1], p.data[idx]
	p.size--
	p.data = p.data[:p.size]
}

func (p *pool) Random() interface{} {
	idx := utils.Random.Intn(p.size)
	return p.data[idx]
}

//func main() {
//	pool := NewPool()
//	pool.Insert(1)
//	pool.Insert(3)
//	pool.Insert(154)
//	pool.Insert(56)
//
//
//	fmt.Println(pool.hash)
//	fmt.Println(pool.data)
//
//	pool.Delete(56)
//
//	fmt.Println(pool.hash)
//	fmt.Println(pool.data)
//
//	count := map[interface{}]int{}
//	for i :=0 ; i < 10000; i++{
//		count[pool.Random()]+=1
//	}
//	fmt.Println(count)
//}
