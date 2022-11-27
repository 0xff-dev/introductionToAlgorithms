package leetcode

import "testing"

func TestCountBinarySubstrings(t *testing.T) {
	s := "00110011"
	if r := countBinarySubstrings(s); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	s = "10101"
	if r := countBinarySubstrings(s); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	s = "000111"
	if r := countBinarySubstrings(s); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
