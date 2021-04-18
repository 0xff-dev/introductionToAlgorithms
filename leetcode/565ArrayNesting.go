package leetcode

func arrayNesting(nums []int) int {
	r := 0
	for idx := 0; idx < len(nums); idx++ {
		if nums[idx] == -1 {
			continue
		}
		c, tmp := 0, idx
		for nums[tmp] != -1 {
			c++
			next := nums[tmp]
			nums[tmp] = -1
			tmp = next
		}
		if r < c {
			r = c
		}
	}
	return r
}
