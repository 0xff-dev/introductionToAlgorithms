package leetcode

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	width := fx - sx
	if width < 0 {
		width = -width
	}
	height := fy - sy
	if height < 0 {
		height = -height
	}
	if width == 0 && height == 0 && t == 1 {
		return false
	}
	if height > width {
		width = height
	}
	return t >= width
}
