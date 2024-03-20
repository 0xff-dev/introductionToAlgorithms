package leetcode

import (
	"reflect"
	"testing"
)

func TestMergeInBetween(t *testing.T) {
	l1 := &ListNode{
		Val: 10,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 13,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  9,
						Next: &ListNode{Val: 5},
					},
				},
			},
		},
	}
	l2 := &ListNode{Val: 1000000, Next: &ListNode{Val: 1000001, Next: &ListNode{Val: 1000002}}}
	exp := &ListNode{
		Val: 10,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 13,
				Next: &ListNode{
					Val: 1000000,
					Next: &ListNode{
						Val:  1000001,
						Next: &ListNode{Val: 1000002, Next: &ListNode{Val: 5}},
					},
				},
			},
		},
	}
	a, b := 3, 4
	if r := mergeInBetween(l1, a, b, l2); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	l1 = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}},
					},
				},
			},
		},
	}
	l2 = &ListNode{Val: 1000000, Next: &ListNode{Val: 1000001, Next: &ListNode{Val: 1000002, Next: &ListNode{Val: 1000003, Next: &ListNode{Val: 1000004}}}}}
	exp = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1000000,
				Next: &ListNode{
					Val: 1000001,
					Next: &ListNode{
						Val:  1000002,
						Next: &ListNode{Val: 1000003, Next: &ListNode{Val: 1000004, Next: &ListNode{Val: 6}}},
					},
				},
			},
		},
	}
	a, b = 2, 5
	if r := mergeInBetween(l1, a, b, l2); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
