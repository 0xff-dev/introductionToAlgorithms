package leetcode

import "testing"

func TestCheckRecord(t *testing.T) {
	n := 2
	exp := 8
	if r := checkRecord(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n = 1
	exp = 3
	if r := checkRecord(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n = 10101
	exp = 183236316
	if r := checkRecord(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
