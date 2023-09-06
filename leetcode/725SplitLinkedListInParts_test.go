package leetcode

import (
	"reflect"
	"testing"
)

func TestSplitListToParts(t *testing.T) {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 3}
	a.Next = b
	b.Next = c
	exp := []*ListNode{a, b, c, nil, nil}
	if ans := splitListToParts(a, 5); !reflect.DeepEqual(exp, ans) {
		t.Fatalf("expect %v get %v", exp, ans)
	}

	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 7,
								Next: &ListNode{
									Val:  8,
									Next: &ListNode{Val: 9, Next: &ListNode{Val: 10}},
								},
							},
						},
					},
				},
			},
		},
	}
	exp = []*ListNode{&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}, &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 7,
			},
		},
	}, &ListNode{
		Val: 8,
		Next: &ListNode{
			Val:  9,
			Next: &ListNode{Val: 10},
		},
	}}
	if ans := splitListToParts(l, 3); !reflect.DeepEqual(exp, ans) {
		t.Fatalf("expect %v get %v", exp, ans)
	}
}
