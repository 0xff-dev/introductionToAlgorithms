package leetcode

import "testing"

func TestLongestIdealString(t *testing.T) {
	s := "acfgbd"
	k := 2
	exp := 4
	if r := longestIdealString(s, k); r != exp {
		t.Fatalf("%s expect %d get %d", s, exp, r)
	}

	s = "abcd"
	k = 3
	exp = 4
	if r := longestIdealString(s, k); r != exp {
		t.Fatalf("%s expect %d get %d", s, exp, r)
	}

	s = "pvjcci"
	k = 4
	exp = 2
	if r := longestIdealString(s, k); r != exp {
		t.Fatalf("%s expect %d get %d", s, exp, r)
	}
}
