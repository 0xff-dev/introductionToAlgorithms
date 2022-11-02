package leetcode

import "testing"

func TestReversePrfix(t *testing.T) {
	word := "abcdefd"
	exp := "dcbaefd"
	if r := reversePrefix(word, 'd'); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
