package leetcode

import "testing"

func TestMinExtraChar(t *testing.T) {
	s := "leetscode"
	d := []string{"leet", "code", "leetcode"}
	if r := minExtraChar(s, d); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	s = "sayhelloworld"
	d = []string{"hello", "world"}
	if r := minExtraChar(s, d); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
