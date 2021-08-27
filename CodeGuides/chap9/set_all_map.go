package main

// 为map添加setAll功能，将所有key的value设置一个值
// 改造值，添加时间作为版本

type myValue struct {
	v interface{}
	t int
}

type MyMap struct {
	a    map[string]myValue
	time int
	all  myValue
}

func NewMyMap() MyMap {
	return MyMap{
		a:    make(map[string]myValue),
		time: 0,
		all:  myValue{nil, -1},
	}
}

func (m MyMap) get(key string) interface{} {
	v, ok := m.a[key]
	if ok {
		if m.all.t > v.t { // all值更新
			return m.all.v
		}
		return v.v
	} else {
		return nil
	}
}

func (m MyMap) put(key string, val interface{}) {
	m.time++
	m.a[key] = myValue{v: val, t: m.time}
}

func (m MyMap) setAll(val interface{}) {
	m.time++
	m.all = myValue{val, m.time}
}

func (m MyMap) contains(key string) bool {
	_, ok := m.a[key]
	return ok
}
