package leetcode

func leastBricks(wall [][]int) int {
	rows := len(wall)
	edges := make(map[int64]int)
	for _, row := range wall {
		cur := int64(0)
		for i, col := range row {
			if i != len(row)-1 {
				cur += int64(col)
				edges[cur]++
			}
		}
	}
	ans := rows

	for _, e := range edges {
		ans = min(ans, rows-e)
	}
	return ans
}
