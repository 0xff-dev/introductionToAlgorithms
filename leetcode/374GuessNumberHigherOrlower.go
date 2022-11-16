package leetcode

func guess(int) int
func guessNumber(n int) int {
	start, end := 1, n
	for {
		mid := (end-start)/2 + start
		r := guess(mid)
		if r == 0 {
			return mid
		}
		if r == 1 {
			start = mid + 1
			continue
		}
		end = mid - 1
	}
}
