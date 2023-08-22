package leetcode

import "bytes"

// A=65-1
// B=66-2
// Z=90-26
func convertToTitle(columnNumber int) string {
	base := 64
	buf := bytes.NewBuffer([]byte{})
	for columnNumber > 0 {
		mod := columnNumber % 26
		columnNumber /= 26
		if mod == 0 {
			mod = 26
			columnNumber--
		}
		buf.WriteByte(uint8(base + mod))
	}
	bs := buf.Bytes()
	for s, e := 0, len(bs)-1; s < e; s, e = s+1, e-1 {
		bs[s], bs[e] = bs[e], bs[s]
	}
	return string(bs)
}
