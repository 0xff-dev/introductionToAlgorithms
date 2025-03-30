package leetcode

func doesAliceWin(s string) bool {
	cnt := 0
	for _, e := range s {
		if e == 'a' || e == 'e' || e == 'i' || e == 'o' || e == 'u' {
			cnt++
		}
	}
	return cnt != 0
}
