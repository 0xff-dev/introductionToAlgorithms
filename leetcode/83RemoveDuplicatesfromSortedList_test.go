package leetcode

import "testing"

func TestDeleteDuplicates1(t *testing.T) {
	list := &ListNode{1, &ListNode{1, &ListNode{2, nil}}}
	head := deleteDuplicates1(list)
	head.Dis()

	list = &ListNode{1, &ListNode{1, &ListNode{1, nil}}}
	head = deleteDuplicates1(list)
	head.Dis()

	list = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	head = deleteDuplicates1(list)
	head.Dis()

	list = &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}
	head = deleteDuplicates1(list)
	head.Dis()
}
