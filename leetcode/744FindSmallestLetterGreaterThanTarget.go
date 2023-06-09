package leetcode

/*
You are given an array of characters letters that is sorted in non-decreasing order, and a character target. There are at least two different characters in letters.

Return the smallest character in letters that is lexicographically greater than target. If such a character does not exist, return the first character in letters.
*/

// letters = ["c","f","j"], target = "a"
func nextGreatestLetter(letters []byte, target byte) byte {
	x := uint8(0)
	for i := 0; i < len(letters); i++ {
		if letters[i] > target {
			if x == 0 || letters[i] < x {
				x = letters[i]
			}
		}
	}
	if x == 0 {
		return letters[0]
	}
	return x
}
