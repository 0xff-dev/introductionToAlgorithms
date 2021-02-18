package leetcode

import "testing"

func TestFindSubstring(t *testing.T) {
	input, words := "barfoothefoobarman", []string{"foo", "bar"}
	r := findSubstring(input, words)
	t.Logf("%v", r)

	input, words = "wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}
	r = findSubstring(input, words)
	t.Logf("%v", r)

	input, words = "barfoofoobarthefoobarman", []string{"bar", "foo", "the"}
	r = findSubstring(input, words)
	t.Logf("%v", r)
}
