package leetcode

func canJump(nums []int) bool {
	//reachable := make([]uint, len(nums))
	//reachable[len(nums)-1] = 1
	//return canJumpFromIndex(0, reachable, nums)
	target := len(nums) - 1
	for idx := len(nums) - 2; idx >= 0; idx-- {
		if idx+nums[idx] >= target {
			// 从这个目标可以到达
			target = idx
		}
	}
	return target == 0
}

//func jumpAux(nums []int, idx int, reachable *bool) {
//	if idx >= len(nums)-1 {
//		*reachable = true
//		return
//	}
//	if !(*reachable) {
//		for step := 1; step <= nums[idx]; step++ {
//			jumpAux(nums, idx+step, reachable)
//		}
//	}
//}

//func canJumpFromIndex(pos int, reachable []uint, nums []int) bool {
//	if reachable[pos] != 0 {
//		if reachable[pos] == 1 {
//			return true
//		}
//		return false
//	}
//
//	nextPos := pos + nums[pos]
//	if nextPos >= len(nums)-1 {
//		nextPos = len(nums)-1
//	}
//	for np := pos+1; np <= nextPos; np++ {
//		if canJumpFromIndex(np, reachable, nums) {
//			reachable[pos] = 1
//			return true
//		}
//	}
//	reachable[pos] = 2
//	return false
//}
