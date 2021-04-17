package leetcode

func isHappy(n int) bool {
	visited := make(map[int]struct{})

	for n != 1 {
		n = next(n)
		if _, ok := visited[n]; ok {
			return false
		}
		visited[n] = struct{}{}
	}
	return true
}

func next(n int) int {
	base := 0
	for n > 0 {
		tmp := n % 10
		base += tmp * tmp
		n = n / 10
	}
	return base
}
