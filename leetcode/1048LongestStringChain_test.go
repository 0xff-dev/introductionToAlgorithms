package leetcode

import "testing"

func TestLongestStrChain(t *testing.T) {
	words := []string{"a", "b", "ba", "bca", "bda", "bdca"}
	if r := longestStrChain(words); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	words = []string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}
	if r := longestStrChain(words); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	words = []string{"a"}
	if r := longestStrChain(words); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}

func TestIsPredecessor(t *testing.T) {
	word1, word2 := "ahd", "ahed"
	if !isPredecessor(word1, word2) {
		t.Fatalf("epxect true get false")
	}
}
