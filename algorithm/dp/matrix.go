package dp

import "fmt"

const (
	keyFormat = `%d-%d`
	inf       = 0xffff
)

type matrix struct {
	Row int
	Col int
}

func MinMatrixCalculateCount(matrixs []matrix) map[string]int {
	allMatrix := len(matrixs)
	if allMatrix == 0 || allMatrix == 1 {
		return nil
	}
	minResult := map[string]int{}
	result := map[string]int{}
	for index := 0; index < allMatrix; index++ {
		minResult[fmt.Sprintf(keyFormat, index, index)] = 0
	}
	// 外层循环在计算长度为length时候, 从i-j最小的计算值
	for length := 2; length <= allMatrix; length++ {
		// 计算长度为length的时候，可以作为开始矩阵的位置
		for start := 0; start < allMatrix-length+1; start++ {
			// 长度为length的结束矩阵
			end := start + length - 1
			key := fmt.Sprintf(keyFormat, start, end)
			minResult[key] = inf
			for k := start; k <= end-1; k++ {
				// start-k, k+1-end + mid
				key1, key2 := fmt.Sprintf(keyFormat, start, k), fmt.Sprintf(keyFormat, k+1, end)
				min := minResult[key1] + minResult[key2] + matrixs[start].Row*matrixs[k].Col*matrixs[end].Col
				if min < minResult[key] {
					minResult[key] = min
					result[key] = k
				}
			}
		}
	}
	findPath(result, 0, 5)
	return minResult
}

func findPath(result map[string]int, i, j int) {
	if i == j {
		fmt.Print("A")
	} else {
		fmt.Print("(")
		findPath(result, i, result[fmt.Sprintf(keyFormat, i, j)])
		findPath(result, result[fmt.Sprintf(keyFormat, i, j)]+1, j)
		fmt.Print(")")
	}
}
