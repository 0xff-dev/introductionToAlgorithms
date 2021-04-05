package leetcode

// n = 2
func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	pow2 := pow(n)
	r := make([]int, pow2[n-1]*2)
	r[0], r[1] = 0, 1
	for idx := 2; idx <= n; idx++ {
		index := pow2[idx-1]
		for i, j := index-1, index; i >= 0; i, j = i-1, j+1 {
			r[j] = index + r[i]
		}
	}
	return r
}

func pow(x int) []int {
	result := make([]int, x)
	base := 1
	for idx := 1; idx <= x; idx++ {
		result[idx-1] = base
		base *= 2
	}
	return result
}
