package datastructure

// prefix tree
// a-z
type TrieNode struct {
	Count    int // word frequency
	Children []*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Count:    0,
		Children: make([]*TrieNode, 26),
	}
}

func (trie *TrieNode) Insert(str string) {
	walker := trie
	for _, _byte := range []byte(str) {
		if walker.Children[_byte-'a'] == nil {
			walker.Children[_byte-'a'] = NewTrieNode()
		}
		walker = walker.Children[_byte-'a']
	}
	walker.Count++
}

// Search returns the number of occurrences
func (trie *TrieNode) Search(str string) int {
	walker := trie
	for _, _byte := range []byte(str) {
		if walker.Children[_byte-'a'] == nil {
			return 0
		}
		walker = walker.Children[_byte-'a']
	}
	if walker == nil {
		return 0
	}
	return walker.Count
}
