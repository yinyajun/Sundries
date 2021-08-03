package main

import (
	"errors"
	"fmt"
	"hash/fnv"
	"strconv"
	"sync"

	"github.com/emirpasic/gods/maps/treemap"
	"github.com/emirpasic/gods/utils"
)

/*
1. 将id映射到2^32的范围内，对应到环中的一个位置
2. 顺时针找最近的机器
*/

var EmptyCircle = errors.New("circle is empty")

type Circle struct {
	sync.RWMutex
	circle     *treemap.Map
	replicaNum int                 // 虚拟节点数目
	count      int                 // 真实节点数目
	members    map[string]struct{} // 真实节点set
}

func NewCircle() *Circle {
	c := new(Circle)
	c.circle = treemap.NewWith(utils.UInt32Comparator)
	c.replicaNum = 20
	c.members = make(map[string]struct{})
	return c
}

func (c *Circle) hashFunc(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func (c *Circle) virtualNode(realNode string, idx int) string { return realNode + strconv.Itoa(idx) }

func (c *Circle) add(realNode string) {
	for i := 0; i < c.replicaNum; i++ {
		vn := c.virtualNode(realNode, i)
		c.circle.Put(c.hashFunc(vn), realNode)
	}
	c.members[realNode] = struct{}{}
	c.count++
}

func (c *Circle) remove(realNode string) {
	for i := 0; i < c.replicaNum; i++ {
		vn := c.virtualNode(realNode, i)
		c.circle.Remove(c.hashFunc(vn))
	}
	delete(c.members, realNode)
	c.count--
}

func (c *Circle) AddNode(realNode string) {
	c.Lock()
	c.add(realNode)
	c.Unlock()
}

func (c *Circle) RemoveNode(realNode string) {
	c.Lock()
	c.add(realNode)
	c.Unlock()
}

func (c *Circle) Empty() bool { return c.count == 0 }

func (c *Circle) Get(key string) (string, error) {
	c.RLock()
	defer c.RUnlock()
	if c.Empty() {
		return "", EmptyCircle
	}
	return c.search(c.hashFunc(key)), nil
}

func (c *Circle) search(key uint32) string {
	_, n := c.circle.Ceiling(key)
	if n == nil {
		_, n = c.circle.Min()
	}
	return n.(string)
}

func (c *Circle) Members() (nodes []string) {
	c.RLock()
	defer c.RUnlock()
	for key := range c.members {
		nodes = append(nodes, key)
	}
	return
}

func main() {
	c := NewCircle()

	c.AddNode("hjchat02.add.*****.*****.net")
	c.AddNode("p28877v.hulk.bjyt.*****.net")
	c.AddNode("hjchat03.add.*****.*****.net")
	c.AddNode("p16727v.hulk.bjyt.*****.net")

	fmt.Println(c.Get("123"))
	fmt.Println(c.Get("234"))
	fmt.Println(c.Members())
}
