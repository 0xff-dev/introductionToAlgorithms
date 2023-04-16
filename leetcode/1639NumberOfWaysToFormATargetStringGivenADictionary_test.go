package leetcode

import "testing"

func TestNumWays(t *testing.T) {
	words := []string{"acca", "bbbb", "caca"}
	target := "aba"
	if r := numWays(words, target); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	words = []string{"abba", "baab"}
	target = "bab"
	if r := numWays(words, target); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}
