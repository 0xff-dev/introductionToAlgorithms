package leetcode

import "testing"

func TestMinFlipMonoIncr(t *testing.T) {
	s := "00110"
	if r := minFlipsMonoIncr(s); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	s = "010110"
	if r := minFlipsMonoIncr(s); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	s = "00011000"
	if r := minFlipsMonoIncr(s); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
