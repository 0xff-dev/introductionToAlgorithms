package leetcode

import (
	"strconv"
	"strings"
)

type Spreadsheet struct {
	rows  int
	sheet [26][]int
}

func Constructor3484(rows int) Spreadsheet {
	sheet := [26][]int{}
	for i := range 26 {
		sheet[i] = make([]int, rows)
	}
	return Spreadsheet{
		rows:  rows,
		sheet: sheet,
	}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	row := int(cell[0] - 'A')
	v, _ := strconv.Atoi(cell[1:])
	this.sheet[row][v-1] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	row := int(cell[0] - 'A')
	v, _ := strconv.Atoi(cell[1:])
	this.sheet[row][v-1] = 0
}

func (this *Spreadsheet) part(str string) int {
	index := 0
	isNumber := true
	if !(str[0] >= '0' && str[0] <= '9') {
		isNumber = false
		index++
	}
	num, _ := strconv.Atoi(str[index:])
	if isNumber {
		return num
	}
	return this.sheet[int(str[0]-'A')][num-1]
}
func (this *Spreadsheet) GetValue(formula string) int {
	left := formula[1:]
	parts := strings.Split(left, "+")
	a, b := this.part(parts[0]), this.part(parts[1])
	return a + b
}

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */
