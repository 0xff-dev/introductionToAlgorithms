package leetcode

import "testing"

func TestMergeKLists(t *testing.T) {
	lists := []*ListNode{
		&ListNode{1, &ListNode{4, &ListNode{5, nil}}},
		&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
		&ListNode{2, &ListNode{6, nil}},
	}
	newHead := mergeKLists(lists)
	newHead.Dis()

	lists = []*ListNode{
		nil,
	}

	if newHead = mergeKLists(lists); newHead != nil {
		t.Fatalf("expect nil")
	}

	lists = []*ListNode{
		{1, &ListNode{2, nil}},
	}
	mergeKLists(lists).Dis()
}
