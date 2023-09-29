package leetcode

func brokenCalc(startValue int, target int) int {
	// 3 1
	/*
		if target <= startValue {
			return startValue - target
		}
		queue := [][]int{{startValue, 0}}
		for len(queue) > 0 {
			nq := make([][]int, 0)
			for _, item := range queue {
				if item[0] == target {
					return item[1]
				}
				if item[0] > target {
					a := item[0] - 1
					nq = append(nq, []int{a, item[1] + 1})
					continue
				}
				a := item[0] * 2
				b := item[0] - 1
				nq = append(nq, []int{a, item[1] + 1}, []int{b, item[1] + 1})
			}
			queue = nq
		}
	*/
	ans := 0
	for target > startValue {
		ans++
		if target&1 == 1 {
			target++
			continue
		}
		target /= 2
	}
	return ans + startValue - target
}
