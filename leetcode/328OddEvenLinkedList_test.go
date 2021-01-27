package leetcode

import "testing"

func TestOddEvenList(t *testing.T) {
	l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head := oddEvenList(l)
	head.Dis()
	l = &ListNode{2, &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{6, &ListNode{4, &ListNode{7, nil}}}}}}}
	head = oddEvenList(l)
	head.Dis()
	if head = oddEvenList(nil); head != nil {
		t.Fatalf("should return nil")
	}

	l = &ListNode{1, nil}
	head = oddEvenList(l)
	head.Dis()

	l = &ListNode{1, &ListNode{2, nil}}
	head = oddEvenList(l)
	head.Dis()

	l = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	head = oddEvenList(l)
	head.Dis()

	l = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	head = oddEvenList(l)
	head.Dis()
}
