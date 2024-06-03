package leetcode

import "testing"

func TestAppendCharacters(t *testing.T) {
	ss, tt := "coaching", "coding"
	exp := 4
	if r := appendCharacters(ss, tt); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	ss, tt = "abcde", "a"
	exp = 0
	if r := appendCharacters(ss, tt); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	ss, tt = "z", "abcde"
	exp = 5
	if r := appendCharacters(ss, tt); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
