package leetcode

func maxDistance(colors []int) int {
	step, length := len(colors)-1, len(colors)-1
	for ; step > 0; step-- {
		for idx := length; idx >= step; idx-- {
			end := idx - step
			if colors[idx] != colors[end] {
				return step
			}
		}
	}
	return 0
}
