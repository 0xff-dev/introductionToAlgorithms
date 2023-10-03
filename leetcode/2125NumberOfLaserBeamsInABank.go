package leetcode

func numberOfBeams(bank []string) int {
	l := len(bank)
	ones := make([]int, l)
	for row, str := range bank {
		for _, b := range str {
			if b == '1' {
				ones[row]++
			}
		}
	}
	ans := 0
	index := 0
	for ; index < l && ones[index] == 0; index++ {
	}

	// index 第一个有1的地方
	for next := index + 1; next < l; next++ {
		if ones[next] == 0 {
			continue
		}
		ans += ones[index] * ones[next]
		index = next
	}
	//fmt.Println(ones)
	return ans
}
