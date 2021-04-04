package leetcode

import "testing"

func TestValidatestackSequences(t *testing.T) {
	push, pop := []int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}
	if !validateStackSequences(push, pop) {
		t.Fatalf("should return true")
	}
}
