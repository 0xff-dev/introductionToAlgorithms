package leetcode

import "testing"

func TestNumDecoding(t *testing.T) {
	s := "12"
	if r := numDecodings(s); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	s = "226"
	if r := numDecodings(s); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	s = "06"
	if r := numDecodings(s); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
