package leetcode

type Trie struct {
	end   bool
	child [26]*Trie
}

func Constructor208() Trie {
	return Trie{child: [26]*Trie{}}
}

func (this *Trie) Insert(word string) {
	root := this
	for i, b := range word {
		if root.child[b-'a'] == nil {
			root.child[b-'a'] = &Trie{end: i == len(word)-1, child: [26]*Trie{}}
			root = root.child[b-'a']
			continue
		}
		if i == len(word)-1 {
			root.child[b-'a'].end = true
			return
		}
		root = root.child[b-'a']
	}
}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}
	root := this
	for i, b := range word {
		if i == len(word)-1 {
			return root.child[b-'a'] != nil && root.child[b-'a'].end
		}
		if root.child[b-'a'] == nil {
			return false
		}
		root = root.child[b-'a']
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}
	root := this
	for _, b := range prefix {
		if root.child[b-'a'] == nil {
			return false
		}
		root = root.child[b-'a']
	}
	return true
}
