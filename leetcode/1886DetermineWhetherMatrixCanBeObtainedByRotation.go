package leetcode

func findRotation(mat [][]int, target [][]int) bool {
	if equalTarget(mat, target) {
		return true
	}
	for i := 0; i < 3; i++ {
		innerRotate(mat)
		if equalTarget(mat, target) {
			return true
		}
	}
	return false
}

func innerRotate(mat [][]int) {
	n := len(mat)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for s, e := 0, n-1; s < e; s, e = s+1, e-1 {
			mat[i][s], mat[i][e] = mat[i][e], mat[i][s]
		}
	}
}

func equalTarget(mat [][]int, target [][]int) bool {
	for row := 0; row < len(mat); row++ {
		for col := 0; col < len(mat[0]); col++ {
			if mat[row][col] != target[row][col] {
				return false
			}
		}
	}
	return true
}
