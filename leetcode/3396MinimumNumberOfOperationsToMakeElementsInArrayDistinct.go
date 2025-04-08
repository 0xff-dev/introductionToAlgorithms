package leetcode

func minimumOperations3396(nums []int) int {
	cnt := make(map[int]struct{})
	i := len(nums) - 1
	for ; i >= 0; i-- {
		if _, ok := cnt[nums[i]]; ok {
			break
		}
		cnt[nums[i]] = struct{}{}
	}
	op := (i + 1) / 3
	if (i+1)%3 != 0 {
		op++
	}
	return op
}
