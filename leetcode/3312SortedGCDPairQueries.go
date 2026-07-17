package leetcode

func gcdValues(nums []int, queries []int64) []int {
	m := nums[0]
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	cnt := make([]int64, m+1)
	for _, num := range nums {
		cnt[num]++
	}
	for i := 1; i <= m; i++ {
		for j := i * 2; j <= m; j += i {
			cnt[i] += cnt[j]
		}
	}
	for i := 1; i <= m; i++ {
		cnt[i] = cnt[i] * (cnt[i] - 1) / 2
	}
	for i := m; i >= 1; i-- {
		for j := i * 2; j <= m; j += i {
			cnt[i] -= cnt[j]
		}
	}
	for i := 1; i <= m; i++ {
		cnt[i] += cnt[i-1]
	}
	ans := make([]int, len(queries))
	for k, q := range queries {
		q++
		left, right := 1, m
		for left < right {
			mid := (left + right) / 2
			if cnt[mid] >= q {
				right = mid
			} else {
				left = mid + 1
			}
		}
		ans[k] = left
	}
	return ans
}
