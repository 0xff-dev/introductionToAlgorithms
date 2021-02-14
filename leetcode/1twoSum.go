package leetcode

func twoSum(nums []int, target int) []int {
	existsNum := make(map[int][]int)
	for idx, n := range nums {
		if _, ok := existsNum[n]; !ok {
			existsNum[n] = make([]int, 0)
		}
		existsNum[n] = append(existsNum[n], idx)
	}

	for idx, n := range nums {
		if result, ok := existsNum[target-n]; ok {
			if (target-n == n && len(result) > 1) || target-n != n {
				return getResultIndex(result, idx, n, target-n)
			}
		}
	}
	return []int{}
}

func getResultIndex(indexs []int, nowIdx, n, diff int) []int {
	if diff != n {
		return []int{nowIdx, indexs[0]}
	}
	for _, i := range indexs {
		if i != nowIdx {
			return []int{nowIdx, i}
		}
	}
	return []int{}
}
