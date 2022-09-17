package leetcode

import "testing"

func TestCanMakeArithmeticProgression(t *testing.T) {
	n := []int{3, 5, 1}
	if !canMakeArithmeticProgression(n) {
		t.Fatalf("expect true get false")
	}

	n = []int{1, 2, 4}
	if canMakeArithmeticProgression(n) {
		t.Fatalf("expect false get true")
	}
}
