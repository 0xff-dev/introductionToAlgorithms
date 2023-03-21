package leetcode

import "testing"

func TestZeroFilledSubarray(t *testing.T) {
	n := []int{1, 3, 0, 0, 2, 0, 0, 4}
	if r := zeroFilledSubarray(n); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
	n = []int{0, 0, 0, 2, 0, 0}
	if r := zeroFilledSubarray(n); r != 9 {
		t.Fatalf("expect 9 get %d", r)
	}

	n = []int{2, 10, 2019}
	if r := zeroFilledSubarray(n); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
