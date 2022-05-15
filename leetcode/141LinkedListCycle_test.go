package leetcode

import "testing"

func TestHasCycle(t *testing.T) {
	one := &ListNode{Val: 3}
	two := &ListNode{Val: 2}
	three := &ListNode{Val: 0}
	four := &ListNode{Val: -4}
	two.Next = three
	three.Next = four
	four.Next = two
	one.Next = two
	if !hasCycle(one) {
		t.Fatalf("expect true get false")
	}
}
