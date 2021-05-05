package leetcode

func oddCells(m int, n int, indices [][]int) int {
	rows := make([]int, m)
	cols := make([]int, n)
	for _, item := range indices {
		rows[item[0]]++
		cols[item[1]]++
	}

	odds := 0
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if (rows[row]+cols[col])&1 != 0 {
				odds++
			}
		}
	}
	return odds
}
