package main

import "fmt"

// 单词间逆序
func ReverseString(a []byte) {
	reverse(a)
	i, j := 0, 0
	for j <= len(a)-1 {
		if a[j] != ' ' {
			j++
		} else { // a[j]=' '
			if i < j { // 每当遇到空格且i<j的时候，反转空格之前的单词
				reverse(a[i:j])
				i = j
			}
			i++
			j++
		}
	}
	if i < j { // 最后一个单词需要单独反转(如果没有以空格结尾上述循环中就不会反转最后一个单词)
		reverse(a[i:j])
	}
	fmt.Println(string(a))
}

func ReverseString2(a []byte) {
	cond := func(expr bool, i, j int) int {
		if expr {
			return i
		}
		return j
	}
	reverse(a)
	i, j := -1, -1 // 注意初值，0不可以，因为i==0合法
	// 在循环中动态确定当前遍历到的单词的左右边界
	// 左边界为：k有值，k-1为space或者k==0
	// 右边界为：k有值，k+1为space或者k==len-1
	for k := 0; k < len(a); k++ {
		if a[k] != ' ' {
			i = cond(k == 0 || a[k-1] == ' ', k, i)
			j = cond(k == len(a)-1 || a[k+1] == ' ', k, j)
		}
		if i != -1 && j != -1 {
			reverse(a[i : j+1])
			i, j = -1, -1
		}
	}
	fmt.Println(string(a))
}

func AdjustString(a []byte, size int) {
	if size > len(a) {
		size = len(a)
	}
	reverse(a)
	reverse(a[:len(a)-size])
	reverse(a[len(a)-size:])
	fmt.Println(string(a))
}

func reverse(a []byte) {
	i, j := 0, len(a)-1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}

func main() {
	a := []byte("dog loves pig")
	ReverseString2(a)
	AdjustString([]byte("ABCDE"), 3)
}
