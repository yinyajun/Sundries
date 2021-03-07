package main

import "fmt"

// 123abcabc, abc->x, 123x

// 时间复杂度为O(N)
func SubstituteString(a, from, to string) string {
	match := 0

	// copy
	aa := make([]byte, len(a))
	for i := range a {
		aa[i] = a[i]
	}

	// 将a中和from一样的部分全部置为空
	for i := range aa {
		if aa[i] == from[match] {
			match++
			if match == len(from) {
				for k := 0; k < len(from); k++ {
					aa[i-k] = 0
				}
				match = 0
			}
		} else {
			match = 0
		}
	}

	//
	ret := []byte{}
	for i := 0; i < len(aa); i++ {
		if aa[i] != 0 {
			ret = append(ret, aa[i])
			continue
		}
		if i == 0 || aa[i-1] != 0 { // 替换的边界
			ret = append(ret, to...)
		}
	}

	return string(ret)
}

func main() {
	fmt.Println(SubstituteString("123abcabc", "abc", "X"))
}
