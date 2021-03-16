package leetcode

import "testing"

func TestCanJump(t *testing.T) {
	input := []int{2, 3, 1, 1, 4}
	if !canJump(input) {
		t.Fatalf("expect true get false")
	}

	input = []int{3, 2, 1, 0, 4}
	if canJump(input) {
		t.Fatalf("expect false get true")
	}

	input = []int{1}
	if !canJump(input) {
		t.Fatalf("expect true get false")
	}

}
