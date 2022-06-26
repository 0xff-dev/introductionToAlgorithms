package leetcode

func maxScore(cardPoints []int, k int) int {
	length := len(cardPoints)
	base := 0
	left := make([]int, k)
	for idx := 0; idx < k; idx++ {
		left[idx] = cardPoints[idx]
	}
	for idx := length - 1; idx >= length-k; idx-- {
		base += cardPoints[idx]
		cardPoints[idx] = base
	}
	base = 0
	ans := cardPoints[length-k] // select all right cards.
	for idx := 1; idx <= k; idx++ {
		base += left[idx-1]
		if idx == k {
			if base > ans {
				ans = base
			}
			continue
		}
		if r := base + cardPoints[length-k+idx]; r > ans {
			ans = r
		}
	}
	return ans
}

func maxScore1(cardPoints []int, k int) int {
	length := len(cardPoints)
	r := 0
	for idx := length-1; idx >= length-k; idx-- {
		r += cardPoints[idx]
	}
	ans := r


	l := 0
	for idx := 0; idx < k; idx++ {
		l += cardPoints[idx]
		r -= cardPoints[length-k+idx]
		if l+r > ans {
			ans = l+r
		}
	}
	return ans
}
