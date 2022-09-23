package leetcode

func distributeCandies(candyType []int) int {
	bucket := make(map[int]int)
	for _, ct := range candyType {
		bucket[ct]++
	}
	mid := len(candyType) / 2
	if len(bucket) >= mid {
		return mid
	}
	return len(bucket)
}
