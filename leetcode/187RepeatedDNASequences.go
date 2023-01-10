package leetcode

func findRepeatedDnaSequences(s string) []string {
	ans := make([]string, 0)
	exist := make(map[string]uint8)
	// abcd =2
	for i := 0; i <= len(s)-10; i++ {
		s1 := s[i : i+10]
		v, ok := exist[s1]
		if !ok {
			exist[s1] = 1
			continue
		}
		exist[s1] = 2
		if v == 1 {
			ans = append(ans, s1)
		}
	}
	return ans
}
