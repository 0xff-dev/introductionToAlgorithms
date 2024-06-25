package leetcode

/*
找公共祖先，么得，超时了。。
func findCommonFather(root *TreeNode, startValue int, destValue int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == startValue || root.Val == destValue {
		return root
	}
	left := findCommonFather(root.Left, startValue, destValue)
	right := findCommonFather(root.Right, startValue, destValue)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

type tmpPath2096 struct {
	node *TreeNode
	path string
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	commonFather := findCommonFather(root, startValue, destValue)
	if commonFather == nil {
		return ""
	}
	buf := strings.Builder{}
	queue := []*TreeNode{commonFather}
	steps := -1
	level := 0
	for len(queue) > 0 {
		nq := make([]*TreeNode, 0)
		for _, node := range queue {
			if node.Val == startValue {
				steps = level
				break
			}
			if node.Left != nil {
				nq = append(nq, node.Left)
			}
			if node.Right != nil {
				nq = append(nq, node.Right)
			}
		}
		if steps != -1 {
			break
		}
		level++
		queue = nq
	}
	for i := 0; i < steps; i++ {
		buf.WriteByte('U')
	}

	q := []tmpPath2096{{node: commonFather, path: ""}}
	rightSteps := "-"
	for len(q) > 0 {
		nq := make([]tmpPath2096, 0)
		for _, node := range q {
			if node.node.Val == destValue {
				rightSteps = node.path
				break
			}
			if node.node.Left != nil {
				nq = append(nq, tmpPath2096{node: node.node.Left, path: node.path + "L"})
			}
			if node.node.Right != nil {
				nq = append(nq, tmpPath2096{node: node.node.Right, path: node.path + "R"})
			}
		}
		if rightSteps != "-" {
			break
		}
		q = nq
	}
	buf.WriteString(rightSteps)
	return buf.String()
}
*/

type path2096 struct {
	node *TreeNode
	path byte
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	var dfs func(*TreeNode, *[]int, *[]byte, int) bool
	dfs = func(tree *TreeNode, nums *[]int, dirs *[]byte, value int) bool {
		if tree == nil {
			return false
		}
		if tree.Val == value {
			return true
		}
		if tree.Left != nil {
			*nums = append(*nums, tree.Left.Val)
			*dirs = append(*dirs, 'L')
			if dfs(tree.Left, nums, dirs, value) {
				return true
			}
			*nums = (*nums)[:len(*nums)-1]
			*dirs = (*dirs)[:len(*dirs)-1]

		}
		if tree.Right != nil {
			*nums = append(*nums, tree.Right.Val)
			*dirs = append(*dirs, 'R')
			if dfs(tree.Right, nums, dirs, value) {
				return true
			}
			*nums = (*nums)[:len(*nums)-1]
			*dirs = (*dirs)[:len(*dirs)-1]
		}
		return false
	}
	nums1, nums2 := []int{}, []int{}
	dirs1, dirs2 := []byte{}, []byte{}
	_ = dfs(root, &nums1, &dirs1, startValue)
	_ = dfs(root, &nums2, &dirs2, destValue)
	//fmt.Printf("--- nums1: %v, dirs1: %s\n", nums1, dirs1)
	//fmt.Printf("--- nums2: %v, dirs2: %s\n", nums2, dirs2)

	i := 0
	for ; i < len(nums1) && i < len(nums2) && nums1[i] == nums2[i]; i++ {
	}
	for i1 := i; i1 < len(dirs1); i1++ {
		dirs1[i1] = 'U'
	}
	return string(dirs1[i:]) + string(dirs2[i:])
}
