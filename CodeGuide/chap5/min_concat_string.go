package main

import (
	"fmt"
	"sort"
)

// ["b", "ba"] = > "bab" "bba"
// 为了使生成的字符串的字典序最小，首先字符串的开头字符要最小
// 不难想到要对strs中的所有字符串排序

// 直接排序可能造成错误答案，bba>bab
// 需要对其concat后的字符串比大小，才能是正确答案

func MinConcatString(a []string) {
	less := func(i, j int) bool {
		return a[i]+a[j] < a[j]+a[i]
	}
	sort.Slice(a, less)
	res := ""
	for i := range a {
		res += a[i]
	}
	fmt.Println(res)
}

func main() {
	MinConcatString([]string{"ba", "b"})
}
