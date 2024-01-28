package leetcode

/* n^2 勉强过。
func findMaxLength(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	ps := make([]int, len(nums))
	ps[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		ps[i] = ps[i-1] + nums[i]
	}
	// 0 1 2 3 4 5 6 7
	// 0 1 0 1 0 0 1 1
	// 0 1 1 2 2 2 3 4

	// 0, 1
	// 0, 1
	// 0, 1
    // 3个1和4个0
	ans := 0
	for end := len(nums) - 1; end > 0; end-- {
		start := 0
		if end&1 == 0 {
			start = 1
		}
		for ; start < end; start += 2 {
			dl := end - start + 1
            if dl < ans {
                break
            }
			cut := 0
			if start > 0 {
				cut = ps[start-1]
			}
			if ps[end]-cut == dl/2 && dl > ans {
				ans = dl
				break
			}
		}
	}
	return ans
}
*/

func findMaxLength(nums []int) int {
	sum := 0
	indies := make(map[int]int)
	ans := 0
	// 0 0 1 0 0 1 1

	// 0 1 0
	// -1:0
	//
	indies[0] = -1
	for i := 0; i < len(nums); i++ {
		add := 1
		if nums[i] == 0 {
			add = -1
		}
		sum += add
		if idx, ok := indies[sum]; ok {
			if r := i - idx; r > ans {
				ans = r
			}
			continue
		}
		indies[sum] = i
	}
	return ans
}
