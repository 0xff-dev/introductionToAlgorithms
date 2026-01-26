package leetcode

func minSteps2186(s string, t string) int {
	ss := [26]int{}
	tt := [26]int{}
	bs := []byte(s)
	bt := []byte(t)
	for i := range bs {
		ss[bs[i]-'a']++
	}
	for i := range bt {
		tt[bt[i]-'a']++
	}
	var ret, diff int
	for i := range 26 {
		if ss[i] == tt[i] {
			continue
		}
		diff = ss[i] - tt[i]
		if diff < 0 {
			diff = -diff
		}
		ret += diff
	}
	return ret
}
