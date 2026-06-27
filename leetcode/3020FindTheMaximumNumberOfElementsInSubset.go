package leetcode

func maximumLength3020(nums []int) int {
	cnt := make(map[int]int)
	single := make([]int, 0)
	for i := range nums {
		if _, ok := cnt[nums[i]]; !ok {
			single = append(single, nums[i])
		}
		cnt[nums[i]]++
	}

	var (
		ret int = 1
		tmp int = 0
		ele int = 0
	)
	skip := make(map[int]struct{})
	for i := range single {
		if _, ok := skip[single[i]]; ok {
			continue
		}

		ele, tmp = single[i], 0
		if ele == 1 {
			if cnt[1]&1 == 0 {
				cnt[1]--
			}
			ret = max(ret, cnt[1])
			skip[1] = struct{}{}
			continue
		}

		for cnt[ele] >= 2 {
			skip[ele] = struct{}{}
			ele *= ele
			tmp += 2
		}
		if cnt[ele] == 0 {
			tmp--
		} else {
			skip[ele] = struct{}{}
			tmp++
		}
		ret = max(ret, tmp)
	}
	return ret
}
