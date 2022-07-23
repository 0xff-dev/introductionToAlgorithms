package leetcode

import "testing"

func TestCanBeEqual(t *testing.T) {
	target, arr := []int{1, 2, 3, 4}, []int{2, 4, 1, 3}
	if !canBeEqual(target, arr) {
		t.Fatalf("expect true get false")
	}

	target, arr = []int{7}, []int{7}
	if !canBeEqual(target, arr) {
		t.Fatalf("expect true get false")
	}

	target, arr = []int{3, 7, 9}, []int{3, 7, 11}
	if canBeEqual(target, arr) {
		t.Fatalf("expect false get true")
	}
}