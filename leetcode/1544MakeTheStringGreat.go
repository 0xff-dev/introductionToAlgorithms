package leetcode

func makeGood(s string) string {
	// leEeetcode
	// leeeetcode
	// A:65 a: 97 diff=32
	length := len(s)
	ans := make([]byte, length)
	if length == 0 {
		return ""
	}
	idx, walker := -1, 0
	for ; walker < length; walker++ {
		if idx == -1 || !byteEqual(ans[idx], s[walker]) {
			idx++
			ans[idx] = s[walker]
			continue
		}
		idx--
	}
	return string(ans[:idx+1])
}

func byteEqual(a, b byte) bool {
	diff := int(a) - int(b)
	return diff == 32 || diff == -32
}
