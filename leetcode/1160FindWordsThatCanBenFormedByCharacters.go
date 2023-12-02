package leetcode

func countCharacters(words []string, chars string) int {
	cc := [26]int{}
	for _, c := range chars {
		cc[c-'a']++
	}

	ans := 0
	tmp := [26]int{}
	for _, word := range words {
		for _, c := range word {
			tmp[c-'a']++
		}
		ok := true
		for i := 0; i < 26; i++ {
			if tmp[i] > cc[i] {
				ok = false
			}
			tmp[i] = 0
		}
		if ok {
			ans += len(word)
		}
	}
	return ans
}
