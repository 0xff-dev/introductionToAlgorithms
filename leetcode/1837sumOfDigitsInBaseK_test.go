package leetcode

import "testing"

func TestSumBase(t *testing.T) {
	n, k := 34, 6
	if r := sumBase(n, k); r != 9 {
		t.Fatalf("expect 9 get %d", r)
	}

	n, k = 10, 10
	if r := sumBase(n, k); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
