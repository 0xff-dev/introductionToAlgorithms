package leetcode

import "testing"

func TestSortVowels(t *testing.T) {
	s := "lEetcOde"
	exp := "lEOtcede"
	if r := sortVowels(s); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
	s = "lYmpH"
	exp = "lYmpH"
	if r := sortVowels(s); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
