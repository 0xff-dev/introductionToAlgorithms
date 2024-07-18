package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	ones := 0
	ans := 0
	for _, n := range nums {
		if n == 1 {
			ones++
			continue
		}
		ans = max(ans, ones)
		ones = 0
	}
	ans = max(ans, ones)
	return ans
}
