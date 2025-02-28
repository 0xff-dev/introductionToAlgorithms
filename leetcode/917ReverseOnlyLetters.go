package leetcode

func ok917(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func reverseOnlyLetters(s string) string {
	bs := []byte(s)
	l, r := 0, len(bs)-1
	for l < r {
		for ; l < r && !ok917(bs[l]); l++ {

		}
		for ; r > l && !ok917(bs[r]); r-- {
		}
		bs[l], bs[r] = bs[r], bs[l]
		l, r = l+1, r-1
	}
	return string(bs)
}
