package leetcode

import "testing"

func TestContructor745(t *testing.T) {
	words := []string{"apple"}
	w := Constructor745(words)
	if r := w.F("a", "e"); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	words = []string{"cabaabaaaa", "ccbcababac", "bacaabccba", "bcbbcbacaa", "abcaccbcaa", "accabaccaa", "cabcbbbcca", "ababccabcb", "caccbbcbab", "bccbacbcba"}
	w = Constructor745(words)
	_ans := []int{9, 4, 5, 0, 8, 1, 2, 5, 3, 1}
	for idx, sp := range [][]string{
		{"bccbacbcba", "a"},
		{"ab", "abcaccbcaa"},
		{"a", "aa"},
		{"cabaaba", "abaaaa"},
		{"cacc", "accbbcbab"},
		{"ccbcab", "bac"},
		{"bac", "cba"},
		{"ac", "accabaccaa"},
		{"bcbb", "aa"},
		{"ccbca", "cbcababac"}} {
		if r := w.F(sp[0], sp[1]); r != _ans[idx] {
			t.Fatalf("prefix: %s, suffix: %s, expect %d get %d", sp[0], sp[1], _ans[idx], r)
		}
	}

	words = []string{
		"dcdddcdccc", "ddddccdddc", "dccccccdcd", "cdcddccddc", "dccddccdcc", "cdccddcddd", "ccddccccdc", "dcccccddcd", "cccdcdccdd", "dcddcdcdcd", "dcdcddddcd", "ccccddcdcc", "dcccdddcdc", "cdccdccddd", "cddcccddcc", "cdddcccccc", "dcdddcddcc", "ccddcdddcd", "dcddcdcccd", "cccdccdddd", "ccdcdcddcc", "ccccddddcd", "cdcdccdccd", "ccdcddcccc", "ddddddcddd", "ddccdccddc", "cdcddcccdd", "dddcdddccc", "cddddcdccd", "dcdcddcdcd", "cdcccccddc", "cddddddccd", "cdcdcdcddc", "dccccccddd", "ddccdccccc", "ccddcdcdcd", "ddddcccdcc", "ccdcdccddc", "cccdcddccd", "cccdcccddd", "dccddcdddc", "ddddcddddc", "cddcddcdcd", "dddccdddcc", "dddcccdddc", "dcccdcdddc", "cccccccccd", "cdcdcccddd", "cccddccddd", "ccccdddccc", "ccddcdcddc", "dccdccddcc", "cddcccdccc", "cdcccdcdcc", "cccddddddd", "dccccdcddc", "ddddccdcdd", "cdddccccdd", "cddcdcdcdd", "dccdccdddc", "dcddcdddcd", "cccdddddcd", "ddcdcdddcc", "ddccccddcd", "ccdccdcddc", "ccccccdccd", "cdcdddcddd", "ddcdccddcc", "dccdcdccdc", "ddcdcccccc", "ddccccdddd", "cdcccddddd", "dcdddddddc", "ccdccccddd", "cdccdddddc", "cddcccdcdd", "ddcddddcdc", "dcddccdccc", "ccddcdcccc", "dddddddcdc", "ccdddddddc", "cdcdccddcc", "dcdcccccdc", "cdcdcdcddd", "ddccccdccc", "ddddccdccd", "dcdcccddcd", "cdcdddcddc", "ccccddccdc", "dcdccccddd", "dccddcccdd", "cccccdccdc", "ccddcdcccd", "dccdcdcdcc", "cdddddccdd", "dcccddcddd", "ccdcdccdcc", "ddcccdcddd", "dcddccccdc", "dddcdcccdc", "dcdcddccdd", "dccddcdcdc", "dcccdcddcc", "ddcdcddcdd", "dddcccccdd", "cdcdcdcdcc", "ddccccddcc", "ddcdcccdcd", "ccccdcdddd", "ccccdcdcdc", "cddddddddd", "cccdcdddcc", "dccccdcccd", "ccccddcccd", "cdddcdcccd", "cddcddcddc", "cdddcddcdd", "ddcddccddd", "ccccdccdcd", "ddcddcdcdc", "dddccddccd", "ccddccdcdd", "ddddccddcc", "ddddcddddd", "cddcdcddcc", "cddccccdcd", "ddcdcddcdc", "cccccdccdd", "cddcccddcd", "dddccdcddc", "ddcddcdccd", "dcdccddcdd", "dcdccdcccc", "cdccdcdccc", "ddcdccdcdc", "dccdddccdc", "ddddcdccdd", "dcddcdddcc", "ccdccddddd", "dddcdcdddc", "dcdccdcdcc", "dccddcddcd", "cdcddcdccc", "dcccccdcdc", "ddccddcddd", "cdddccddcd", "ccdddccddd", "ccdcdddddd", "cccdccccdc", "cddcdcddcd", "cdccddccdc", "cdcccddccd", "dcddcdcccc", "cccddcccdd", "dcdcccdddc", "dcdcdcccdd", "dccdcddcdc", "cccdcdcddd", "ddcdddddcc", "dccdcdccdd", "cddcdddddd", "cccccccddd", "dddddddccd", "ccdccdddcd", "ccdcdcccdd", "dccdcdddcd", "dcdddcdcdd", "dcddcccddc", "dcccddddcd", "ddcdccdddd", "dcdcdddddc", "ddcddcdddd", "ccdddddccd", "dcccdcccdc", "dddcccdcdc", "ddcdddcdcd", "dccccddccc", "ccccddcddc", "dccccddcdd", "ddcddcccdd", "cdcdcccccc", "cdcddddcdd", "ccccdcdcdd", "cddccccddd", "cdcdcccdcd", "cdccdddccc", "ccdcdccccd", "cdddccddcc", "ddcdddccdd", "dcdccdcddc", "dcddcddcdc", "ccccddddcc", "dcdcdcdccc", "ccdcdcdccd", "dcdcdcccdc", "dddcddcddd", "dccccdddcd", "cdccddcdcd", "ccccccccdc", "ddcdccdccd", "dddddcdccc", "ccdccddccc", "dccdcccdcd", "ddcddccccc", "dddcdddcdd", "ddcdddcddc", "cdddcdcdcd", "ccdddddcdc", "ddccdddddc", "dcccddcdcd", "cddddcdddc", "ccdddccccc", "dddcdccdcc", "cdccccccdc", "ddccdcdddc", "ccdcccdcdd", "ccddcccddd", "ccdcddccdc", "dccdcdcccc", "cdddddcccd", "dccddddddd", "ddccdcdccc", "cdcccdcddd", "dcddcccddd", "dccdccccdc", "cddcddddcd", "cccddcddcd", "cccddcccdc", "ccccdcccdc", "ddccdddccd", "ddddddddcc", "cddcddcccd", "dccccdddcc", "cccddcdccd", "dcdddcdddd", "dddcccddcd", "cdccccddcd", "ddcccdcdcd", "cdddcccdcc", "dddccddcdd", "ddccdcdcdc", "ddddccccdd", "dccdccdddd", "cddccdcccc", "cddcdccdcc", "cddcdddcdc", "dcdcdccdcd", "ddcdddccdc", "cccccddccc", "cddcdcccdc", "ddddcccdcd", "cdcdcddccd", "ccdddddcdd", "ddddcccccc", "ddddddcdcd", "cdddcddcdc", "ccccccdcdd", "dcdddcccdc", "dcdcdccddc", "ddcdcdccdc", "ddddcdcdcc", "ddcddddcdd", "ccddddcdcd", "dddddcdddd", "cddccdccdc", "dcddcccccd", "ccdcdcdcdc", "dddccccdcc", "ccddddcccd", "cddddcdddd", "ddccdcddcd", "ddccdccccd", "dcddcdcdcc", "dddccdddcd", "cddcccccdd", "ccdccddcdc", "dcddcccccc", "dccccdcddd", "ccdccdcccc", "cddcdccccd", "cdcccdccdd", "ccccdccccc", "cdcddccddd", "cdcccccddd", "cddddcddcc", "ddcdcdcdcd", "ccdcdddddc", "dccdcccddd", "cddcdcdccd", "cdcddddccd", "ddddddccdd", "ddccdcddcc", "dcddccdddc", "ccdccccccc", "ddcdcddccd", "dddcccdccd", "dcdccddcdc", "ddccdcdccd", "cddddcdcdd", "ccccdddcdd", "cdcddcdddd", "cccccdddcc", "cdcdddddcc", "dcccccddcc", "dccddddddc", "cdddcccddc", "dcccccdddd", "cdcdddccdc", "dcdccdddcc", "dddddcddcd", "ccdddcdddd", "ddcddcddcc", "cdcddccccd", "ccdcdccdcd", "ccdcdddccc", "ddcccccccd", "ddccdddddd", "dccddccccd", "ccccddcdcd", "cdcccddcdd", "cdcdcddddc", "ddddcdccdc", "cdcdddccdd", "dccccdcdcd", "ddcccccdcd", "dcccccdccc", "cccccccddc", "dcddcddccd", "ddccccdccd", "dddcdddddc", "cdcccdddcd", "ccdcddcccd", "cddcccdccd", "dcdcccdcdc", "dcddcdccdc", "dddcddddcc", "ddddddcccd", "cccdddddcc", "dccdddddcc", "ddcccccddd", "dccdccdccc", "ccdcddcdcd", "dccddccdcd", "dcdccccdcd", "cdddccccdc", "cdcddddccc", "cdccdccdcd", "dccccdccdd", "cdcdddcdcc", "dddccddddc", "cddddcddcd", "cdccdccccc", "dcdddddccc", "cccccdcccd", "dcdccdcddd", "ddcccddddc", "cdccddddcc", "dcddddcdcc", "dcdccddddc", "dccccccccc", "cdccdddddd", "dddcddddcd", "ccddddccdd", "ddddccdcdc", "ccccdccccd", "ddddccccdc", "dcccdccddc", "ccccdddddd", "dcccdcddcd", "dcdcddccdc", "cccdddccdc", "dddddccdcd", "dddccdcccc", "cdccccdccd", "cdccddccdd", "cccddcdcdc", "dddcdccddd", "dccddddcdd", "cccdcccccc", "dddccddccc", "ddddcddcdc", "dddccdcdcd", "dcccddcccd", "dcdcdddcdd", "dddddcdcdc", "ddcdccddcd", "ccccdcddcd", "ccdcdcdcdd", "ddcccdccdc", "dcccdccddd", "ccddccdddd", "cdddcccdcd", "cdccddcddc", "dcdccccccc", "dcccdcdddd", "dddcddcccd", "dddcdcccdd", "ccdddcdccc", "cdcdcccccd", "dcdccccdcc", "ccdcddcdcc", "ccdcdcddcd", "cddddddccc", "dcddcdcddc", "cddcccccdc", "dccdcdcddc", "dcdccdccdd", "ccddccddcc", "ddcdcddccc", "dcddcccdcd", "ddcdcccddc", "cdccdcddcc", "ddddcddccc", "cdddccdccc", "dcddcdcddd", "ccdccdcdcc", "cdcdcdcccd", "dccdddccdd", "ccddddcccc", "dcdddccccc", "ccdccddcdd", "cdcdccccdc", "cdcddddddd", "cdddcdcccc", "cccccccdcc", "dddccdcddd", "ccddccdddc", "cdddddddcd", "ddccccdddc", "cccccddddc", "dcccccccdc", "dcdddddccd", "ddddddccdc", "ddddcdcccd", "cccdcddddd", "ddcdddcdcc", "dcccdccccc", "ddddccdddd", "cdcdccdccc", "ccdcccdccc", "dddccccccd", "dccdcddcdd", "ccdddcdcdd", "ddddcccccd", "dccdcddccc", "cdccccdddc", "ccccccdcdc", "cdcdcdccdc", "dcdcddcccc", "cdddcddccc", "cccdddccdd", "ddcccddddd", "dccccddcdc", "cccccdcddc", "ddcccccdcc", "ccdcdddccd", "cdccccddcc", "dccdcccddc", "ddddddcddc", "dcdddddddd", "ccdccccdcd", "dccdcddddc", "ddddddddcd", "cddcdccddc", "ccccdccddd", "dddddcdddc", "cddccdccdd", "dddcddccdc", "cdccccdddd", "cccccdcddd", "cddccddccd", "dddddccccd", "ccddcdcddd", "dcdcddcdcc", "cdccddddcd", "dddcdddcdc", "cddccccccd", "ddcddcdccc", "cdcddcddcc", "ddddcccddc", "dddcccdddd", "cdccdccccd", "dcdcddcddd", "dddcccccdc", "ddcdccccdc", "cdddddccdc", "ddddcddccd", "dcccddddcc", "cccddddddc", "cccccddccd", "cdccdcccdd", "cdddddcddc", "ddcdcccddd", "cccddccccd", "ccdcccdccd", "dddcdcddcc", "ccddcccccc", "cddddccdcd", "ddcdcdcddd", "cdccdccdcc", "dcdcdddddd", "ccddcddcdd", "dccddddcdc", "ddccdddcdd", "ccdcccccdd", "dddcdccccc", "cdcddcddcd", "ddddccdccc", "ddccdccddd", "dcccddccdd", "ccccddcccc", "ccdcccdddc", "dcddcccdcc", "cddcdddccc", "dcccdddcdd", "cddddcdccc", "dddcdddddd", "ccddcddccd", "cdddddcdcd", "dcccdccdcd", "dcdcccdccd", "cddcdddcdd", "dcccddcccc", "dcddcddddc", "ccccdddccd", "dcddccddcd", "ccdcddccdd", "ddccddcdcd", "dccccddddd", "dddccdcdcc", "ddcddddddc", "dccdccdcdc", "dccdccccdd", "cdddcdccdc", "dcdcdccddd", "cddccccccc", "ccdcdcdddc", "ccdcdcdccc", "cccddcdddc", "dcdcdddccd", "ddddcdcccc", "dddddddddc", "ccdcdddcdc", "dcccdcdccd", "ccccddccdd", "cdcdddcccc", "cdcccccccc", "cccdddcdcc", "ddcdccccdd", "cdccdcdddd", "ccdccdcddd", "dddcdccdcd", "dcdcdcddcc", "cccccddcdc", "cdccdccddc", "cdcccdddcc", "dddccdccdc", "cccddccdcd", "dccdcccdcc", "cddcddcddd", "dcccdcdccc", "ccdddcdccd", "cdccdddcdc", "dddcdcdcdd", "dcdcddddcc", "cdccdcddcd", "cddcccdcdc", "ccccdddddc", "ccddddddcc", "dccdccdcdd", "ccccccddcd", "cccddccddc", "cddcdddccd", "cccddccccc", "dccdcdcccd", "dccccddccd", "ccddddddcd", "cdcdccdcdc", "ccdcdcdddd", "dcdddccddd", "ccddcccdcc", "dcdcccdcdd", "dcdddcccdd", "dccddcddcc", "dccccdcdcc", "ccccdcdccc", "cddcdcdccc", "cdcddcdccd", "dccdcccccc", "cdccdcdcdc", "dddcccddcc", "cdcddcccdc", "ccddddccdc", "dddccccccc", "ccdddddddd", "dddddccdcc", "ddcdcccccd", "dddccdcccd", "cccddcdccc", "ccdcccddcc", "dddcddcdcd", "cccdccdccc", "ddddcdddcc", "cccdcddccc", "ccccdcdccd", "cddcccdddd", "cdddcdcddc", "cdccddcdcc", "cdddcccccd", "cddcddcccc", "ddccccdcdd", "cddccccddc", "cddcdccccc", "cddddccddd", "cdcccccccd", "dddcdcdccc", "dcdddccdcd", "cdcdcddccc", "ccdccdddcc", "ccccdcdddc", "ddcddccccd", "cddddcccdc", "cccdcdcddc", "ccdcccdddd", "cdddcdccdd", "ddddcdcdcd", "ddccddddcc", "ddcccdccdd", "dccccdccdc", "ddddcdcddc", "cccdcdccdc", "ddcdddcccc", "ccdddccdcc", "ccdcddcddc", "dcddcdccdd", "dcddcddcdd", "ccddcdcdcc", "dcddccdccd", "ccccccddcc", "cddccccdcc", "dcccccdcdd", "ddddcddcdd", "cddcddccdc", "dccddccddd", "cddddccddc", "cdcdddcccd", "cccdddcdcd", "ddcdcdcccd", "ddcdcdccdd", "cccddccdcc", "cdddcdddcd", "dcdcdcdddc", "ddccddccdd", "cdcccddcdc", "cdcdcccdcc", "dccccccccd", "ddcdcdcddc", "dddcdcddcd", "dcdccccddc", "ccccccdccc", "dcdcdcdcdd", "cccdcccdcc", "dddccccddc", "ddcdcddddc", "ddcccddccc", "dddcdcdcdc", "ccddcddddd", "cdddcccddd", "cccdcddcdc", "dcdcccdccc", "cddddccdcc", "cdccdddccd", "ddccdcdddd", "ddccddcdcc", "ccdccddccd", "ddccdddccc", "dcddcddccc", "cdddccdddc", "ddcccdddcc", "cddccdddcc", "cdcddcdcdd", "dddcccdccc", "cddccddcdd", "cdcccccdcd", "ddcccccccc", "cdcccccdcc", "cddccddddc", "ddccdcccdd", "dcdddcddcd", "cdddcddddc", "cdcccdcdcd", "cdccccdccc", "cddcdcdddc", "dcddddccdc", "dcdcddcddc", "dcdddcdcdc", "dddcdddccd", "cddccdcddc", "cddcddddcc", "cddddcccdd", "ddcdccdccc", "ddcccddccd", "ddddddcccc", "ddccdccdcc", "ccdddcddcc", "cdccccccdd", "dccdddcccd", "ddddcdddcd", "ccdddcccdd", "ddccccccdd", "cdddcdcddd", "ddcccdddcd", "cccdccdcdd", "cdcdcdccdd", "cdcdccddcd", "ddcddcccdc", "cccddcdddd", "dcddccccdd", "ccdccdccdc", "dcdccddccd", "ccdccdcccd", "ddccddddcd", "ccccdccddc", "ccccccdddc", "cdcddcdcdc", "ccdccdccdd", "ccdcddcddd", "ccdddcddcd", "ccddcddcdc", "cdccccdcdc", "dccdddcddd", "dcccccccdd", "ddccdcdcdd", "cdcddccdcc", "cdcddccccc", "ccdcccdcdc", "cccdcdcccc", "dddcdcdccd", "cdddddddcc", "cccdcdcdcc", "cccddddccd", "dccccccdcc", "dddddccddc", "dccdcdcddd", "dccdcddccd", "dccddcdddd", "cddccddddd", "ddccddcddc", "ccccdddcdc", "dccdcddddd", "dcdddcdddc", "ddddcccddd", "ddcdcdddcd", "cdddcddddd", "ccddccddcd", "dcccddccdc", "ccddccccdd", "ddccddcccd", "cccdddcccc", "cdddccdccd", "ccdddddccc", "ccddcdccdc", "ccccdcddcc", "ccddddcddc", "dddccccdcd", "ddcddddccc", "dddddcccdd", "dcdcccdddd", "ddccccccdc", "cddccddccc", "dcdcdddccc", "ddccccdcdc", "dccdddcdcc", "dcdddddcdd", "ccdccccdcc", "ccccdccdcc", "cccdccddcd", "dddcdccddc", "ddcdcdcccc", "dcdcccccdd", "dddddddddd", "dcdcdddcdc", "cddcddccdd", "dcddddcccd", "ccddccdcdc", "dddccdccdd", "dcdddccddc", "dddddccddd", "cdcccdcddc", "cdccddcccd", "cccdccddcc", "cdccccdcdd", "ddcdccdcdd", "dcddccddcc", "cddcdcccdd", "cdddccdddd", "dddddcddcc", "ddddddcdcc", "ccddcddddc", "dddcddcccc", "ccddcccccd", "cdcccdcccc", "dcdccdcdcd", "dcdcdccccd", "dccdddcccc", "dcddddcccc", "ccddccdccd", "cddddddddc", "cccdcccccd", "dcdccdcccd", "dcccdddccc", "ccddddcdcc", "cdcdcddddd", "dccdddddcd", "dccccddddc", "dcccdddddc", "ccdddcdcdc", "ddcccdcccc", "cdccdcdddc", "ccdcdccddd", "ddcccdcddc", "cccddddcdd", "cddcddcdcc", "cddddcdcdc", "cdcdccdcdd", "cdcddcdddc", "cdcdcdcdcd", "cccdccdcdc", "ccdcccccdc", "cddcdcdddd", "dccddcdccd", "dddccddcdc", "dccdddcdcd", "ccdcdccccc", "dccccccddc", "ccdcddddcd", "dcdcccddcc", "dccddccddc", "cccccdcdcc", "dddccccddd", "cddccdcdcc", "ccddccdccc", "cccdcdcdcd", "dcccdddddd", "cdcdccdddc", "dcdddccccd", "cddccddcdc", "cccccccccc", "cdddccdcdd", "ddcdddcccd", "ddcddcddcd", "ccddcccddc", "cdddddcdcc", "ddccddcccc", "ddcddcdddc", "cdddcddccd", "ccdcdcccdc", "ddcddccdcd", "ccdddccddc", "cdddcdcdcc", "cdcdccdddd", "dcdcdcdccd", "cccdddcddd", "cccdcccddc", "dcccddcddc", "cccdccccdd", "dcddccdcdd", "cccccdddcd", "ccddcccdcd", "cdcdcdddcd", "cccddddcdc", "cddcdccdcd", "dcdccccccd", "ccccccdddd", "ccdcdddcdd", "dccddcdcdd", "cdcdccccdd", "dddcddcdcc", "dddddccccc", "dddcddccdd", "cdddddcccc", "cccdccdddc", "cddcccdddc", "dcddccdcdc", "cdcdcdcccc", "dccddcdccc", "dccddddccd", "cdccddcccc", "dcdcdcdcdc", "ccdddcdddc", "cdccdcccdc", "dcdcddcccd", "cdcddccdcd", "cccdcddddc", "dccdcdcdcd", "cdcdcdddcc", "cddcdcdcdc", "ddcddddccd", "dcdcdcddcd", "ddcddcdcdd", "dcccddcdcc", "dcdddcdccd", "ccddcdddcc", "dcddcddddd", "dccddccccc", "ddddccddcd", "dcddddcddd", "cccdcccdcd", "cccdccdccd", "dddddddccc", "cddddccccc", "cccdcdddcd", "cdddccdcdc", "dddcccdcdd", "dccdcccccd", "ddccdcccdc", "dcddddddcd", "dddddcdcdd", "ddcddccddc", "ccddddcddd", "ddccddccdc", "cccdcdcccd", "dcccccdddc", "dcdccddddd", "cdcdddddcd", "dcdccdddcd", "cccccdcdcd", "cdcdcccddc", "ddcdddddcd", "cdccdcdcdd", "cccddcddcc", "ccccdcccdd", "cdddddcddd", "dccdccdccd", "dddcdcdddd", "ccdcddddcc", "ddcddddddd", "cdcccdccdc", "cccddddccc", "dcdccddccc", "cddccdcddd", "ddcccccddc", "cdcdddcdcd", "dddcdccccd", "dcddddccdd", "cdcddddcdc", "dcccdddccd", "dddddcccdc", "ccccccccdd", "cccddcdcdd", "cddccdcdcd", "dcdcdcdddd", "dcccdccdcc", "dddddcdccd", "cccdddcddc", "dcddddddcc", "cddcdccddd", "cdcdcddcdd", "dccccdcccc", "cddddddcdd", "dcccdcdcdd", "cdddcdddcc", "ccddcdccdd", "ccdccccccd", "cddddccccd", "ccdccddddc", "ddcccdcdcc", "dcdccdccdc", "dddddddcdd", "dccddddccc", "ccdccccddc", "ccdccdcdcd", "dcddccdddd", "ccdddcccdc", "cccccddcdd", "cddccdcccd", "dcddddcdcd", "cdcccddccc", "cddccdddcd", "dcccdccccd", "ccdcccddcd", "ddcddccdcc", "dccdccddcd", "ddcdcdcdcc", "dcccdcdcdc", "ddddcdcddd", "cccdddcccd", "cddddddcdc", "cdccdcdccd", "ddcdcddddd", "cdcccddddc", "cddcdddddc", "cdcdcddcdc", "cccccddddd", "dddcddcddc", "ccccddcddd", "ccddcddccc", "dcdcdccccc", "dcdddddcdc", "ddcdccdddc", "ddcdddcddd", "dccddcccdc", "dcccdcccdd", "ccdddccdcd", "dcdcdccdcc", "cccdcddcdd", "ccdddccccd", "ddcdcccdcc", "dddccddddd", "ddcccddcdc", "ddccdccdcd", "dccdcdddcc", "cccccccdcd", "dcddddcddc", "dcccccdccd", "dcdddccdcc", "cdcddddddc", "cdccdddcdd", "cdcccdcccd", "dccdddcddc", "cccccdcccc", "ddcccdcccd", "ddcccddcdd", "ddccdddcdc",
	}
	w = Constructor745(words)
	_ans = []int{63, 103, 443, 956}
	for idx, sp := range [][]string{
		{"ddcccc", "ccddcd"},
		{"ddcdcdd", "cdcddcdd"},
		{"ddddcccc", "dcccccd"},
		{"cdd", "cdccddd"},
	} {
		if r := w.F(sp[0], sp[1]); r != _ans[idx] {
			t.Fatalf("prefix: %s, suffix: %s, expect %d get %d\n", sp[0], sp[1], _ans[idx], r)
		}
	}

	words = []string{"b", "b", "b", "b", "b", "b", "b"}
	w = Constructor745(words)
	if r := w.F("a", "a"); r != -1 {
		t.Fatalf("expect 0 get %d", r)
	}
}
