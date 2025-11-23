package leetcode

func maximumXOR(nums []int) int {
	var ret, index int
	l := len(nums)
	mask := 1
	for i := 0; i < 31; i, mask = i+1, mask<<1 {
		index = 0
		for ; index < l; index++ {
			if nums[index]&mask == mask {
				break
			}
		}
		if index != l {
			ret |= mask
		}
	}
	return ret
}
