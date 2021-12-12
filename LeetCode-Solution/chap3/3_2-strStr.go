/*
Implement strStr().
Returns a pointer to the first occurrence of needle in haystack, or null if needle is not part of haystack.

* @Author: Yajun
* @Date:   2021/12/12 19:14
*/

package chap3

func strStr(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		var j int
		for j = i; (j-i < len(needle)) && haystack[j] == needle[j-i]; j++ {
		}
		if j-i == len(needle) {
			return i
		}
	}
	return -1
}
