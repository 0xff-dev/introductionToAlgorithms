package leetcode

import (
	"reflect"
	"testing"
)

func TestMergeNodes(t *testing.T) {
	l := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 2,
								Next: &ListNode{
									Val: 0,
								},
							},
						},
					},
				},
			},
		},
	}
	exp := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 11,
		},
	}
	if r := mergeNodes(l); !reflect.DeepEqual(exp, l) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
