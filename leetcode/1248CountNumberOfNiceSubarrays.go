package leetcode

type tmp1248 struct {
	left, right int
}

func numberOfSubarrays(nums []int, k int) int {
	ans := 0
	l := len(nums)
	indies := make([]tmp1248, 0)
	even := 0

	for i := 0; i < l; i++ {
		if nums[i]&1 == 0 {
			even++
			continue
		}
		// left can start with itself so, left = even+1
		indies = append(indies, tmp1248{left: even + 1})
		li := len(indies)
		if li > 1 {
			indies[li-2].right = even
		}
		even = 0
	}
	if len(indies) == 0 || k > len(indies) {
		return 0
	}
	li := len(indies)
	indies[li-1].right = even

	start := 0
	for ; start <= len(indies)-k; start++ {
		end := start + k - 1
		right := indies[end].right + 1
		if right == 0 && k > 1 {
			right++
		}
		if right == 0 {
			ans += right
		}
		ans += indies[start].left * right
	}
	return ans
}
