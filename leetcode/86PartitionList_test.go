package leetcode

import "testing"

func TestPartition(t *testing.T) {
	list := &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}}
	head := partition(list, 3)
	head.Dis()

	list = &ListNode{2, &ListNode{1, nil}}
	head = partition(list, 2)
	head.Dis()

	list = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	head = partition(list, 4)
	head.Dis()

	head = partition(list, 0)
	head.Dis()
}
