package leetcode

import (
	"reflect"
	"testing"
)

func TestModifiedList(t *testing.T) {
	n, h, e := []int{1, 2, 3}, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, &ListNode{4, &ListNode{5, nil}}
	if r := modifiedList(n, h); !reflect.DeepEqual(e, r) {
		t.Fatalf("expect %v get %v", e, r)
	}
	n, h, e = []int{1}, &ListNode{1, &ListNode{2, &ListNode{1, &ListNode{2, &ListNode{1, &ListNode{2, nil}}}}}}, &ListNode{2, &ListNode{2, &ListNode{2, nil}}}
	if r := modifiedList(n, h); !reflect.DeepEqual(e, r) {
		t.Fatalf("expect %v get %v", e, r)
	}
	n, h, e = []int{5}, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	if r := modifiedList(n, h); !reflect.DeepEqual(e, r) {
		t.Fatalf("expect %v get %v", e, r)
	}
}
