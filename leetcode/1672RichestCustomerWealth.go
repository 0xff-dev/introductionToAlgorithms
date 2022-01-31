package leetcode

func maximumWealth(accounts [][]int) int {
	result := 0
	for row := 0; row < len(accounts); row++ {
		s := 0
		for col := 0; col < len(accounts[0]); col++ {
			s += accounts[row][col]
		}
		if s > result {
			result = s
		}
	}
	return result
}
