package leetcode

func maxDepth559(root *MultiChildNode) int {
	if root == nil {
		return 0
	}
	dep := 0
	for _, c := range root.Children {
		dep = max(dep, maxDepth559(c))
	}
	return dep + 1
}
