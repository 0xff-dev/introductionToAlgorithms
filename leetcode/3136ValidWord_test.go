package leetcode

import "testing"

func TestIsValid3136(t *testing.T) {
	word, exp := "234Adas", true
	if r := isValid3136(word); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}

	word, exp = "b3", false
	if r := isValid3136(word); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	word, exp = "a3$e", false
	if r := isValid3136(word); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
