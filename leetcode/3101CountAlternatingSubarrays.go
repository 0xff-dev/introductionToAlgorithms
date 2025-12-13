package leetcode

func countAlternatingSubarrays(nums []int) int64 {
	l := len(nums)
	ret := int64(l)
	cnt := 0
	start, end := 0, 1
	for ; end < l; end++ {
		if nums[end] != nums[end-1] {
			continue
		}

		length := end - start
		cnt = (length - 1) * length / 2
		ret += int64(cnt)

		start = end
	}

	length := end - start
	cnt = (length - 1) * length / 2
	ret += int64(cnt)

	return ret
}
