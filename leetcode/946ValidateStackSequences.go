package leetcode

func validateStackSequences(pushed []int, popped []int) bool {
	length := len(popped)
	if length == 0 {
		return true
	}
	popTopIdx, pushIdx := 0, 0
	pushStack := make([]int, length)
	for idx := 0; idx < length; idx++ {
		pushStack[pushIdx] = pushed[idx]
		for popTopIdx < length && pushIdx >= 0 && popped[popTopIdx] == pushStack[pushIdx] {
			pushIdx, popTopIdx = pushIdx-1, popTopIdx+1
		}
		pushIdx++
	}
	if pushIdx == 0 {
		return true
	}
	return false
}
