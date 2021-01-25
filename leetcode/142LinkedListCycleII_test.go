package leetcode

import "testing"

func TestDetectCycle(t *testing.T) {
	one := &ListNode{1, nil}
	two := &ListNode{2, one}
	one.Next = two
	cyclePoint := detectCycle(one)
	if cyclePoint == nil {
		t.Fatalf("has cycle, point value is: %d", one.Val)
	}

	t.Logf("find cycle point value is: %d", cyclePoint.Val)

	one = &ListNode{3, nil}
	two = &ListNode{2, nil}
	three := &ListNode{0, &ListNode{4, two}}
	two.Next = three
	one.Next = two
	cyclePoint = detectCycle(one)
	if cyclePoint == nil {
		t.Fatalf("has cycle, point value is: %d", two.Val)
	}
	t.Logf("find cycle point value is: %d", cyclePoint.Val)

	one = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	if cyclePoint = detectCycle(one); cyclePoint != nil {
		t.Fatalf("list 1234 don't have cycle.")
	}
}
