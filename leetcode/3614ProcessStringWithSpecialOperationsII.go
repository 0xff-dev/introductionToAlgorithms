package leetcode

func processStr3614(s string, k int64) byte {
	var length int64
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '*':
			if length > 0 {
				length--
			}
		case '#':
			length *= 2
		case '%':
			// no length change
		default:
			length++
		}
	}
	if k+1 > length {
		return '.'
	}
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '*':
			length++
		case '#':
			if k+1 > (length+1)/2 {
				k -= length / 2
			}
			length = (length + 1) / 2
		case '%':
			k = length - k - 1
		default:
			if k+1 == length {
				return s[i]
			}
			length--
		}
	}
	return '.'
}
