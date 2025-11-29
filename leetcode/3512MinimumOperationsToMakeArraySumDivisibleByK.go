package leetcode

func minOperations3512(nums []int, k int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum % k
}
