package main

import "fmt"

// 两个字符串是否是异形词，先将str1转变为hist，然后在遍历str2的时候，将hist消除，如果遇到负数，就false

func IsDeformation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	m := make([]int, 256)
	for i := range str1 {
		m[int(str1[i])] += 1
	}
	for i := range str2 {
		m[int(str2[i])] -= 1
		if m[int(str2[i])] < 0 {
			return false
		}
	}
	return true
}

func IsDeformation2(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	m := make(map[uint8]int)
	for i := range str1 {
		m[str1[i]] += 1
	}
	for i := range str2 {
		m[str2[i]] -= 1
		if m[str2[i]] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsDeformation("ffyt", "tyfe"))
	fmt.Println(IsDeformation2("ffyt", "tyff"))
}
