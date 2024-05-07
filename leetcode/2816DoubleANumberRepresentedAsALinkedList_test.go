package leetcode

import (
	"reflect"
	"testing"
)

func TestDoubleIt(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9}}}
	exp := &ListNode{Val: 3, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8}}}
	if r := doubleIt(head); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	head = &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}
	exp = &ListNode{Val: 1, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 8}}}}
	if r := doubleIt(head); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
