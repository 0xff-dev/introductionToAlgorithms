package leetcode

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	for end := 1; end <= 5; end++ {
		l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
		head := removeNthFromEnd(l, end)
		t.Logf("loop[%d]------", end)
		head.Dis()
	}

	t.Logf("single node[%d]", 1)
	singleNode := &ListNode{1, nil}
	head := removeNthFromEnd(singleNode, 1)
	head.Dis()
}
