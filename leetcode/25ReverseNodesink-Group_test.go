package leetcode

import "testing"

func TestReverseKGroup(t *testing.T) {
	l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	r := reverseKGroup(l, 2)
	r.Dis()

	l = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	r = reverseKGroup(l, 1)
	r.Dis()

	l = &ListNode{1, nil}
	r = reverseKGroup(l, 2)
	r.Dis()
}
