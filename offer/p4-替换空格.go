package offer

func ReplaceSpace(str string) string {
	spaceCount := 0
	for _, char := range str {
		if char == ' ' {
			spaceCount++
		}
	}
	bytesArray := make([]byte, len(str)+2*spaceCount)
	newIndex := len(bytesArray) - 1
	for index := len(str) - 1; index >= 0; index-- {
		if str[index] != ' ' {
			bytesArray[newIndex] = str[index]
			newIndex--
			continue
		}
		bytesArray[newIndex] = '0'
		bytesArray[newIndex-1] = '2'
		bytesArray[newIndex-2] = '%'
		newIndex -= 3
	}
	return string(bytesArray)
}
