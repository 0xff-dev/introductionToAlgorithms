package leetcode

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	maxIdx, _ := max3(height)
	leftWater, rightWater := 0, 0
	leftNow := 0
	for leftIdx := 1; leftIdx < maxIdx; leftIdx++ {
		if height[leftIdx] >= height[leftNow] {
			leftNow = leftIdx
			continue
		}
		leftWater += height[leftNow] - height[leftIdx]
	}
	rightNow := len(height) - 1
	for rightIdx := len(height) - 1; rightIdx > maxIdx; rightIdx-- {
		if height[rightIdx] > height[rightNow] {
			rightNow = rightIdx
			continue
		}
		rightWater += height[rightNow] - height[rightIdx]
	}
	return leftWater + rightWater
}

func max3(args []int) (int, int) {
	i := 0
	m := args[i]
	for idx, height := range args {
		if height > m {
			i = idx
			m = height
		}
	}
	return i, m
}
