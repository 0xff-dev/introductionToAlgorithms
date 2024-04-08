package leetcode

func countStudents(students []int, sandwiches []int) int {
	// 找0，1的差值
	// 0:2, 1:4
	ones, zeros := 0, 0
	for _, n := range students {
		if n == 1 {
			ones++
			continue
		}
		zeros++
	}

	for _, n := range sandwiches {
		if n == 1 {
			if ones == 0 {
				break
			}
			ones--
			continue
		}
		if zeros == 0 {
			break
		}
		zeros--
	}
	return ones + zeros
}
