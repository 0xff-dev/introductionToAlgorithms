package leetcode

func isInterleave(s1 string, s2 string, s3 string) bool {
	// 交替的来, 动态规划

	ls1, ls2, ls3 := len(s1), len(s2), len(s3)
	if ls1+ls2 != ls3 {
		return false
	}
	if ls3 == 0 {
		return true
	}
	if ls1 == 0 {
		return s2 == s3
	}
	if ls2 == 0 {
		return s1 == s3
	}

	cache := make([][]int, ls1)
	for i := 0; i < ls1; i++ {
		cache[i] = make([]int, ls2)
		for j := 0; j < ls2; j++ {
			cache[i][j] = -1
		}
	}

	var (
		dfs1 func(int, int, int) bool
		dfs2 func(int, int, int) bool
	)
	dfs1 = func(i, j, si int) bool {
		if i == ls1 && j == ls2 {
			return true
		}

		if i >= ls1 {
			return s2[j:] == s3[si:]
		}
		if j >= ls2 {
			return s1[i:] == s3[si:]
		}
		if cache[i][j] != -1 {
			return cache[i][j] == 1
		}
		cache[i][j] = 0
		for ie := i + 1; ie <= ls1; ie++ {
			l := ie - i
			a := s1[i:ie]
			if a != s3[si:si+l] {
				continue
			}

			for je := j + 1; je <= ls2; je++ {
				ll := je - j
				b := s2[j:je]
				if b != s3[si+l:si+l+ll] {
					continue
				}
				if dfs1(ie, je, si+l+ll) {
					cache[i][j] = 1
					return true
				}
			}
		}
		return false
	}
	dfs2 = func(i, j, si int) bool {
		if i == ls1 && j == ls2 {
			return true
		}

		if i >= ls1 {
			return s2[j:] == s3[si:]
		}
		if j >= ls2 {
			return s1[i:] == s3[si:]
		}
		if cache[i][j] != -1 {
			return cache[i][j] == 1
		}
		cache[i][j] = 0
		for je := j + 1; je <= ls2; je++ {
			ll := je - j
			b := s2[j:je]
			if b != s3[si:si+ll] {
				continue
			}
			for ie := i + 1; ie <= ls1; ie++ {
				l := ie - i
				a := s1[i:ie]
				if a != s3[si+ll:si+l+ll] {
					continue
				}
				if dfs2(ie, je, si+l+ll) {
					cache[i][j] = 1
					return true
				}
			}
		}
		return false
	}

	dfs1(0, 0, 0)
	if cache[0][0] == 1 {
		return true
	}
	for i := 0; i < ls1; i++ {
		for j := 0; j < ls2; j++ {
			cache[i][j] = -1
		}
	}
	dfs2(0, 0, 0)
	return cache[0][0] == 1
}
