package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveNodes(t *testing.T) {
	head := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 13,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{Val: 8},
				},
			},
		},
	}
	exp := &ListNode{Val: 13, Next: &ListNode{Val: 8}}
	if r := removeNodes(head); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 1},
			},
		},
	}
	exp = head
	if r := removeNodes(head); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

}
