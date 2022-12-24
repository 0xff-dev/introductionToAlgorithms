package leetcode

import "testing"

func TestNumTilings(t *testing.T) {
	n := 3
	if r := numTilings(n); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	n = 9
	if r := numTilings(n); r != 569 {
		t.Fatalf("expect 569 get %d", r)
	}
}
