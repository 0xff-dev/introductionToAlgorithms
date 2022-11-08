package leetcode

import "testing"

func TestMakeGood(t *testing.T) {
	s := "leEeetcode"
	if r := makeGood(s); r != "leetcode" {
		t.Fatalf("expect %s get %s", "leetcode", r)
	}

	s = "abBAcC"
	if r := makeGood(s); r != "" {
		t.Fatalf("expect '' get %s", r)
	}

	s = "s"
	if r := makeGood(s); r != "s" {
		t.Fatalf("expect s get %s", r)
	}
}
