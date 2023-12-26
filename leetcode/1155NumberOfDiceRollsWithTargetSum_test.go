package leetcode

import "testing"

func TestNumRollsToTarget(t *testing.T) {
	n, k, target := 1, 6, 3
	if r := numRollsToTarget(n, k, target); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	n, k, target = 2, 6, 7
	if r := numRollsToTarget(n, k, target); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
	n, k, target = 30, 30, 500
	if r := numRollsToTarget(n, k, target); r != 222616187 {
		t.Fatalf("expect 222616187 get %d", r)
	}
}
