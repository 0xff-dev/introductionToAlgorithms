package leetcode

import "testing"

func TestRemoveElements(t *testing.T) {
	list := &ListNode{1, &ListNode{2, &ListNode{6, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}
	r := removeElements(list, 6)
	r.Dis()

	list = &ListNode{7, &ListNode{7, &ListNode{7, nil}}}
	r = removeElements(list, 7)
	if r != nil {
		r.Dis()
		t.Fatalf("expect nil")
	}

	list = &ListNode{7, nil}
	if r = removeElements(list, 7); r != nil {
		t.Fatalf("expect nil")
	}

	list = &ListNode{7, nil}
	if r = removeElements(list, 8); r != nil {
		r.Dis()
	}

}
