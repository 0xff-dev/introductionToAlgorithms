package leetcode

import "testing"

func TestGetDecimalValue(t *testing.T) {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
			},
		},
	}

	if r := getDecimalValue(l); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	l = &ListNode{
		Val: 0,
	}
	if r := getDecimalValue(l); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	l = &ListNode{
		Val: 1,
	}
	if r := getDecimalValue(l); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
