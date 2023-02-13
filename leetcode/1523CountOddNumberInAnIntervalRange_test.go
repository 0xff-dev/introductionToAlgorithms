package leetcode

import "testing"

func TestCountOdds(t *testing.T) {
	low, high := 3, 7
	if r := countOdds(low, high); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	low, high = 8, 10
	if r := countOdds(low, high); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
