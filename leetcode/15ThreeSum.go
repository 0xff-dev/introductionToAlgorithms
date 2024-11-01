package leetcode

func key15(a, b, c int) [3]int {
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}
	return [3]int{a, b, c}
}

func threeSum(nums []int) [][]int {
	used := make(map[[3]int]struct{})
	nc := make(map[int]int)
	for _, n := range nums {
		nc[n]++
	}
	l := len(nums)
	ans := make([][]int, 0)

	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			left := 0 - (nums[i] + nums[j])
			v, ok := nc[left]
			if !ok {
				continue
			}
			need := 1
			if nums[i] == left {
				need++
			}
			if nums[j] == left {
				need++
			}
			if v >= need {
				key := key15(nums[i], nums[j], left)
				if _, ok := used[key]; ok {
					continue
				}
				ans = append(ans, []int{nums[i], nums[j], left})
				used[key] = struct{}{}
			}
		}
	}

	return ans
}
