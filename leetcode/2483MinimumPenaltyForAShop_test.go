package leetcode

import "testing"

func TestBestClosingTime(t *testing.T) {
	if r := bestClosingTime("YYNY"); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	if r := bestClosingTime("NNNNN"); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	if r := bestClosingTime("YYYY"); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}
