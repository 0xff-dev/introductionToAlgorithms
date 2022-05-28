package leetcode

/*
14: 6
8: 4
123: 12
*/
func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}

	ans := numberOfSteps(num/2) + 1
	if num&1 == 1 {
		ans++
	}

	return ans
}
