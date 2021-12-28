package leetcode

import "testing"

func TestMiddleNode(t *testing.T) {
	l := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	m := middleNode(l)
	if m.Val != 2 {
		t.Fatalf("expect 2 get %d", m.Val)
	}

	l = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	m = middleNode(l)
	if m.Val != 3 {
		t.Fatalf("expect 3 get %d", m.Val)
	}

	l = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}
	m = middleNode(l)
	if m.Val != 4 {
		t.Fatalf("expect 4 get %d", m.Val)
	}
}
