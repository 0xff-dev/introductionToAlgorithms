package leetcode

import "testing"

func TestTitleToNumber(t *testing.T) {
	input := "A"
	if r := titleToNumber(input); r != 1 {
		t.Fatalf("exepct 1 get %d", r)
	}

	input = "D"
	if r := titleToNumber(input); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	input = "AA"
	if r := titleToNumber(input); r != 27 {
		t.Fatalf("expect 27 get %d", r)
	}

	input = "Z"
	if r := titleToNumber(input); r != 26 {
		t.Fatalf("expect 26 get %d", r)
	}

	input = "FXSHRXW"
	if r := titleToNumber(input); r != 2147483647 {
		t.Fatalf("expect 2147483647 get %d", r)
	}

	input = "ZY"
	if r := titleToNumber(input); r != 701 {
		t.Fatalf("expect 701 get %d", r)
	}
}
