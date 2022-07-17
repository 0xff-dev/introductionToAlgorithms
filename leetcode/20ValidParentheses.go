package leetcode

func isValid(s string) bool {
	length := len(s)
	if length&1 == 1 {
		return false
	}
	//brackets := [3]int{}
	//
	//for idx := 0; idx < length; idx++ {
	//	switch s[idx] {
	//	case '(', ')':
	//		// 0
	//		if s[idx] == '(' {
	//			brackets[0]++
	//			continue
	//		}
	//		if brackets[0] == 0 {
	//			return false
	//		}
	//		brackets[0]--
	//
	//	case '{', '}':
	//		if s[idx] == '{' {
	//			brackets[1]++
	//			continue
	//		}
	//		if brackets[1] == 0 {
	//			return false
	//		}
	//		brackets[1]--
	//	case '[', ']':
	//		if s[idx] == '[' {
	//			brackets[2]++
	//			continue
	//		}
	//		if brackets[2] == 0 {
	//			return false
	//		}
	//		brackets[2]--
	//		// 2
	//	}
	//
	//}
	//return brackets[0] == 0 && brackets[1] == 0 && brackets[2] == 0
	pair := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	brackets := make([]byte, length)
	walker := 0
	for idx := 0; idx < length; idx++ {
		if _, ok := pair[s[idx]]; !ok {
			brackets[walker] = s[idx]
			walker++
			continue
		}
		// 2
		if walker == 0 || brackets[walker-1] != pair[s[idx]] {
			return false
		}
		walker--
	}
	return walker == 0
}
