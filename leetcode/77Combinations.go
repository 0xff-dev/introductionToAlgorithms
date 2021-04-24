package leetcode

func combine(n int, k int) [][]int {
	r := make([][]int, 0)
	if k == 0 {
		return r
	}
	combineAux(n, k, 1, &r, []int{})
	return r
}

func combineAux(n, deep, now int, r *[][]int, path []int) {
	if deep == 0 {
		dest := make([]int, len(path))
		copy(dest, path)
		*r = append(*r, dest)
		return
	}

	if now <= n {
		combineAux(n, deep-1, now+1, r, append(path, now))
		combineAux(n, deep, now+1, r, path)
	}
}
