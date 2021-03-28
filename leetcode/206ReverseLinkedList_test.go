package leetcode

import "testing"

func TestReverseLinkedList(t *testing.T) {
	list := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	list.Dis()
	newHead := reverseList(list)
	newHead.Dis()

	list = &ListNode{1, nil}
	list.Dis()
	newHead = reverseList(list)
	newHead.Dis()

	list = &ListNode{1, &ListNode{2, nil}}
	list.Dis()
	newHead = reverseList(list)
	newHead.Dis()
}
