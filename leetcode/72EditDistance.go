package leetcode

func minDistance(word1 string, word2 string) int {
	lenW1, lenW2 := len(word1), len(word2)
	if lenW2 == 0 {
		return lenW1
	}
	if lenW1 == 0 {
		return lenW2
	}
	array := make([][]int, 2)
	array[0], array[1] = make([]int, lenW2+1), make([]int, lenW2+1)
	for idx := 1; idx <= lenW2; idx++ {
		array[0][idx] = idx
	}
	array[1][0] = 1

	loopIdx := 1
	for row := 1; row < lenW1+1; row++ {
		for col := 1; col < lenW2+1; col++ {
			if word1[row-1] == word2[col-1] {
				array[loopIdx][col] = array[1-loopIdx][col-1]
			} else {
				_min := array[1-loopIdx][col-1]
				if array[loopIdx][col-1] < _min {
					_min = array[loopIdx][col-1]
				}
				if array[1-loopIdx][col] < _min {
					_min = array[1-loopIdx][col]
				}
				array[loopIdx][col] = _min + 1
			}
		}
		loopIdx = 1 - loopIdx
		array[loopIdx][0] = array[1-loopIdx][0] + 1
	}

	return array[1-loopIdx][lenW2]
}
