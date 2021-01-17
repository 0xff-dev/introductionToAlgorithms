package leetcode

import "testing"

func TestReverseBetween(t *testing.T) {
	l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head := reverseBetween(l, 2, 4)
	head.Dis()

	l = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head = reverseBetween(l, 1, 5)
	head.Dis()

	l = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head = reverseBetween(l, 3, 3)
	head.Dis()

	l = &ListNode{1, nil}
	head = reverseBetween(l, 1, 1)
	head.Dis()
}
