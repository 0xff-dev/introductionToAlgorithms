package leetcode

import "testing"

func TestWays(t *testing.T) {
	pizza := []string{"A..", "AAA", "..."}
	k := 3
	if r := ways(pizza, k); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	pizza = []string{"A..", "AA.", "..."}
	k = 3
	if r := ways(pizza, k); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
