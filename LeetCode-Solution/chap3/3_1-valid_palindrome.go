/*
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring
cases.
For example,
"A man, a plan, a canal: Panama" is a palindrome.
"race a car" is not a palindrome.
Note: Have you consider that the string might be empty? This is a good question to ask during an
interview.
For the purpose of this problem, we define empty string as valid palindrome.

* @Author: Yajun
* @Date:   2021/12/12 17:15
*/

package chap3

import (
	"strings"
	"unicode"
)

func isValidPalindrome(str string) bool {
	if len(str) == 0 {
		return true
	}

	str = strings.ToLower(str)

	left, right := 0, len(str)-1
	for left < right {
		if !isAlNum(str[left]) {
			left++
			continue
		}
		if !isAlNum(str[right]) {
			right--
			continue
		}
		if str[left] == str[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

func isAlNum(a byte) bool {
	r := rune(a)
	if unicode.IsDigit(r) || unicode.IsLetter(r) {
		return true
	}
	return false
}
