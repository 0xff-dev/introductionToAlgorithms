package leetcode

func maxDifference(s string) int {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}
	a1, a2 := 0, 101
	for i := 0; i < 26; i++ {
		if cnt[i]&1 == 1 {
			a1 = max(a1, cnt[i])
			continue
		}
		if cnt[i] == 0 {
			continue
		}
		a2 = min(a2, cnt[i])
	}
	return a1 - a2
}
