package leetcode

func reverse2946(src []int, start, end int) {
	for ; start < end; start, end = start+1, end-1 {
		src[start], src[end] = src[end], src[start]
	}
}

func areSimilar(mat [][]int, k int) bool {
	m, n := len(mat), len(mat[0])
	k %= n
	if k == 0 {
		return true
	}
	cp := make([][]int, m)
	for i := range m {
		cp[i] = make([]int, n)
		copy(cp[i], mat[i])
	}

	for i := range m {
		step := k
		if i&1 == 0 {
			step = -k
		}
		if step < 0 {
			step += n
		}
		reverse2946(cp[i], 0, n-1)
		reverse2946(cp[i], 0, step-1)
		reverse2946(cp[i], step, n-1)
		for j := range n {
			if cp[i][j] != mat[i][j] {
				return false
			}
		}
	}
	return true
}
