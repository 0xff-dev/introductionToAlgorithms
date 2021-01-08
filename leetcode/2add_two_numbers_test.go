package leetcode

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	var l1, l2 *ListNode
	if r := addTwoNumbers(l1, l2); r != nil {
		t.Fatalf("nil list, should return nil")
	}
	l1 = &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 = &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	resultArray := []int{7, 0, 8}
	r := addTwoNumbers(l1, l2)
	if r == nil {
		t.Fatalf("Both L1 and L2 are not empty")
	}
	idx := 0
	for walker := r; walker != nil; walker = walker.Next {
		if walker.Val != resultArray[idx] {
			t.Fatalf("[%d] don't equal [%d]", walker.Val, resultArray[idx])
		}
		idx++
	}

	l1, l2 = &ListNode{0, nil}, &ListNode{0, nil}
	resultArray = []int{0}
	r = addTwoNumbers(l1, l2)
	if r == nil {
		t.Fatalf("Both L1 and L2 are not empty")
	}
	idx = 0
	for walker := r; walker != nil; walker = walker.Next {
		if walker.Val != resultArray[idx] {
			t.Fatalf("[%d] don't equal [%d]", walker.Val, resultArray[idx])
		}
		idx++
	}

	l1 = &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	l2 = &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	resultArray = []int{8, 9, 9, 9, 0, 0, 0, 1}
	r = addTwoNumbers(l1, l2)
	if r == nil {
		t.Fatalf("Both L1 and L2 are not empty")
	}
	idx = 0
	for walker := r; walker != nil; walker = walker.Next {
		if walker.Val != resultArray[idx] {
			t.Fatalf("[%d] don't equal [%d]", walker.Val, resultArray[idx])
		}
		idx++
	}

	l1 = &ListNode{2, &ListNode{4, &ListNode{9, nil}}}
	l2 = &ListNode{5, &ListNode{6, &ListNode{4, &ListNode{9, nil}}}}
	r = addTwoNumbers(l1, l2)
	resultArray = []int{7, 0, 4, 0, 1}
	if r == nil {
		t.Fatalf("Both L1 and L2 are not empty")
	}
	idx = 0
	for walker := r; walker != nil; walker = walker.Next {
		if walker.Val != resultArray[idx] {
			t.Fatalf("[%d] don't equal [%d]", walker.Val, resultArray[idx])
		}
		idx++
	}

}
