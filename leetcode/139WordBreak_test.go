package leetcode

import "testing"

func TestWordBreak(t *testing.T) {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	if !wordBreak(s, wordDict) {
		t.Fatalf("expect true get false")
	}

	s = "applepenapple"
	wordDict = []string{"apple", "pen"}
	if !wordBreak(s, wordDict) {
		t.Fatalf("expect true get false")
	}

	s = "catsandog"
	wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	if wordBreak(s, wordDict) {
		t.Fatalf("expect false get true")
	}
}
