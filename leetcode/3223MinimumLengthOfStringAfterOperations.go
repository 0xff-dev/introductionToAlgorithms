package leetcode

func minimumLength3223(s string) int {
	count := [26]int{}
	for _, b := range s {
		count[b-'a']++
	}
	ans := 0
	var add int
	for i := range 26 {
		if count[i] != 0 {
			add = 1
			if count[i]&1 == 0 {
				add = 2
			}
			ans += add
		}
	}
	return ans
}
