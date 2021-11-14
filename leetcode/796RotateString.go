package leetcode

func rotateString(s, goal string) bool {
	if s == goal {
		return true
	}
	if len(s) != len(goal) {
		return false
	}

	length := len(s)
	for idx := 1; idx < length; idx++ {
		if s[:idx] == goal[length-idx:] && s[idx:] == goal[:length-idx] {
			return true
		}
	}
	return false
}
