package datastructure

import "testing"

func TestTrieTree(t *testing.T) {
	root := NewTrieNode()
	for _, word := range []string{"word", "tomorrow", "today"} {
		root.Insert(word)
	}
	for _, searchWord := range []string{"word", "wo"} {
		t.Logf("word %s count: %d", searchWord, root.Search(searchWord))
	}
}
