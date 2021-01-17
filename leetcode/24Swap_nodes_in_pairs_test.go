package leetcode

import "testing"

func TestSwapPairs(t *testing.T) {
	var l *ListNode
	t.Log("nil link list")
	head := swapPairs(l)
	head.Dis()

	t.Logf("single link node")
	l = &ListNode{1, nil}
	head = swapPairs(l)
	head.Dis()

	t.Logf("two nodes")
	l = &ListNode{1, &ListNode{2, nil}}
	head = swapPairs(l)
	head.Dis()

	t.Logf("three nodes")
	l = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	head = swapPairs(l)
	head.Dis()

	t.Logf("four nodes")
	l = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	head = swapPairs(l)
	head.Dis()

	t.Logf("five nodes")
	l = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head = swapPairs(l)
	head.Dis()

	t.Logf("six nodes")
	l = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	head = swapPairs(l)
	head.Dis()
}
