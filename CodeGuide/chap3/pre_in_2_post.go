package main

import "fmt"

func PreIn2Post(pre, in []int) []int {

	m := make(map[int]int)
	for i, num := range in {
		m[num] = i
	}

	post := make([]int, len(pre))
	preIn2Post(pre, 0, len(pre)-1, in, 0, len(in)-1, post, 0, len(post)-1, m)
	return post
}

func preIn2Post(pre []int, preLo, preHi int, in []int, inLo, inHi int, post []int, postLo, postHi int, m map[int]int) {
	if preLo > preHi {
		return
	}

	post[postHi] = pre[preLo]
	index := m[pre[preLo]]
	preIn2Post(pre, preLo+1+index-inLo-1+1, preHi, in, index+1, inHi, post, postLo+index-inLo, postHi-1, m)
	preIn2Post(pre, preLo+1, preLo+1+index-inLo-1, in, inLo, index-1, post, postLo, postLo+index-inLo-1, m)
}

func main() {
	pre := []int{1, 2, 4, 5, 3, 6, 7}
	in := []int{4, 2, 5, 1, 6, 3, 7}
	post := PreIn2Post(pre, in)
	fmt.Println(post)
}
