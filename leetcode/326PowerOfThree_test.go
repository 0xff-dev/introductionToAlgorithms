package leetcode

import "testing"

func TestIsPowerOfThree(t *testing.T) {
	// 1, 3, 9, 27
	n := 1
	if !isPowerOfThree(n) {
		t.Fatalf("expect true get false")
	}
	if !isPowerOfThree(3) {
		t.Fatalf("expect true get false")
	}
	if isPowerOfThree(5) {
		t.Fatalf("expect false get true")
	}
}
