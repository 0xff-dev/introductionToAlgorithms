package leetcode

func isOneBitCharacter(bits []int) bool {
	// 1, 0 , 0
	i := 0
	for i < len(bits)-1 {
		if bits[i] == 1 {
			i += 2
			continue
		}
		i++
	}
	return i == len(bits)-1
}
