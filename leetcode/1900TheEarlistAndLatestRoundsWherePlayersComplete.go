package leetcode

import "math"

func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
	const maxN = 30
	var F [maxN][maxN][maxN]int
	var G [maxN][maxN][maxN]int
	if firstPlayer > secondPlayer {
		firstPlayer, secondPlayer = secondPlayer, firstPlayer
	}

	var dp func(n, f, s int) (int, int)
	dp = func(n, f, s int) (int, int) {
		if F[n][f][s] != 0 {
			return F[n][f][s], G[n][f][s]
		}
		if f+s == n+1 {
			return 1, 1
		}
		// F(n,f,s) = F(n,n+1-s,n+1-f)
		if f+s > n+1 {
			x, y := dp(n, n+1-s, n+1-f)
			F[n][f][s] = x
			G[n][f][s] = y
			return x, y
		}

		earliest := math.MaxInt32
		latest := math.MinInt32
		n_half := (n + 1) / 2
		if s <= n_half {
			// On the left or in the middle
			for i := 0; i < f; i++ {
				for j := 0; j < s-f; j++ {
					x, y := dp(n_half, i+1, i+j+2)
					earliest = min(earliest, x)
					latest = max(latest, y)
				}
			}
		} else {
			// s on the right
			s_prime := n + 1 - s
			mid := (n - 2*s_prime + 1) / 2
			for i := 0; i < f; i++ {
				for j := 0; j < s_prime-f; j++ {
					x, y := dp(n_half, i+1, i+j+mid+2)
					earliest = min(earliest, x)
					latest = max(latest, y)
				}
			}
		}

		F[n][f][s] = earliest + 1
		G[n][f][s] = latest + 1
		return F[n][f][s], G[n][f][s]
	}

	earliest, latest := dp(n, firstPlayer, secondPlayer)
	return []int{earliest, latest}
}
