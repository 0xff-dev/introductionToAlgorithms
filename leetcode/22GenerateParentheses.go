package leetcode

func generateParenthesis(n int) []string {
	r := make([]string, 0)
	generateParenthesisAux(n, n, "", &r)
	return r
}

func generateParenthesisAux(left, right int, now string, r *[]string) {
	if left == 0 && right == 0 {
		*r = append(*r, now)
		return
	}

	if left > 0 {
		generateParenthesisAux(left-1, right, now+"(", r)
	}
	if right > left {
		generateParenthesisAux(left, right-1, now+")", r)
	}
}
