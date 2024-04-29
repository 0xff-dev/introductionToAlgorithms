package leetcode

func minOperations2997(nums []int, k int) int {
	//我的理解就是，看k的每一个位
	mask := 1
	ans := 0
	for shift := 0; shift < 32; shift++ {
		cur := mask << shift
		kbit := k & cur
		xor := 0
		for _, n := range nums {
			xor ^= (n & cur)
		}
		if xor != kbit {
			ans++
		}
	}
	return ans
}
