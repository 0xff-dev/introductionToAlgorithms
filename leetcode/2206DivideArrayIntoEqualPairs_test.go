package leetcode

import "testing"

func TestDivideArray(t *testing.T) {
	n := []int{3, 2, 3, 2, 2, 2}
	if !divideArray(n) {
		t.Fatalf("expect true get false")
	}

	n = []int{1, 2, 3, 4}
	if divideArray(n) {
		t.Fatalf("expect false get true")
	}
}
