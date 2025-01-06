package leetcode

func numTimesAllBlue(flips []int) int {
	ans := 0
	sum, cur := 0, 0
	index := 1
	for _, f := range flips {
		sum += index
		index++
		cur += f
		if sum == cur {
			ans++
		}
	}
	return ans
}
