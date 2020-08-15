package offer

func FirstNotRepeatingChar(chars string) byte {
	checkMap := map[byte]int{}
	for _, char := range []byte(chars) {
		if _, ok := checkMap[char]; !ok {
			checkMap[char] = 0
		}
		checkMap[char]++
	}
	for _, char := range []byte(chars) {
		if val, ok := checkMap[char]; ok && val == 1 {
			return char
		}
	}
	return ' '
}
