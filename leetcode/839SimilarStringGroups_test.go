package leetcode

import "testing"

func TestNumSimilarGroups(t *testing.T) {
	s := []string{"tars", "rats", "arts", "star"}
	if r := numSimilarGroups(s); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	s = []string{"omv", "ovm"}
	if r := numSimilarGroups(s); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	s = []string{"abc", "abc"}
	if r := numSimilarGroups(s); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
