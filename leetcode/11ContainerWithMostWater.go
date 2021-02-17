package leetcode

func maxArea(height []int) int {
	start, end := 0, len(height)-1
	max := 0
	for start < end {
		tmp := compare(height[start], height[end], false) * (end - start)
		max = compare(max, tmp, true)
		if height[start] < height[end] {
			start++
			continue
		}
		end--
	}
	return max
}

func compare(a, b int, max bool) int {
	if max {
		if a > b {
			return a
		}
		return b
	}
	if a < b {
		return a
	}
	return b
}
