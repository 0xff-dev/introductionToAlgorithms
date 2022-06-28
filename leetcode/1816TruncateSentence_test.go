package leetcode

import "testing"

func TestTruncateSentence(t *testing.T) {
	s, k := "Hello how are you Contestant", 4
	if r := truncateSentence(s, k); r != "Hello how are you" {
		t.Fatalf("expect 'Hello how are you' get %s", r)
	}

	s, k = "What is the solution to this problem", 4
	if r := truncateSentence(s, k); r != "What is the solution" {
		t.Fatalf("expect 'What is the solution' get %s", r)
	}
	s, k = "chopper is not a tanuki", 5
	if r := truncateSentence(s, k); r != s {
		t.Fatalf("expect '%s' get %s", s, r)
	}
}
