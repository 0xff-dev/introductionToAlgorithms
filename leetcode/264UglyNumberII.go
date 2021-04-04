package leetcode

func nthUglyNumber(n int) int {
	if n < 0 {
		return 0
	}

	r := make([]int, n)
	p2, p3, p5 := 0, 0, 0
	r[0] = 1
	index := 1
	for index < n {
		_min := r[p2] * 2
		if r[p3]*3 < _min {
			_min = r[p3] * 3
		}
		if r[p5]*5 < _min {
			_min = r[p5] * 5
		}

		r[index] = _min
		for r[p2]*2 <= r[index] {
			p2++
		}
		for r[p3]*3 <= r[index] {
			p3++
		}
		for r[p5]*5 <= r[index] {
			p5++
		}
		index++
	}
	return r[n-1]
}
