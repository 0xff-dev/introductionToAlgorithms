package leetcode

import (
	"reflect"
	"testing"
)

func TestNodesBetweenCriticalPoints(t *testing.T) {
	l := &ListNode{Val: 3, Next: &ListNode{Val: 1}}
	exp := []int{-1, -1}
	if r := nodesBetweenCriticalPoints(l); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	l = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 2,
							},
						},
					},
				},
			},
		},
	}
	exp = []int{1, 3}
	if r := nodesBetweenCriticalPoints(l); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	l = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 2,
								Next: &ListNode{
									Val:  2,
									Next: &ListNode{Val: 7},
								},
							},
						},
					},
				},
			},
		},
	}
	exp = []int{3, 3}
	if r := nodesBetweenCriticalPoints(l); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
