/*
* @Author: Yajun
* @Date:   2021/11/27 22:05
 */

package utils

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

func NewListNode(val interface{}) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func ShowList(node *ListNode) {
	for cur := node; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, "  ")
	}
	fmt.Println()
}

func NewListFromSlice(a interface{}) *ListNode {
	rv := reflect.ValueOf(a)
	length := rv.Len()
	dummy := new(ListNode)
	pre := dummy
	for i := 0; i < length; i++ {
		pre.Next = NewListNode(rv.Index(i).Interface())
		pre = pre.Next
	}
	return dummy.Next
}
