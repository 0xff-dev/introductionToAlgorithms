package leetcode

import (
	"reflect"
	"testing"
)

func TestSwapNodes(t *testing.T) {
	h := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	h2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	h1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5}}}}}
	if r := swapNodes(h, 2); !reflect.DeepEqual(r, h1) {
		t.Fatalf("expect %v get %v", h1, r)
	}
	if r := swapNodes1(h2, 2); !reflect.DeepEqual(r, h1) {
		t.Fatalf("expect %v get %v", h1, r)
	}
}
