package leetcode

import "testing"

func TestCountGoodNumbers(t *testing.T) {
	n := int64(1)
	exp := 5
	if r := countGoodNumbers(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n = 4
	exp = 400
	if r := countGoodNumbers(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n = 50
	exp = 564908303
	if r := countGoodNumbers(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
