package leetcode

func minHeightShelves(books [][]int, shelf_width int) int {
	if len(books) == 0 {
		return 0
	}

	dp := make([]int, len(books))
	dp[0] = books[0][1]

	for idx := 1; idx < len(books); idx++ {
		length, maxHeight := 0, 0
		dp[idx] = 1000000
		for j := idx; j >= 0; j-- {
			length += books[j][0]
			if books[j][1] > maxHeight {
				maxHeight = books[j][1]
			}
			if length <= shelf_width {
				if j == 0 {
					if dp[idx] > maxHeight {
						dp[idx] = maxHeight
					}
				} else {
					if dp[idx] > dp[j-1]+maxHeight {
						dp[idx] = dp[j-1] + maxHeight
					}
				}
				continue
			}
			break
		}
	}
	return dp[len(books)-1]
}
