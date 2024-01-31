package leetcode

import "testing"

func TestFindTheLongestSubstring(t *testing.T) {
	s := "eleetminicoworoep"
	e := 13
	if r := findTheLongestSubstring(s); r != e {
		t.Fatalf("expect %d get %d", e, r)
	}
	s = "leetcodeisgreat"
	e = 5
	if r := findTheLongestSubstring(s); r != e {
		t.Fatalf("expect %d get %d", e, r)
	}

	s = "bcbcbc"
	e = 6
	if r := findTheLongestSubstring(s); r != e {
		t.Fatalf("expect %d get %d", e, r)
	}

}
