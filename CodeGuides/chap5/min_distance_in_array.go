package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 从左到右遍历
// 分别记录最近一次遇到的str1，str2
// 如果遍历到str1，就比较cur-last1
// 如果遍历到str2，就比较cur-last2
func MinDistance(a []string, str1, str2 string) int {
	last1, last2 := -1, -1
	dist := len(a)
	for i := range a {
		if a[i] == str1 {
			last1 = i
			if last2 != -1 {
				dist = utils.MinInt(dist, last1-last2)
			}
		}
		if a[i] == str2 {
			last2 = i
			if last1 != -1 {
				dist = utils.MinInt(dist, last2-last1)
			}
		}
	}
	if last1 == -1 || last2 == -1 {
		return -1
	}
	return dist
}

type Record struct {
	m map[string]map[string]int
}

func NewRecord(a []string) *Record {
	m := make(map[string]map[string]int)
	r := &Record{m}
	indexMap := make(map[string]int)
	for i := range a {
		r.Update(a[i], indexMap, i)
		indexMap[a[i]] = i
	}
	return r
}

func (r *Record) Update(key string, indexMap map[string]int, index int) {
	if _, ok := r.m[key]; !ok {
		r.m[key] = make(map[string]int)
	}
	keyMap := r.m[key]
	for k, idx := range indexMap {
		if k == key {
			continue
		}
		kMap := r.m[k]
		curMin := index - idx
		if _, ok := keyMap[k]; !ok {
			keyMap[k] = curMin
			kMap[key] = curMin
		} else {
			min := keyMap[k]
			if curMin < min {
				keyMap[k] = curMin
				kMap[key] = curMin
			}
		}
	}
}

func (r *Record) Query(a, b string) int {
	if a == b {
		return 0
	}
	if m, ok := r.m[a]; ok {
		if ret, ok := m[b]; ok {
			return ret
		}
	}
	return -1
}

func main() {
	a := []string{"1", "3", "4", "3", "2", "3", "1"}
	fmt.Println(MinDistance(a, "4", "2"))
	r := NewRecord(a)
	fmt.Println(r.Query("4", "2"))
}
