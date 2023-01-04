package leetcode

func minimumRounds(tasks []int) int {
	taskLevel := make(map[int]int)
	for _, t := range tasks {
		taskLevel[t]++
	}
	rounds := 0
	// 2,2,2,3,3,4,4,4,4,4

	for _, count := range taskLevel {
		//fmt.Printf("id=%d --> count=%d", id, count)
		if count == 1 {
			return -1
		}
		if count == 4 {
			rounds += 2
			//fmt.Printf(" reason 4: +2 = %d\n", rounds)
			continue
		}
		cond := count % 3
		if cond == 0 {
			rounds += count / 3
			//fmt.Printf(" reason mod0 +%d = %d\n", count/3, rounds)
			continue
		}
		// 4(2+2), 7(3+2+2), 10(3+3+2+2), 13(3+3+3+2+2)
		if cond == 1 {
			rounds += (count-4)/3 + 2
			//fmt.Printf(" readon mod1 + %d = %d\n", (count-4)/3+2, rounds)
			continue
		}
		// 5(3+2), 8(3+3+2), 11(3+3+3+2)
		rounds += count/3 + 1
		//fmt.Printf(" readon mod2 + %d = %d\n", count/3+1, rounds)
	}
	return rounds
}
