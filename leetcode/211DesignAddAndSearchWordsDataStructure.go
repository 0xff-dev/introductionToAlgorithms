package leetcode

type WordDictionary struct {
	end      bool
	children [26]*WordDictionary
}

func Constructor211() WordDictionary {
	return WordDictionary{
		children: [26]*WordDictionary{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	root := this
	for i, b := range word {
		if root.children[b-'a'] == nil {
			root.children[b-'a'] = &WordDictionary{children: [26]*WordDictionary{}}
		}
		root.children[b-'a'].end = i == len(word)-1
		root = root.children[b-'a']
	}
}

func (this *WordDictionary) dfs(word string, index int) bool {
	if index == len(word)-1 {
		if word[index] == '.' {
			for i := 0; i < 26; i++ {
				if this.children[i] != nil && this.children[i].end {
					return true
				}
			}
			return false
		}
		return this.children[word[index]-'a'] != nil && this.children[word[index]-'a'].end
	}
	if word[index] == '.' {
		for i := 0; i < 26; i++ {
			if this.children[i] != nil {
				if this.children[i].dfs(word, index+1) {
					return true
				}
			}
		}
		return false
	}
	if this.children[word[index]-'a'] == nil {
		return false
	}
	return this.children[word[index]-'a'].dfs(word, index+1)

}
func (this *WordDictionary) Search(word string) bool {
	return this.dfs(word, 0)
}
