package leetcode

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	input := "Heloo World"
	if r := lengthOfLastWord(input); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	input = " "
	if r := lengthOfLastWord(input); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	input = "hello world   "
	if r := lengthOfLastWord(input); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	input = "a"
	if r := lengthOfLastWord(input); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
