package leetcode

import "testing"

func TestReverseWords3(t *testing.T) {
	s := "Let's take LeetCode contest"
	expect := "s'teL ekat edoCteeL tsetnoc"
	if r := reverseWords3(s); r != expect {
		t.Fatalf("input: %s, expect: %s, get %s", s, expect, r)
	}

	s = "God Ding"
	expect = "doG gniD"
	if r := reverseWords3(s); r != expect {
		t.Fatalf("input: %s, expect: %s, get %s", s, expect, r)
	}
}
