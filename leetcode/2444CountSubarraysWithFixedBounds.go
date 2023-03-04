package leetcode

// type bp struct {
// 	// include self
// 	index int
// 	left  int64
// 	right int64 //
// }

// // 1,2,3,5
// func countSubarrays(nums []int, minK int, maxK int) int64 {
// 	ans := int64(0)
// 	points := make([]bp, 0)
// 	sum := int64(0)
// 	allEqual := true
// 	p := nums[0]
// 	for i := 0; i < len(nums); i++ {
// 		if allEqual && nums[i] != p {
// 			allEqual = false
// 		}
// 		if nums[i] < maxK && nums[i] > minK {
// 			sum++
// 			continue
// 		}
// 		newEle := false
// 		if nums[i] == maxK || nums[i] == minK {
// 			newEle = true
// 			points = append(points, bp{index: i, left: sum + 1, right: -1})
// 		}
// 		if len(points) > 1 {
// 			if newEle {
// 				if points[len(points)-2].right == -1 {
// 					points[len(points)-2].right = sum
// 				}
// 			} else if points[len(points)-1].right == -1 {
// 				points[len(points)-1].right = sum
// 			}
// 		}
// 		sum = 0
// 	}
// 	if len(points) == 0 {
// 		return 0
// 	}
// 	if points[len(points)-1].right == -1 {
// 		points[len(points)-1].right = sum
// 	}
// 	if allEqual {
// 		a := int64(len(nums)) * int64(len(nums)+1)
// 		return a / 2
// 	}
// 	for i := 0; i < len(points)-1; i++ {
// 		both := minK == maxK
// 		for next := i + 1; next < len(points); next++ {
// 			if points[next].index-points[next-1].index != int(points[next].left) {
// 				break
// 			}
// 			if !both && nums[points[next].index] != nums[points[i].index] {
// 				both = true
// 			}

// 			if both {
// 				right := points[next].right + 1
// 				ans += points[i].left * right
// 			}

// 		}
// 	}
// 	if minK == maxK {
// 		ans += int64(len(points))
// 	}
// 	return ans
// }

// 还的是高人指点啊~~~~~~~~
func countSubarrays(nums []int, minK, maxK int) int64 {
	ans := int64(0)
	preMin, preMax, edge := -1, -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] < minK || nums[i] > maxK {
			edge = i
			continue
		}
		if nums[i] == minK {
			preMin = i
		}
		if nums[i] == maxK {
			preMax = i
		}
		which := preMin
		if preMax < which {
			which = preMax
		}
		if diff := which - edge; diff > 0 {
			ans += int64(diff)
		}
	}
	return ans
}
