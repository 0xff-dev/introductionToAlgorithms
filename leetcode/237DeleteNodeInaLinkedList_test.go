package leetcode

import "testing"

func TestDeleteNode1(t *testing.T) {
	head := &ListNode{4, nil}
	del := &ListNode{5, &ListNode{1, &ListNode{9, nil}}}
	head.Next = del
	deleteNode1(del)
	head.Dis()
}
