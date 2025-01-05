package leetcode

func shiftingLetters(str string, shifts [][]int) string {
	n := len(str)
	diffArray := make([]int, n)

	for _, shift := range shifts {
		if shift[2] == 1 {
			diffArray[shift[0]]++
			if shift[1]+1 < n {
				diffArray[shift[1]+1]--
			}
		} else {
			diffArray[shift[0]]--
			if shift[1]+1 < n {
				diffArray[shift[1]+1]++
			}
		}
	}

	result := make([]byte, n)
	numberOfShifts := 0

	for i := 0; i < n; i++ {
		numberOfShifts = (numberOfShifts + diffArray[i]) % 26
		if numberOfShifts < 0 {
			numberOfShifts += 26
		}

		result[i] = 'a' + (str[i]-'a'+byte(numberOfShifts))%26
	}

	return string(result)
}
