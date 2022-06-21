package leetcode

import "testing"

func TestTrieNode(t *testing.T) {
	baseStr := "abcdefg"
	trie := &TrieNode{Child: [26]*TrieNode{}, Count: 0}
	trie.Insert(baseStr)

	if !trie.Search("a") {
		t.Fatalf("expect true get false")
	}
	if !trie.Search("abc") {
		t.Fatalf("expect true get false")
	}

	if !trie.Search(baseStr) {
		t.Fatalf("expect true get false")
	}

	if trie.Search("abce") {
		t.Fatalf("expect false get true")
	}

	if trie.Search("abcdefgh") {
		t.Fatalf("expect false get true")
	}
}