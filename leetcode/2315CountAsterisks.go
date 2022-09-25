package leetcode

func countAsterisks(s string) int {
	ans := 0
	// l|*e*et|c**o|*de|
	addFlag := 0
	for _, b := range s {
		if b == '|' {
			addFlag = 1 - addFlag
			continue
		}
		if b == '*' && addFlag&1 == 0 {
			ans++
		}
	}

	return ans
}
