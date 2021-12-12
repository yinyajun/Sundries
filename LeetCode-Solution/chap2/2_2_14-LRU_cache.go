/*
Design and implement a data structure for Least Recently Used (LRU) cache. It should support the
following operations: get and set.
get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise
return -1.
set(key, value) - Set or insert the value if the key is not already present. When the cache reached its
capacity, it should invalidate the least recently used item before inserting a new item.

* @Author: Yajun
* @Date:   2021/12/12 16:22
*/

package chap2

import "container/list"

type Pair struct {
	key, val interface{}
}

// hashmap：查找O(1)
// doubly-list: 插入O(1)，单链表插入需要前驱节点

type Cache struct {
	list *list.List
	m    map[interface{}]*list.Element
	cap  int
}

func (c *Cache) Get(key interface{}) interface{} {
	n, exist := c.m[key]
	if !exist {
		return nil
	}
	val := n.Value.(Pair).val
	c.Set(key, val)
	return val
}

func (c *Cache) Set(key, val interface{}) {
	pair := Pair{key: key, val: val}

	if n, exist := c.m[key]; exist {
		c.list.Remove(n)
		c.m[key] = c.list.PushFront(pair)
	} else {
		if len(c.m) == c.cap {
			last := c.list.Back()
			delete(c.m, c.list.Remove(last).(Pair).key)
		}
		c.m[key] = c.list.PushFront(pair)
	}
}
