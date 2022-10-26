package leetcode

// 23,2,4,6,7
// 23,25,29,35,42
// 23  2  4  6  7
/*
func checkSubarraySum(nums []int, k int) bool {

	for idx := 1; idx < len(nums); idx++ {
		nums[idx] += nums[idx-1]
		if nums[idx]%k == 0 {
			return true
		}
		for i := idx - 2; i >= 0; i-- {
			if target := nums[idx] - nums[i]; target%k == 0 {
				return true
			}
		}
	}

	return false
}
*/
// 用余数计算是真的骚啊
// 5%6=5,    11%6=5 如果出现了两次余数是5，那么就说明在整个加和的过程中出现了n*k 所以就可以判断
// 出是否有k的倍数
// 5, 6
// 5, 2, 4
func checkSubarraySum(nums []int, k int) bool {

	checker := map[int]int{0: 0}
	sum := 0
	for idx := 0; idx < len(nums); idx++ {
		sum += nums[idx]
		i, ok := checker[sum%k]
		if !ok {
			checker[sum%k] = idx + 1
			continue
		}
		if i < idx {
			// > 保证至少有两个数字
			return true
		}
	}
	return false
}
