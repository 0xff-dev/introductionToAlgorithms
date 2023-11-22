package leetcode

/*
14 12 19 16 9
13 14 15 8 11
11 13 1
1
*/
func findDiagonalOrder1424(nums [][]int) []int {
	/*
		ans := make([]int, 0)
		l := len(nums)
		ml := 0
		for row := 0; row < l; row++ {
			if len(nums[row]) > ml {
				ml = len(nums[row])
			}

			r, c := row, 0
			for r >= 0 {
				if c < len(nums[r]) {
					ans = append(ans, nums[r][c])
				}
				r, c = r-1, c+1
			}
		}
		for col := 1; col < ml; col++ {
			r, c := l-1, col
			for r >= 0 {
				if c < len(nums[r]) {
					ans = append(ans, nums[r][c])
				}
				r, c = r-1, c+1
			}
		}
		return ans
	*/
	ans := make([]int, 0)
	q := [][2]int{{0, 0}}
	for len(q) > 0 {
		nq := make([][2]int, 0)
		for _, item := range q {
			r, c := item[0], item[1]
			ans = append(ans, nums[r][c])
			if c == 0 && r+1 < len(nums) {
				nq = append(nq, [2]int{r + 1, c})
			}
			if c+1 < len(nums[r]) {
				nq = append(nq, [2]int{r, c + 1})
			}
		}
		q = nq
	}
	return ans
}
