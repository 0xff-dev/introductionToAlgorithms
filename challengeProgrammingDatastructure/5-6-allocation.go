package challengeprogrammingdatastructure

// 12, 8, 8 不能计算综合，平均
func checkP(k, p int, weight []int) int {
	i := 0
	for car := 0; car < k; car++ {
		total := 0
		for total+weight[i] <= p {
			total += weight[i]
			i++
			if i == len(weight) {
				return i
			}
		}
	}
	return i
}
func goodsAlloc(k int, weight []int) int {
	total := 0
	for _, w := range weight {
		total += w
	}
	left, right := 0, total
	ans := -1
	for left <= right {
		mid := (right-left)/2 + left
		v := checkP(k, mid, weight)
		if v >= len(weight) {
			ans = mid
			right = mid - 1
			continue
		}
		left = mid + 1
	}
	return ans
}
