package leetcode

func doesValidArrayExist(derived []int) bool {
	ans := 0
	for _, n := range derived {
		ans ^= n
	}
	return ans == 0
}
