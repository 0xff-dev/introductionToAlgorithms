package leetcode

func generateMatrix(n int) [][]int {
	r := make([][]int, n)
	for idx := 0; idx < n; idx++ {
		r[idx] = make([]int, n)
	}
	start := 1
	loop := 0
	for start <= n*n {
		x, y := loop, loop
		for ; y < n-loop; y, start = y+1, start+1 {
			r[x][y] = start
		}
		x, y = x+1, y-1
		for ; x < n-loop; x, start = x+1, start+1 {
			r[x][y] = start
		}
		x, y = x-1, y-1
		for ; y >= loop; y, start = y-1, start+1 {
			r[x][y] = start
		}
		x, y = x-1, y+1
		for ; x > loop; x, start = x-1, start+1 {
			r[x][y] = start
		}
		loop++
	}
	return r
}
