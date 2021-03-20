package leetcode

import "testing"

func TestDeleteDuplicate(t *testing.T) {
	list := &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}}
	head := deleteDuplicates(list)
	head.Dis()

	list = &ListNode{1, &ListNode{1, nil}}
	if head = deleteDuplicates(list); head != nil {
		t.Fatalf("should return nil")
	}

	list = &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, nil}}}}
	head = deleteDuplicates(list)
	head.Dis()

	list = &ListNode{1, &ListNode{2, &ListNode{2, nil}}}
	head = deleteDuplicates(list)
	head.Dis()

	list = &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}
	head = deleteDuplicates(list)
	head.Dis()

	list = &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, nil}}}}}}
	head = deleteDuplicates(list)
	head.Dis()

	list = &ListNode{1, &ListNode{1, &ListNode{2, nil}}}
	head = deleteDuplicates(list)
	head.Dis()

	list = &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, nil}}}}
	if head = deleteDuplicates(list); head != nil {
		t.Fatalf("should return nil")
	}
}
