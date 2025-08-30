package leetcode

func isValidSudoku(board [][]byte) bool {
	var rows [9]map[byte]bool
	var cols [9]map[byte]bool
	var boxes [9]map[byte]bool

	// 初始化每个集合
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num != '.' {
				// 计算对应的box索引
				boxIndex := (i/3)*3 + j/3

				// 检查行
				if rows[i][num] {
					return false
				}
				rows[i][num] = true

				// 检查列
				if cols[j][num] {
					return false
				}
				cols[j][num] = true

				// 检查3x3盒子
				if boxes[boxIndex][num] {
					return false
				}
				boxes[boxIndex][num] = true
			}
		}
	}
	return true
}
