package leetcode

import "sort"

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	ans := 0
	for idx := 0; idx < len(seats); idx++ {
		moves := students[idx] - seats[idx]
		if moves < 0 {
			moves = -moves
		}
		ans += moves
	}
	return ans
}
