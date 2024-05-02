package leetcode

import "testing"

func TestMaximumScore1442(t *testing.T) {
	s := "011101"
	if r := maxScore1442(s); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
	s = "00111"
	if r := maxScore1442(s); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
	s = "1111"
	if r := maxScore1442(s); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
