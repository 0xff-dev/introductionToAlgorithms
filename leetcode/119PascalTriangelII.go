package leetcode

func getRow(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	ans[0] = 1
	for idx := 1; idx <= rowIndex; idx++ {
		for in := idx; in >= 1; in-- {
			ans[in] += ans[in-1]
		}
	}
	return ans
}
