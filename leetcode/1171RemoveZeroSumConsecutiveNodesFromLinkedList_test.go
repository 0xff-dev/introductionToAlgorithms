package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveZerosumSublists(t *testing.T) {
	// 1,2,-3,3,1]
	// 3,1   1,2,1
	l := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: -3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 1}}}}}
	exp := &ListNode{Val: 3, Next: &ListNode{Val: 1}}
	if r := removeZeroSumSublists(l); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	// 1,2,3,-3,4
	// 1, 2, 4
	l = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: -3, Next: &ListNode{Val: 4}}}}}
	exp = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	if r := removeZeroSumSublists(l); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	// 1,2,3,-3,-2
	// 1
	l = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: -3, Next: &ListNode{Val: -2}}}}}
	exp = &ListNode{Val: 1}
	if r := removeZeroSumSublists(l); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	// 1,3,2,-3,-2,5,5,-5,1
	// 151
	l = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: -3, Next: &ListNode{Val: -2, Next: &ListNode{
		Val: 5, Next: &ListNode{Val: 5, Next: &ListNode{Val: -5, Next: &ListNode{Val: 1}}},
	}}}}}}
	exp = &ListNode{Val: 1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1}}}
	if r := removeZeroSumSublists(l); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
