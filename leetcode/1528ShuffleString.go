package leetcode

func restoreString(s string, indices []int) string {
	res := make([]byte, len(s))

	for sIdx, idx := range indices {
		res[idx] = s[sIdx]
	}

	return string(res)
}
