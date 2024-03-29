/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 10:54
 */

package utils

import (
	"CodeGuide/base/abstract"
	"reflect"
)

func Less(a, b interface{}) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		Panic("Less error: type mismatch")
	}
	switch a.(type) {
	case int:
		return a.(int) < b.(int)
	case string:
		s1, s2 := a.(string), b.(string)
		if len(s1) == len(s2) {
			return s1 < s2
		} else {
			return len(s1) < len(s2)
		}
	case float64:
		return a.(float64) < b.(float64)
	case abstract.Comparable:
		return a.(abstract.Comparable).CompareTo(b.(abstract.Comparable)) < 0
	default:
		panic("Less error: unsupported type")
		return false
	}
}

func Leq(a, b interface{}) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		Panic("Leq error: type mismatch")
	}
	switch a.(type) {
	case int:
		return a.(int) <= b.(int)
	case string:
		return a.(string) <= b.(string)
	case float64:
		return a.(float64) <= b.(float64)
	case abstract.Comparable:
		return a.(abstract.Comparable).CompareTo(b.(abstract.Comparable)) <= 0
	default:
		panic("Leq error: unsupported type")
	}
}

func Great(a, b interface{}) bool {
	return !Leq(a, b)
}

func MinInt(a ...int) int {
	if len(a) == 0 {
		Panic("invalid arguments")
	}
	m := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < m {
			m = a[i]
		}
	}
	return m
}

func MaxInt(a ...int) int {
	if len(a) == 0 {
		Panic("invalid arguments")
	}
	m := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > m {
			m = a[i]
		}
	}
	return m
}

func MaxFloat(a ...float32) float32 {
	Assert(len(a) > 0)

	m := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > m {
			m = a[i]
		}
	}
	return m
}

func MinFloat(a ...float32) float32 {
	Assert(len(a) > 0)

	m := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < m {
			m = a[i]
		}
	}
	return m
}

func CompareTo(a, b interface{}) int {
	if Less(a, b) {
		return -1
	} else if a == b {
		return 0
	}
	return 1
}
