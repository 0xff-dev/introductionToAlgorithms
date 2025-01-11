package leetcode

import "testing"

func TestCanConstruct1400(t *testing.T) {
	s, k, exp := "annabelle", 2, true
	if r := canConstruct1400(s, k); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	s, k, exp = "leetcode", 3, false
	if r := canConstruct1400(s, k); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	s, k, exp = "true", 4, true
	if r := canConstruct1400(s, k); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
