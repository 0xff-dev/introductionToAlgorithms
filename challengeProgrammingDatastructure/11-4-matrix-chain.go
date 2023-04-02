package challengeprogrammingdatastructure

type Matrix struct {
	row, col int
}

func MatrixChainMulti(matrixs []Matrix) int {
	if len(matrixs) < 2 {
		return 0
	}
	l := len(matrixs)
	if l < 2 {
		return 0
	}
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}

	// 3:5, 5:4 两个矩阵相乘，需要3*5*4次计算
	for ml := 2; ml <= l; ml++ {
		for start := 0; start <= l-ml; start++ {
			end := start + ml - 1
			dp[start][end] = -1
			for k := start; k < end; k++ {
				cal := matrixs[start].row * matrixs[k].col * matrixs[end].col
				cal += dp[start][k] + dp[k+1][end]
				if dp[start][end] == -1 || dp[start][end] > cal {
					dp[start][end] = cal
				}
			}
		}
	}
	return dp[0][l-1]
}
