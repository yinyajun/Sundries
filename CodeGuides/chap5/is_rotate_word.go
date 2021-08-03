package main

import (
	"fmt"
	"strings"
)

// 时间复杂度为O(N)，但是循环了两遍
// 旋转词，类似于循环数组
// b如果是在a的基础上旋转而成，b的元素在a中对应的index应该和a的index满足同余关系
func IsRotateWord(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}

	offset := strings.Index(a, b[0:1])
	if offset == -1 {
		return false
	}

	for i := 0; i < len(a); i++ {
		if b[i] != a[(i+offset)%len(a)] {
			return false
		}
	}
	return true
}

func IsRotateWord2(a, b string) bool {
	temp := a + a
	return strings.Index(temp, b) != -1
}

func main() {
	fmt.Println(IsRotateWord("12345", "34512"))
	fmt.Println(IsRotateWord2("12345", "34512"))
}
