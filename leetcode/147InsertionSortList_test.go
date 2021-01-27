package leetcode

import "testing"

func TestInsertionSortList(t *testing.T) {
	l := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	head := insertionSortList(l)
	head.Dis()
	l = &ListNode{-1, &ListNode{5, &ListNode{3, &ListNode{4, &ListNode{0, nil}}}}}
	head = insertionSortList(l)
	head.Dis()

	if head = insertionSortList(nil); head != nil {
		t.Fatalf("shoule return nil")
	}
	l = &ListNode{7, &ListNode{1, nil}}
	head = insertionSortList(l)
	head.Dis()

}
