package leetcode

import "testing"

func TestFindReplaceString(t *testing.T) {
	s := "abcd"
	i := []int{0, 2}
	sources := []string{"a", "cd"}
	targets := []string{"eee", "ffff"}
	exp := "eeebffff"
	if r := findReplaceString(s, i, sources, targets); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}

	sources = []string{"ab", "ec"}

	exp = "eeecd"
	if r := findReplaceString(s, i, sources, targets); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}

	i = []int{0, 3, 2, 1}
	sources = []string{"a", "d", "c", "b"}
	targets = []string{"d", "a", "b", "c"}
	exp = "dcba"
	if r := findReplaceString(s, i, sources, targets); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
