package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
	list := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	if !isPalindrome(list) {
		t.Fatalf("expect true get false")
	}

	list = &ListNode{1, &ListNode{2, nil}}
	if isPalindrome(list) {
		t.Fatalf("expect false get true")
	}

	list = &ListNode{1, nil}
	if !isPalindrome(list) {
		t.Fatalf("expect true get false")
	}
}
