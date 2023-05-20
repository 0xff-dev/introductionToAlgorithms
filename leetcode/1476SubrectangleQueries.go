package leetcode

type SubrectangleQueries struct {
	rec [][]int
}

func Constructor1476(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rec: rectangle}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for r := row1; r <= row2; r++ {
		for c := col1; c <= col2; c++ {
			this.rec[r][c] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.rec[row][col]
}
