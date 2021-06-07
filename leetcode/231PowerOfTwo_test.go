package leetcode

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	n := 4
	if !isPowerOfTwo(n) {
		t.Fatalf("expect true get false")
	}

	n = 5
	if isPowerOfTwo(n) {
		t.Fatalf("expect false get true")
	}
}
