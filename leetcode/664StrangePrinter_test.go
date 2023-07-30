package leetcode

import "testing"

func TestStrangePrinter(t *testing.T) {
	s := "aaabbb"
	if r := strangePrinter(s); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	s = "aababab"
	if r := strangePrinter(s); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}
