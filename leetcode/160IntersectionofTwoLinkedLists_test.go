package leetcode

import "testing"

func TestGetIntersectionNode(t *testing.T) {
	common := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}
	ha := &ListNode{Val: 1, Next: &ListNode{Val: 9, Next: common}}
	hb := &ListNode{Val: 3, Next: common}
	if r := getIntersectionNode(ha, hb); r != common {
		t.Fatalf("expect %v get %v", common, r)
	}

	common = &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}
	ha = &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: common}}
	hb = &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: common}}}
	if r := getIntersectionNode(ha, hb); r != common {
		t.Fatalf("expect %v get %v", common, r)
	}

	ha = &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	hb = &ListNode{Val: 1, Next: &ListNode{Val: 5}}
	if r := getIntersectionNode(ha, hb); r != nil {
		t.Fatalf("expect nil get %v", r)
	}
}
