package leetcode

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	evenSum := 0
	for _, n := range nums {
		if n&1 == 0 {
			evenSum += n
		}
	}

	for i, query := range queries {
		idx, val := query[1], query[0]
		j1, j2 := nums[idx]&1 == 0, val&1 == 0
		if j1 && j2 {
			evenSum += val
		} else if !j1 && !j2 {
			evenSum += nums[idx] + val
		} else if j1 {
			evenSum -= nums[idx]
		}
		nums[idx] += val
		ans[i] = evenSum
	}
	return ans
}
