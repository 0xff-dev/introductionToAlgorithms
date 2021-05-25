package leetcode

import "testing"

func TestWordPattern(t *testing.T) {
	pattern, s := "abba", "dog cat cat dog"
	if !wordPattern(pattern, s) {
		t.Fatalf("expect true get false")
	}

	pattern, s = "abba", "dog cat cat fish"
	if wordPattern(pattern, s) {
		t.Fatalf("expect false get true")
	}

	pattern, s = "aaaa", "dog cat cat dog"
	if wordPattern(pattern, s) {
		t.Fatalf("expect false get true")
	}

	pattern, s = "abba", "dog dog dog dog"
	if wordPattern(pattern, s) {
		t.Fatalf("expect flase get true")
	}
}
