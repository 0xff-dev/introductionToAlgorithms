package leetcode

func strWithout3a3b(a int, b int) string {
	ret := make([]byte, 0, a+b)
	var (
		l      int
		writeA bool
	)

	for a > 0 || b > 0 {
		writeA = false
		l = len(ret)
		if l >= 2 && ret[l-1] == ret[l-2] {
			writeA = ret[l-1] == 'b'
		} else {
			writeA = a >= b
		}
		if writeA {
			ret = append(ret, 'a')
			a--
			continue
		}
		ret = append(ret, 'b')
		b--
	}

	return string(ret)
}
