package leetcode

import (
	"bytes"
)

func addBinary(a, b string) string {
	buf := bytes.NewBufferString("")

	ai, bi := len(a)-1, len(b)-1
	cf := uint8(0)
	for ; ai >= 0 && bi >= 0; ai, bi = ai-1, bi-1 {
		tmp := a[ai] - '0' + b[bi] - '0' + cf
		buf.WriteByte(tmp%2 + '0')
		cf = tmp / 2
	}

	for ; ai >= 0; ai-- {
		tmp := a[ai] - '0' + cf
		buf.WriteByte(tmp%2 + '0')
		cf = tmp / 2
	}
	for ; bi >= 0; bi-- {
		tmp := b[bi] - '0' + cf
		buf.WriteByte(tmp%2 + '0')
		cf = tmp / 2
	}
	if cf > 0 {
		buf.WriteByte('1')
	}
	bs := buf.Bytes()
	for s, e := 0, len(bs)-1; s < e; s, e = s+1, e-1 {
		bs[s], bs[e] = bs[e], bs[s]
	}
	return string(bs)
}
