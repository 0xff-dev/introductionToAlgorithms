package leetcode

func closestTarget(words []string, target string, startIndex int) int {
	if words[startIndex] == target {
		return 0
	}
	size := len(words)
	steps := 0
	left, right := -1, -1
	for i := (startIndex - 1 + size) % size; i != startIndex; i = (i - 1 + size) % size {
		steps++
		if words[i] == target {
			left = steps
			break
		}
	}
	steps = 0
	for i := (startIndex + 1) % size; i != startIndex; i = (i + 1) % size {
		steps++
		if words[i] == target {
			right = steps
			break
		}
	}
	if left == -1 && right == -1 {
		return -1
	}
	return min(left, right)
}
