package leetcode

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	n := len(nums1)

	// 1. 找出最大可能差值，用于开辟频次数组
	maxDiff := 0
	diffs := make([]int, n)
	for i := 0; i < n; i++ {
		d := nums1[i] - nums2[i]
		if d < 0 {
			d = -d
		}
		diffs[i] = d
		if d > maxDiff {
			maxDiff = d
		}
	}

	// 如果最大差值已经是 0，直接返回 0
	if maxDiff == 0 {
		return 0
	}

	// 2. 统计每个差值的频次
	freq := make([]int, maxDiff+2)
	for _, d := range diffs {
		freq[d]++
	}

	// 总可用操作次数
	k := k1 + k2

	// 3. 从最大差值开始，从高到低贪心削减
	for d := maxDiff; d > 0; d-- {
		if freq[d] == 0 {
			continue
		}

		// 当前差值 d 的个数与可用次数 k 取最小值
		take := k
		if freq[d] < take {
			take = freq[d]
		}

		k -= take
		freq[d] -= take
		freq[d-1] += take // 削减后差值变成 d-1

		if k == 0 {
			break
		}
	}

	// 4. 计算最终的平方和
	var ans int64 = 0
	for d := 1; d <= maxDiff; d++ {
		if freq[d] > 0 {
			ans += int64(d) * int64(d) * int64(freq[d])
		}
	}

	return ans
}
