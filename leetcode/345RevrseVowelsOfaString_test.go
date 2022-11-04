package leetcode

import "testing"

func TestReverseVowels(t *testing.T) {
	s := "hello"
	exp := "holle"
	if r := reverseVowels(s); r != exp {
		t.Fatalf("expect %s get %v", exp, r)
	}

	s = "leetcode"
	exp = "leotcede"
	if r := reverseVowels(s); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
