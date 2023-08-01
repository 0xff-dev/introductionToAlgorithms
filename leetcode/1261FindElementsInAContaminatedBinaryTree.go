package leetcode

type FindElements struct {
	cache map[int]struct{}
}

func (f *FindElements) dfs(root *TreeNode) {
	if root == nil {
		return
	}
	f.cache[root.Val] = struct{}{}
	if root.Left != nil {
		root.Left.Val = 2*root.Val + 1
		f.dfs(root.Left)
	}
	if root.Right != nil {
		root.Right.Val = 2*root.Val + 2
		f.dfs(root.Right)
	}
}

func Constructor1261(root *TreeNode) FindElements {
	o := FindElements{cache: make(map[int]struct{})}
	// add root
	root.Val = 0
	o.dfs(root)

	return o
}

func (this *FindElements) Find(target int) bool {
	_, ok := this.cache[target]
	return ok
}
