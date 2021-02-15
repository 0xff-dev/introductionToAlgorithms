package leetcode

import "testing"

func TestMergeTwoLists(t *testing.T) {
	l1, l2 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	r := mergeTwoLists(l1, l2)
	r.Dis()

	l1, l2 = nil, nil
	r = mergeTwoLists(l1, l2)
	if r != nil {
		t.Fatal("should return nil")
	}
	l1, l2 = nil, &ListNode{1, nil}
	r = mergeTwoLists(l1, l2)
	r.Dis()
}
