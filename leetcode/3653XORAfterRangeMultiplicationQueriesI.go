package leetcode

const mod3653 = 1000000007

func xorAfterQueries(nums []int, queries [][]int) int {
	ret := 0
	for _, q := range queries {
		for i := q[0]; i <= q[1]; i += q[2] {
			nums[i] = (nums[i] * q[3]) % mod3653
		}
	}
	for _, n := range nums {
		ret ^= n
	}
	return ret
}
