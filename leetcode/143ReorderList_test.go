package leetcode

import (
	"testing"
)

func TestReorderList(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	reorderList(head)
	head.Dis()

	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	reorderList(head)
	head.Dis()

	head = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	reorderList(head)
	head.Dis()
}
