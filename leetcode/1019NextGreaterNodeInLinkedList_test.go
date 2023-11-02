package leetcode

import (
	"reflect"
	"testing"
)

func TestNextLargerNodes(t *testing.T) {
	l := &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 5}}}
	exp := []int{5, 5, 0}
	if r := nextLargerNodes(l); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	l = &ListNode{Val: 2, Next: &ListNode{Val: 7, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}}
	exp = []int{7, 0, 5, 5, 0}
	if r := nextLargerNodes(l); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
