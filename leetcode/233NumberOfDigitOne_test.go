package leetcode

import "testing"

func TestCountDigitOne(t *testing.T) {
	if r := countDigitOne(100); r != 21 {
		t.Fatalf("expect 21 get %d", r)
	}

	if r := countDigitOne(21345); r != 18821 {
		t.Fatalf("expect 18821 get %d", r)
	}
}
