package leetcode

// a^b^a=b
func decode(encoded []int, first int) []int {
	ans := make([]int, len(encoded)+1)
	ans[0] = first
	for idx, code := range encoded {
		ans[idx+1] = ans[idx] ^ code
	}
	return ans
}
