package leetcode

import "testing"

func TestPairSum(t *testing.T) {
	l := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  2,
				Next: &ListNode{Val: 1},
			},
		},
	}
	if r := pairSum(l); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	l = &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  2,
				Next: &ListNode{Val: 3},
			},
		},
	}
	if r := pairSum(l); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
}
