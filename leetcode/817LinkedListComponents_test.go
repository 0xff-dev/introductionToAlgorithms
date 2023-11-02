package leetcode

import "testing"

func TestNumComponents(t *testing.T) {
	head := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}
	nums := []int{0, 1, 3}
	if r := numComponents(head, nums); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	head = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{Val: 4},
				},
			},
		},
	}
	nums = []int{0, 3, 1, 4}
	if r := numComponents(head, nums); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
