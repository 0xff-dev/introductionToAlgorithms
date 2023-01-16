package leetcode

import "testing"

func TestSmalletEquivalentString(t *testing.T) {
	s1, s2, baseStr := "parker", "morris", "parser"
	if r := smallestEquivalentString(s1, s2, baseStr); r != "makkek" {
		t.Fatalf("expect 'makkek' get %s", r)
	}

	s1, s2, baseStr = "hello", "world", "hold"
	if r := smallestEquivalentString(s1, s2, baseStr); r != "hdld" {
		t.Fatalf("expect 'hdld' get %s", r)
	}

	s1, s2, baseStr = "leetcode", "programs", "sourcecode"
	if r := smallestEquivalentString(s1, s2, baseStr); r != "aauaaaaada" {
		t.Fatalf("expect 'aauaaaaada' get %s", r)
	}
}
