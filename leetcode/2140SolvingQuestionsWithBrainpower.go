package leetcode

/*
type mode2140 struct {
	yes, no int64
}

func mostPoints(questions [][]int) int64 {
	ans := int64(questions[0][0])
	dp := make([]mode2140, len(questions))
	dp[0] = mode2140{yes: int64(questions[0][0]), no: int64(0)}

	for next := 1; next < len(questions); next++ {
		dp[next] = mode2140{yes: int64(questions[next][0]), no: int64(-1)}

		for pre := next - 1; pre >= 0; pre-- {
			expectStart := pre + questions[pre][1] + 1
			if expectStart <= next {
				x := dp[pre].no
				if dp[pre].yes > x {
					x = dp[pre].yes
				}
				if r := x + int64(questions[next][0]); r > dp[next].yes {
					dp[next].yes = r
				}
			} else {
				if dp[pre].no > dp[next].no {
					dp[next].no = dp[pre].no
				}
			}
		}

		if dp[next].yes > ans {
			ans = dp[next].yes
		}
		if dp[next].no > ans {
			ans = dp[next].no
		}
	}
	return ans
}
*/

func mostPoints(questions [][]int) int64 {
	l := len(questions)
	dp := make([]int64, l)
	dp[l-1] = int64(questions[l-1][0])
	for pre := l - 2; pre >= 0; pre-- {
		dp[pre] = int64(questions[pre][0])
		reach := pre + questions[pre][1] + 1
		if reach < l {
			dp[pre] += dp[reach]
		}
		if dp[pre+1] > dp[pre] {
			dp[pre] = dp[pre+1]
		}
	}
	return dp[0]
}
