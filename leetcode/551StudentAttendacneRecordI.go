package leetcode

func checkRecord551(s string) bool {
	a := 0
	l := 0
	for _, b := range s {
		if b == 'P' {
			l = 0
			continue
		}
		if b == 'A' {
			l = 0
			a++
			if a >= 2 {
				return false
			}
			continue
		}
		l++
		if l >= 3 {
			return false
		}
	}
	return true
}
