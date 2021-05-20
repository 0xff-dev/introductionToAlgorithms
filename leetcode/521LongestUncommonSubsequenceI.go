package leetcode

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	lena, lenb := len(a), len(b)
	if lena > lenb {
		return lena
	}
	return lenb
}
