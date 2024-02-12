package leetcode

func numTilePossibilities(tiles string) int {
	count := [26]int{}
	for _, c := range tiles {
		count[c-'a']++
	}
	var dfs func() int
	dfs = func() int {
		r := 1
		for i := 0; i < 26; i++ {
			if count[i] == 0 {
				continue
			}
			count[i]--
			r += dfs()
			count[i]++
		}
		return 1

	}
	return dfs()
}
