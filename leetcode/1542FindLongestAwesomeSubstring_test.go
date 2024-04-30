package leetcode

import "testing"

func TestLongestAwesome(t *testing.T) {
	s := "3242415"
	exp := 5
	if r := longestAwesome(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s = "12345678"
	exp = 1
	if r := longestAwesome(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s = "213123"
	exp = 6
	if r := longestAwesome(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
