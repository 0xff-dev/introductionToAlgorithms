package leetcode

var digitsLetterMap = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	result := make([]byte, 0)
	r := make([]string, 0)
	if digits == "" {
		return r
	}
	digitAux(digits, 0, &result, &r)
	return r
}

func digitAux(digits string, index int, result *[]byte, r *[]string) {
	if index == len(digits) {
		*r = append(*r, string(*result))
		return
	}

	for _, _byte := range digitsLetterMap[digits[index]] {
		*result = append(*result, _byte)
		digitAux(digits, index+1, result, r)
		*result = (*result)[:len(*result)-1]
	}
}
