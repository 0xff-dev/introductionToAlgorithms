package leetcode

import (
	"sort"
)

type TopVotedCandidate struct {
	matrix [][]int
	mm     []int
	times  []int
}

func Constructor911(persons []int, times []int) TopVotedCandidate {
	l := len(times)

	matrix := make([][]int, l)
	for i := 0; i < l; i++ {
		matrix[i] = make([]int, l)
	}
	// 记录当前时间票数最多的人
	mm := make([]int, l)
	voteTime := make([]int, l)
	for i := range l {
		voteTime[i] = -1
	}
	// 按时间统计
	matrix[0][persons[0]]++
	mm[0] = persons[0]
	voteTime[persons[0]] = times[0]

	for i := 1; i < l; i++ {
		matrix[i][persons[i]]++
		voteTime[persons[i]] = times[i]
		selectuser := 0
		maxvalue := 0
		for c := range matrix[i] {
			matrix[i][c] += matrix[i-1][c]
			if matrix[i][c] == maxvalue && maxvalue != 0 {
				if voteTime[c] > voteTime[selectuser] {
					selectuser = c
				}
			}

			if matrix[i][c] > maxvalue {
				selectuser = c
				maxvalue = matrix[i][c]
			}
		}
		mm[i] = selectuser
	}
	return TopVotedCandidate{matrix: matrix, times: times, mm: mm}
}

func (this *TopVotedCandidate) Q(t int) int {
	l := len(this.times)
	index := sort.Search(l, func(i int) bool {
		return this.times[i] >= t
	})
	if index == l || this.times[index] != t {
		index--
	}
	return this.mm[index]
}
