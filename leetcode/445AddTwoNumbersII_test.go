package leetcode

import "testing"

func TestGetListIntArray(t *testing.T) {
	l := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	array := make([]int, 0)
	getListIntArray(l, &array)
	t.Log(array)
}

func TestAddTwoNumbers1(t *testing.T) {
	l := &ListNode{7, &ListNode{2, &ListNode{4, &ListNode{3, nil}}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	head := addTwoNumbers1(l, l2)
	head.Dis()
	head = addTwoNumbers1(l2, l)
	head.Dis()
	head = addTwoNumbers1(nil, l2)
	head.Dis()
	head = addTwoNumbers1(l, nil)
	head.Dis()
	if head = addTwoNumbers1(nil, nil); head != nil {
		t.Fatalf("should return nil pointer")
	}

	l = &ListNode{1, nil}
	l2 = &ListNode{3, nil}
	head = addTwoNumbers1(l, l2)
	head.Dis()

}
