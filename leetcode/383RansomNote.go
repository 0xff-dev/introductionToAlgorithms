package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	chars := make([]int, 26)
	for _, b := range magazine {
		chars[b-'a']++
	}
	for _, b := range ransomNote {
		key := b - 'a'
		if chars[key] <= 0 {
			return false
		}
		chars[key]--
	}
	return true
}
