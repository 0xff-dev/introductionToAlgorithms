package leetcode

import "testing"

func TestWordBreak140(t *testing.T) {
	s := "catsanddog"
	dict := []string{"cat", "cats", "and", "sand", "dog"}
	r := wordBreak140(s, dict)
	t.Logf("--- ans: %+v\n", r)
}
