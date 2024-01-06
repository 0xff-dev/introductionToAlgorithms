package leetcode

type MagicDictionary struct {
	source map[int]map[string]struct{}
}

func Constructor676() MagicDictionary {
	return MagicDictionary{source: make(map[int]map[string]struct{})}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		if _, ok := this.source[len(word)]; !ok {
			this.source[len(word)] = make(map[string]struct{})
		}
		this.source[len(word)][word] = struct{}{}
	}
}

func (this *MagicDictionary) search(a, b string) bool {
	diff := 0
	for idx := range a {
		if a[idx] != b[idx] {
			diff++
		}
		if diff > 1 {
			return false
		}
	}
	return diff == 1
}
func (this *MagicDictionary) Search(searchWord string) bool {
	words, ok := this.source[len(searchWord)]
	if !ok {
		return false
	}
	for word := range words {
		if this.search(word, searchWord) {
			return true
		}
	}
	return false
}
