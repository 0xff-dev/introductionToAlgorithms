package leetcode

func wordBreak140(s string, wordDict []string) []string {
	words := make(map[string]struct{})
	ml := 0
	for _, word := range wordDict {
		words[word] = struct{}{}
		ml = max(ml, len(word))
	}

	cache := make(map[int][]string)
	var dfs func(int) []string
	dfs = func(index int) []string {
		if v, ok := cache[index]; ok {
			return v
		}
		if index == len(s) {
			return []string{}
		}
		// ca tsanddog
		// 01 2
		ans := []string{}
		for end := index + 1; end <= len(s) && end <= index+ml; end++ {
			cur := s[index:end]
			if _, ok := words[cur]; ok {
				v := dfs(end)
				if len(v) == 0 {
					// 自己
					if end == len(s) {
						ans = append(ans, cur)
					}
				} else {
					for _, vv := range v {
						ans = append(ans, cur+" "+vv)
					}
				}
			}
		}
		cache[index] = ans
		return ans
	}
	return dfs(0)
}
