package leetcode

import (
	"reflect"
	"testing"
)

func TestReverseEvenLengthGroups(t *testing.T) {
	l := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 7,
								Next: &ListNode{
									Val: 3,
									Next: &ListNode{
										Val:  8,
										Next: &ListNode{Val: 4},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	exp := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 8,
									Next: &ListNode{
										Val:  3,
										Next: &ListNode{Val: 7},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	l1 := reverseEvenLengthGroups(l)
	exp.Dis()
	t.Log("\n")
	l1.Dis()
	if !reflect.DeepEqual(l1, exp) {
		t.Fatalf("expect %v get %v", exp, l)
	}
}
