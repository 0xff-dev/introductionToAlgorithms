package leetcode

/*
func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	ans := make([]int, len(obstacles))
	for i := 0; i < len(obstacles); i++ {
		ans[i] = 1
		for pre := i - 1; pre >= 0; pre-- {
			if obstacles[i] >= obstacles[pre] && ans[pre]+1 > ans[i] {
				ans[i] = ans[pre] + 1
			}
		}
	}
	return ans
}
*/

func bisectRight(a []int, target, right int) int {
	if right == 0 {
		return 0
	}
	left := 0
	for left < right {
		mid := (right-left)/2 + left
		if a[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	n := len(obstacles)
	l := 0
	ans := make([]int, n)
	lis := make([]int, n)
	for i := 0; i < n; i++ {
		h := obstacles[i]
		x := bisectRight(lis, h, l)
		if x == l {
			l++
		}
		lis[x] = h
		ans[i] = x + 1
	}
	return ans
}
