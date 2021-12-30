/*
Given two binary strings, return their sum (also a binary string).
For example,
a = "11"
b = "1"
Return "100".

* @Author: Yajun
* @Date:   2021/12/28 10:00
*/

package chap3

import (
	"strings"
)

func addBinary(a, b string) string {
	var (
		res   strings.Builder
		carry uint8
	)

	if len(a) < len(b) {
		b, a = a, b
	}
	// len(a) >= len(b)
	a = reverse(a)
	b = reverse(b)

	for i := 0; i < len(a); i++ {
		aa := a[i] - '0'
		bb := uint8(0)
		if i < len(b) {
			bb = b[i] - '0'
		}
		t := aa + bb + carry
		val := t % 2
		carry = t / 2
		res.WriteByte(val + '0')
	}
	if carry == 1 {
		res.WriteByte('1')
	}
	return reverse(res.String())
}

func addBinaryB(a, b string) string {
	var (
		res   strings.Builder
		carry uint8
	)
	sz := len(a)
	if len(b) > sz {
		sz = len(b)
	}

	for i := 0; i < sz; i++ {
		ai := uint8(0)
		if i < len(a) {
			ai = a[len(a)-1-i]
		}
		bi := uint8(0)
		if i < len(b) {
			bi = b[len(b)-1-i]
		}
		t := ai + bi + carry
		val := t % 2
		carry = t / 2
		res.WriteByte(val + '0')
	}
	if carry > 0 {
		res.WriteByte('1')
	}
	return reverse(res.String())
}

func reverse(a string) string {
	runes := []rune(a)
	for i, j := 0, len(runes)-1; i < j; {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}
