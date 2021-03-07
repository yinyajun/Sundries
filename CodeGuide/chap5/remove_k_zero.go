package main

import "fmt"

// "A0000B000" =》 "A0000B"

// 遇到0就计数，直到遇到非0，如果此时cnt==k，就舍弃，否则就加上cnt个0，然后加上这个非0元素，在清空cnt
// 如果最后cnt!=k,需要将cnt个0补上
// 时间复杂度为O(n)
func RemoveKZero(a string, k int) string {
	cnt := 0 // 连0计数器
	ret := []byte{}
	for i := range a {
		if a[i] == '0' {
			cnt += 1
		} else { // 遇到非0元素，开始处理连0
			if cnt != k {
				for i := 0; i < cnt; i++ {
					ret = append(ret, '0')
				}
			}
			ret = append(ret, a[i])
			cnt = 0
		}
	}
	// 遇到非0元素才处理连0，如果a以0结尾，那么最后一组0可能没有处理
	if cnt != k {
		for i := 0; i < cnt; i++ {
			ret = append(ret, '0')
		}
	}
	return string(ret)
}

func main() {
	fmt.Println(RemoveKZero("A0000B000", 3))
}
