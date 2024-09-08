package leetcode

import "testing"

func TestExpressiveWords(t *testing.T) {
	s, words, exp := "heeellooo", []string{"hello", "hi", "helo"}, 1
	if r := expressiveWords(s, words); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s, words, exp = "zzzzzyyyyy", []string{"zzyy", "zy", "zyy"}, 3
	if r := expressiveWords(s, words); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
