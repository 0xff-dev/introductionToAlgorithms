package leetcode

import "testing"

func TestReverseBits(t *testing.T) {
	var (
		n, exp uint32
	)
	n = 0b00000010100101000001111010011100
	exp = 964176192
	if r := reverseBits(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	n = 0b11111111111111111111111111111101
	exp = 3221225471
	if r := reverseBits(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
