package leetcode

import "testing"

func TestRotateRight(t *testing.T) {
	var l *ListNode
	t.Log("nil link list")
	head := rotateRight(l, 1)
	head.Dis()

	t.Logf("single link node")
	l = &ListNode{1, nil}
	head = rotateRight(l, 2)
	head.Dis()

	t.Logf("two nodes")
	for idx := 0; idx <= 2; idx++ {
		l = &ListNode{1, &ListNode{2, nil}}
		t.Logf("[%d]\n", idx)
		head = rotateRight(l, idx)
		head.Dis()
	}

	t.Logf("three nodes")
	for idx := 0; idx <= 3; idx++ {
		l = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
		t.Logf("[%d]\n", idx)
		head = rotateRight(l, idx)
		head.Dis()
	}

	t.Logf("four nodes")
	for idx := 0; idx <= 4; idx++ {
		l = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
		t.Logf("[%d]\n", idx)
		head = rotateRight(l, idx)
		head.Dis()
	}

	t.Logf("five nodes")
	for idx := 0; idx <= 5; idx++ {
		l = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
		t.Logf("[%d]\n", idx)
		head = rotateRight(l, idx)
		head.Dis()
	}

	t.Logf("six nodes")
	for idx := 0; idx <= 10; idx++ {
		l = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
		t.Logf("[%d]\n", idx)
		head = rotateRight(l, idx)
		head.Dis()
	}

}
