package leetcode

import "testing"

func TestValiateBinaryTreeNodes(t *testing.T) {
	n := 4
	left := []int{1, -1, 3, -1}
	right := []int{2, -1, -1, -1}
	if !validateBinaryTreeNodes(n, left, right) {
		t.Fatalf("expect true get false")
	}

	n = 4
	left = []int{1, -1, 3, -1}
	right = []int{2, 3, -1, -1}
	if validateBinaryTreeNodes(n, left, right) {
		t.Fatalf("expect false get true")
	}

	n = 2
	left = []int{1, 0}
	right = []int{-1, -1}
	if validateBinaryTreeNodes(n, left, right) {
		t.Fatalf("expect false get true")
	}
}
