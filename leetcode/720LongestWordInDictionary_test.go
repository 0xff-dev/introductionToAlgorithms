package leetcode

import "testing"

func TestLongestWord(t *testing.T) {
	words := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	if r := longestWord(words); r != "apple" {
		t.Fatalf("expect apple get %s", r)
	}
	words = []string{"w", "wo", "wor", "worl", "world"}
	if r := longestWord(words); r != "world" {
		t.Fatalf("expect world get %s", r)
	}
}
