package leetcode

func reversePrefix(word string, ch byte) string {
	end := 0
	for ; end < len(word); end++ {
		if word[end] == ch {
			break
		}
	}

	if end == len(word) {
		return word
	}

	bs := []byte(word)
	for s, e := 0, end; s < e; s, e = s+1, e-1 {
		bs[s], bs[e] = bs[e], bs[s]
	}

	return string(bs)
}
