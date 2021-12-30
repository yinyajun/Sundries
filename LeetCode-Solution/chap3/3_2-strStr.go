/*
Implement strStr().
Returns a pointer to the first occurrence of needle in haystack, or null if needle is not part of haystack.

* @Author: Yajun
* @Date:   2021/12/12 19:14
*/

package chap3

// 暴力解法
// time: O(mn); space: O(1)
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

// KMP1 解法(DFA)
// time: O(n); space: O(mn)
func strStrB(haystack, needle string) int {
	return NewKMP1(needle).search(haystack)
}

type KMP1 struct {
	DFA [][]int // row: 字符集；column：状态数（m+1）
	// 总共有【0，1，...,m】共m+1个状态，其中m状态为停机状态
}

// 每个状态代表pattern能匹配到第几个字母，当匹配到停机状态就说明匹配成功
// 所有非停机状态[0...m-1]，都会有状态转移，所以有m个状态需要状态转移

// dfa[c][j]: 当前状态j，意味着已经匹配了j个字符pat[0...j)，下一个待匹配字符为c。根据pattern[j]和c的值，会跳转到不同的状态
// * 如果匹配，会跳转到下一个状态j+1: dfa[c][j] = j+1
// * 如果失配，pattern会回退，而txt不回退，为了尽可能减少pattern的回退, 将pattern回退到重启状态x: dfa[c][j] = dfa[c][x]
//        此时pattern的前缀[0...x]和text后缀[...j]对齐（所以需要记录公共后缀，dfa值的另一种理解）

func NewKMP1(pattern string) *KMP1 {
	r := 256
	m := len(pattern)
	kmp := &KMP1{DFA: make([][]int, r)}

	// init DFA
	for i := 0; i < r; i++ {
		kmp.DFA[i] = make([]int, m)
	}

	// base
	kmp.DFA[int(pattern[0])][0] = 1
	// 在状态0（未匹配任何字符）
	// * 如果遇到字符为为pattern[0]，那么能够匹配，进入状态1
	// * 遇到其他字符必然不匹配，所以仍然转移到状态0

	// construct dfa
	x := 0 // 重启状态初始化为0
	for j := 1; j < m; j++ {
		for c := 0; c < r; c++ {
			kmp.DFA[c][j] = kmp.DFA[c][x] // 默认失配，回到重启状态
		}
		kmp.DFA[int(pattern[j])][j] = j + 1 // 更新匹配状态
		x = kmp.DFA[int(pattern[j])][x]     // 更新重启状态（重启状态比当前状态j慢一个状态，也就是说，重启状态的dfa已经被更新过）
		// 重启状态意味着未匹配前的最大公共前后缀，此时遇到pattern[j]，它会怎么更新？直接调用DFA[][x]即可
	}
	return kmp
}

func (k *KMP1) search(text string) int {
	n := len(text)
	m := len(k.DFA[0])

	j := 0 // pattern的初始状态
	for i := 0; i < n; i++ {
		j = k.DFA[int(text[i])][j] // 计算pattern的下一个状态
		if j == m {                // 到达停机状态
			return i - m + 1
		}
	}
	return n
}

// 构建DFA的过程和search的过程十分相似
// search在text中匹配pattern
// construct在pattern[1...]中匹配pattern（即找公共前后缀）

// kmp2解法(next数组)
// time: O(n); space: O(m)
func strStrC(haystack, needle string) int {
	return kmp2(haystack, needle)
}

func kmp2(text, pattern string) int {
	n := len(text)
	m := len(pattern)
	next := getNext(pattern)

	j := 0
	for i := 0; i < n; i++ {
		for j > 0 && text[i] != pattern[j] {
			j = next[j] // 遇到失配，查询next数组获取公共前缀的长度（或者重启位置）
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return n
}

// next[i]: pattern[0...i)已经匹配，在已经匹配的串中的最大公共前后缀的长度(或者认为是重启位置)
// next[i] = pre    =>    pattern[0...pre-1] = pattern[i-pre...i-1]
// * if pattern[pre]==pattern[i], next[i+1]=pre+1
// * if pattern[pre]!=pattern[i], 不能在已有的最大公共前后缀上构成新的最大公共前后缀
// 	    需要在当前的最大公共前后缀的子集中寻找，pre = next[pre], 直到pattern[i] == pattern[pre]，此时next[i]=pre+1
func getNext(pattern string) []int {
	m := len(pattern)
	next := make([]int, m+1)

	j := 0
	for i := 2; i < m+1; i++ {
		for j != 0 && pattern[j] != pattern[i-1] {
			j = next[j]
		}
		if pattern[i-1] == pattern[j] {
			j++
		}
		next[i] = j
	}
	return next
}

func strStrD(haystack, needle string) int {
	return kmp(haystack, needle)
}

func kmp(text, pattern string) int {
	n := len(text)
	next := genNext(pattern) // next[i]: 已经匹配pattern[0...i]的最长公共前缀

	j := 0
	for i := 0; i < n; i++ {
		for j > 0 && text[i] != pattern[j] {
			j = next[j-1]
		}
		// j ==0 || text[i] == pattern[j]
		if text[i] == pattern[j] {
			j++
		}
		if j == len(pattern) {
			return i - len(pattern) + 1
		}
	}
	return n
}

// next[i]: pattern[0...i]的最长公共前缀
// next[i] = pre
// if pat[pre] == pat[i+1], next[i+1] = pre +1
// else, pre = next[pre] ... util, pat[pre] == pat[i+1], next[i+1] = pre+1
func genNext(pattern string) []int {
	m := len(pattern)
	next := make([]int, m)

	j := 0
	for i := 1; i < m; i++ {
		for j > 0 && pattern[j] != pattern[i] {
			j = next[j-1]
		}
		// j ==0 || pattern[j] == pattern[i]
		if pattern[i] == pattern[j] {
			j++
		}
		next[i] = j
	}
	return next
}
