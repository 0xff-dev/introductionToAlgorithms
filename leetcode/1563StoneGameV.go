package leetcode

func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)

	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + stoneValue[i]
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == j {
			return 0
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}

		maxScore := 0
		for k := i; k < j; k++ {
			leftSum := prefixSum[k+1] - prefixSum[i]
			rightSum := prefixSum[j+1] - prefixSum[k+1]

			var currentScore int
			if leftSum < rightSum {
				currentScore = leftSum + dfs(i, k)
			} else if leftSum > rightSum {
				currentScore = rightSum + dfs(k+1, j)
			} else {
				score1 := leftSum + dfs(i, k)
				score2 := rightSum + dfs(k+1, j)
				if score1 > score2 {
					currentScore = score1
				} else {
					currentScore = score2
				}
			}

			// 更新当前区间的最大得分
			if currentScore > maxScore {
				maxScore = currentScore
			}
		}

		memo[i][j] = maxScore
		return maxScore
	}

	return dfs(0, n-1)
}

/*
func stoneGameV(stoneValue []int) int {
	size, sum := len(stoneValue), stoneValue[0]
	for i := 1; i < size; i++ {
		sum += stoneValue[i]
	}
	start, end := 0, size-1

	leftSearch := func() (int, int, int) {
		got, s, e := 0, start, end
		leftScore, leftIndex := 0, start
		for ; leftIndex < end; leftIndex++ {
			leftScore += stoneValue[leftIndex]
			rightScore := sum - leftScore
			if leftScore >= rightScore {
				s, e, got = leftIndex+1, end, rightScore
				break
			}
		}
		return got, s, e
	}
	rightSearch := func() (int, int, int) {
		got, s, e := 0, start, end
		rightScore, rightIndex := 0, end
		for ; rightIndex > start; rightIndex-- {
			rightScore += stoneValue[rightIndex]
			leftScore := sum - rightScore
			if rightScore >= leftScore {
				s, e, got = start, rightIndex-1, leftScore
				break
			}
		}
		return got, s, e
	}
	var ret int
	for start < end {
		lScore, lStart, lEnd := leftSearch()
        //fmt.Printf("--lScore = %d, lStart = %d, lEnd = %d\n", lScore, lStart, lEnd)
		rScore, rStart, rEnd := rightSearch()
        //fmt.Printf("--rScore = %d, rStart = %d, rEnd = %d\n", rScore, rStart, rEnd)
		if lScore > rScore {
			ret += lScore
			sum = lScore
			start, end = lStart, lEnd
			continue
		}
		ret += rScore
		sum = rScore
		start, end = rStart, rEnd
	}
	return ret
}
*/
