/*
* @Author: Yajun
* @Date:   2021/12/12 17:34
 */

package chap3

import (
	"fmt"
	"testing"
)

func Test3_1_1(t *testing.T) {
	str := "A man, a plan, a canal: Panama"
	fmt.Println(isValidPalindrome(str))
}

func Test3_1_2(t *testing.T) {
	str := "ATGTGAGCTGGTGTGTGCFAA"
	pattern := "GTGTGC"
	fmt.Println(strStr(str, pattern))
	fmt.Println(strStrB(str, pattern))
	fmt.Println(strStrC(str, pattern))
	fmt.Println(strStrD(str, pattern))
}

func Test3_1_3(t *testing.T) {
	fmt.Println(atoi("-3924X8fc"))
	fmt.Println(atoi("+413"))
	fmt.Println(atoi("++c"))
	fmt.Println(atoi("++1"))
	fmt.Println(atoi("-2147483648"))
}

func Test3_1_4(t *testing.T) {
	fmt.Println(addBinary("110", "10"))
	fmt.Println(addBinaryB("110", "10"))
}

func Test3_1_5(t *testing.T) {
	str := "feuislghahgiuagrefg"
	fmt.Println(longPalindromicSubstring(str))
	fmt.Println(longPalindromicSubstringB(str))
	fmt.Println(longPalindromicSubstringC(str))
	fmt.Println(longPalindromicSubstringD(str))
	fmt.Println(longPalindromicSubstringE(str))
	fmt.Println(longPalindromicSubstringF(str))
	fmt.Println(longPalindromicSubstringG(str))
	fmt.Println(longPalindromicSubstringH(str))
	fmt.Println(longPalindromicSubstringI(str))
}

func Test3_1_6(t *testing.T) {
	text := "ab"
	pattern := "a*b*"
	fmt.Println(isMatch(text, pattern))
	fmt.Println(isMatchB(text, pattern))
	fmt.Println(isMatchC(text, pattern))
	fmt.Println(isMatchD(text, pattern))
	fmt.Println(isMatchE(text, pattern))
}

func Test3_1_7(t *testing.T) {
	text := "adceb"
	pattern := "*a*?"
	fmt.Println(isMatch2(text, pattern))
	//fmt.Println(isMatch2B(text, pattern))
	fmt.Println(isMatch2C(text, pattern))
	fmt.Println(isMatch2D(text, pattern))
	fmt.Println(isMatch2E(text, pattern))
	fmt.Println(isMatch2E2(text, pattern))
	fmt.Println(isMatch2F(text, pattern))
	fmt.Println(isMatch2F2(text, pattern))
	fmt.Println(isMatch2G(text, pattern))
	fmt.Println(isMatch2H(text, pattern))
	fmt.Println(isMatch2I(text, pattern))
}
