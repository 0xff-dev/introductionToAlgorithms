package leetcode

import "testing"

func TestSortList(t *testing.T) {
	list := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	head := sortList(list)
	head.Dis()

	list = &ListNode{1, &ListNode{2, nil}}
	head = sortList(list)
	head.Dis()

	list = &ListNode{1, nil}
	head = sortList(list)
	head.Dis()

	list = &ListNode{18, &ListNode{1, &ListNode{1324, &ListNode{89, &ListNode{135, &ListNode{10, nil}}}}}}
	head = sortList(list)
	head.Dis()

	if head = sortList(nil); head != nil {
		t.Fail()
	}
}
