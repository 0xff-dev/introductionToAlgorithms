package leetcode

import "math"

const mod3655 = 1_000_000_007

func pow3655(x, y int64) int {
	res := int64(1)
	for y > 0 {
		if y&1 == 1 {
			res = res * x % mod3655
		}
		x = x * x % mod3655
		y >>= 1
	}
	return int(res)
}

func xorAfterQueries3655(nums []int, queries [][]int) int {
	n := len(nums)
	T := int(math.Sqrt(float64(n)))

	groups := make([][][]int, T)
	for i := 0; i < T; i++ {
		groups[i] = make([][]int, 0)
	}

	for _, q := range queries {
		l, r, k, v := q[0], q[1], q[2], q[3]
		if k < T {
			groups[k] = append(groups[k], []int{l, r, v})
		} else {
			for i := l; i <= r; i += k {
				nums[i] = int((int64(nums[i]) * int64(v)) % mod3655)
			}
		}
	}

	dif := make([]int64, n+T)
	for k := 1; k < T; k++ {
		if len(groups[k]) == 0 {
			continue
		}
		for i := range dif {
			dif[i] = 1
		}
		for _, q := range groups[k] {
			l, r, v := q[0], q[1], q[2]
			dif[l] = dif[l] * int64(v) % mod3655
			R := ((r-l)/k+1)*k + l
			dif[R] = dif[R] * int64(pow3655(int64(v), mod3655-2)) % mod3655
		}

		for i := k; i < n; i++ {
			dif[i] = dif[i] * dif[i-k] % mod3655
		}
		for i := 0; i < n; i++ {
			nums[i] = int((int64(nums[i]) * dif[i]) % mod3655)
		}
	}

	res := 0
	for _, x := range nums {
		res ^= x
	}
	return res
}
