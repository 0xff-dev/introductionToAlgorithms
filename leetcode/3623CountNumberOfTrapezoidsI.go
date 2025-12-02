package leetcode

const mod3623 = 1000000007

func countTrapezoids(points [][]int) int {
	columns := make(map[int]int)
	for _, p := range points {
		columns[p[1]]++
	}

	pairs := []int{}
	sum, count := 0, 0
	for _, c := range columns {
		if c < 2 {
			continue
		}
		count = c * (c - 1) / 2
		sum = (sum + count) % mod3623
		pairs = append(pairs, count)
	}
	var ret int
	for _, pair := range pairs {
		ret = (ret + pair*((sum-pair+mod3623)%mod3623)) % mod3623
	}

	// 使用模逆元乘以 1/2
	inv2 := (mod3623 + 1) / 2
	return (ret * inv2) % mod3623
}
