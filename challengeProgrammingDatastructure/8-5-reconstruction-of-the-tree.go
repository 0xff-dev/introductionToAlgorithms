package challengeprogrammingdatastructure

func ReconstructTree(pre, in []int) []int {
	if len(pre) == 0 {
		return []int{}
	}
	if len(pre) == 1 {
		return []int{pre[0]}
	}
	root := pre[0]
	index := 0
	for ; index < len(in) && in[index] != root; index++ {
	}
	ans := make([]int, 0)
	if index > 0 {
		ans = append(ans, ReconstructTree(pre[1:index+1], in[:index])...)
	}
	if index < len(in)-1 {
		ans = append(ans, ReconstructTree(pre[index+1:], in[index+1:])...)
	}
	ans = append(ans, root)
	return ans
}
