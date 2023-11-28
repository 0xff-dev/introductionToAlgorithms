package leetcode

import "testing"

func TestNumberOfWays(t *testing.T) {
	if r := numberOfWays("SSPPSPS"); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	if r := numberOfWays("PPSPSP"); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	if r := numberOfWays("S"); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
