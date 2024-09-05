package leetcode

func missingRolls(rolls []int, mean int, n int) []int {
	/*
		target := (n + len(rolls)) * mean
		for _, r := range rolls {
			target -= r
		}

		// 无法组成
		if n*6 < target || target < n {
			return []int{}
		}
		ans := make([]int, n)
		loop := 0
		// 3
		for loop < 6 && target > 0 {
			// 一次去掉1
			if target >= n {
				target -= n
				for i := range n {
					ans[i]++
				}
			} else {
				for i := range target {
					ans[i]++
				}
				target = 0
			}
			loop++
		}
		return ans
	*/
	target := (n + len(rolls)) * mean
	for _, r := range rolls {
		target -= r
	}

	if n*6 < target || target < n {
		return []int{}
	}
	ans := make([]int, n)
	base := target / n
	for i := range n {
		ans[i] = base
	}
	target %= n
	for i := range target {
		ans[i]++
	}
	return ans
}
