package leetcode

func findLucky(arr []int) int {
	bucket := make([]int, 501)
	for _, n := range arr {
		bucket[n]++
	}
	for idx := 500; idx > 0; idx-- {
		if bucket[idx] == idx {
			return idx
		}
	}
	return -1
}
