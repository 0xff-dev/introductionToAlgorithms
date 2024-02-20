package leetcode

func missingNumber(nums []int) int {
	ans := 0
	for i := 0; i <= len(nums); i++ {
		ans ^= i
	}
	for _, n := range nums {
		ans ^= n
	}
	return ans
}
