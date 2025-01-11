package leetcode

func canConstruct1400(s string, k int) bool {
	if len(s) < k {
		return false
	}
	count := [26]int{}
	for _, b := range s {
		count[b-'a']++
	}
	odd, even := 0, 0
	single := 0
	for i, n := range count {
		if n != 0 {
			single++
			if n&1 == 1 {
				odd++
				count[i]--
			}
			even += count[i]
		}
	}
	if odd > k {
		return false
	}
	if k == odd {
		return true
	}
	return k-odd <= even
}
