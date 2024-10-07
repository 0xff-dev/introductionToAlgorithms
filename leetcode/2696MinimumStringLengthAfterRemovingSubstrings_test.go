package leetcode

import "testing"

func TestMinLength(t *testing.T) {
	s, exp := "ABFCACDB", 2
	if r := minLength(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s, exp = "ACBBD", 5
	if r := minLength(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
