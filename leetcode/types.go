package leetcode

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Dis() {
	for walker := l; walker != nil; walker = walker.Next {
		log.Printf("val is [%d]", walker.Val)
	}
}
