package leetcode

import "testing"

func TestMaximumOddBinaryNumber(t *testing.T) {
	s := "010"
	exp := "001"
	if r := maximumOddBinaryNumber(s); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}

	s = "0101"
	exp = "1001"
	if r := maximumOddBinaryNumber(s); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
