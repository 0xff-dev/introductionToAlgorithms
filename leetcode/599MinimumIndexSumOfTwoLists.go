package leetcode

func findRestaurant(list1 []string, list2 []string) []string {
	map2 := make(map[string]int)
	for index, str := range list2 {
		map2[str] = index
	}

	ret := make([]string, 0)
	sum, minSum := 0, 2000
	for index, str := range list1 {
		if i, ok := map2[str]; ok {
			sum = index + i
			if sum == minSum {
				ret = append(ret, str)
				continue
			}
			if sum < minSum {
				ret = []string{str}
				minSum = sum
			}
		}
	}
	return ret
}
