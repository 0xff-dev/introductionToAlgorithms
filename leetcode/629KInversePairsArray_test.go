package leetcode

import "testing"

func TestKInversePairs(t *testing.T) {
	n, k := 3, 0
	if r := kInversePairs(n, k); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	n, k = 4, 2
	if r := kInversePairs(n, k); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
	n, k = 38, 36
	if r := kInversePairs(n, k); r != 506617666 {
		t.Fatalf("expect 506617666 get %d", r)
	}
}
