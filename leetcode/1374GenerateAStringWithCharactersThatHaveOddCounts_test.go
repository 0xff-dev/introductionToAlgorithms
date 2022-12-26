package leetcode

import "testing"

func TestGenerateTheString(t *testing.T) {
	n := 4
	exp := "abbb"
	if r := generateTheString(n); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
