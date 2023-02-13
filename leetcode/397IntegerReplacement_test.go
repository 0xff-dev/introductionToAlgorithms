package leetcode

import "testing"

func TestIntegerReplacement(t *testing.T) {
	if r := integerReplacement(8); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	if r := integerReplacement(7); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	if r := integerReplacement(4); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
