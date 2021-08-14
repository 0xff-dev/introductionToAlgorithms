package leetcode

import (
	"bytes"
)

func multiply(num1, num2 string) string {
	l1, l2 := len(num1), len(num2)

	a := make([]byte, 0)
	first := true
	buf := bytes.NewBufferString("")
	addIdx := 1
	for end := l2 - 1; end >= 0; end-- {
		cf := uint8(0)
		n := num2[end] - '0'
		for ie := l1 - 1; ie >= 0; ie-- {
			tmp := n*(num1[ie]-'0') + cf
			cf = tmp / 10
			tmp %= 10
			buf.WriteByte(tmp + '0')
		}

		if cf > 0 {
			buf.WriteByte(cf + '0')
		}

		bs := buf.Bytes()
		buf.Reset()
		if first {
			a = append(a, bs...)
			first = false
			continue
		}
		// add bs and a
		bsIdx := 0
		innerCf := uint8(0)
		for idx := addIdx; idx < len(a); idx, bsIdx = idx+1, bsIdx+1 {
			tmp := a[idx] - '0' + bs[bsIdx] - '0' + innerCf
			innerCf = tmp / 10
			a[idx] = tmp%10 + '0'
		}

		for ; bsIdx < len(bs); bsIdx++ {
			tmp := bs[bsIdx] - '0' + innerCf
			innerCf = tmp / 10
			tmp %= 10
			a = append(a, tmp+'0')
		}
		if innerCf > 0 {
			a = append(a, innerCf+'0')
		}
		addIdx++
	}

	for s, e := 0, len(a)-1; s < e; s, e = s+1, e-1 {
		a[s], a[e] = a[e], a[s]
	}

	idx := 0
	for ; idx < len(a) && a[idx] == '0'; idx++ {
	}
	if idx == len(a) {
		return "0"
	}
	return string(a[idx:])
}
