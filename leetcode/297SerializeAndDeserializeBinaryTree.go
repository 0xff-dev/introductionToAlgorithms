package leetcode

import "fmt"

type Codec297 struct {
}

func Constructor297() Codec297 {
	return Codec297{}
}

// Serializes a tree to a single string.
func (this *Codec297) serialize(root *TreeNode) string {
	if root == nil {
		return "{}"
	}
	cur := fmt.Sprintf("{%d", root.Val)
	cur += this.serialize(root.Left)
	cur += this.serialize(root.Right)
	return cur + "}"
}

// Deserializes your encoded data to tree.
func (this *Codec297) deserialize(data string) *TreeNode {
	index := 0
	return this.helper(&index, data)
}

func (this *Codec297) helper(index *int, data string) *TreeNode {
	if *index == len(data) {
		return nil
	}
	if data[*index] == '}' {
		return nil
	}
	*index++
	neg := false
	if data[*index] == '-' {
		*index++
		neg = true
	}
	v := 0
	i := *index
	for ; i < len(data); i++ {
		if !(data[i] >= '0' && data[i] <= '9') {
			break
		}
		v = v*10 + int(data[i]-'0')
	}

	if i == *index {
		return nil
	}

	if neg {
		v = -v
	}
	*index = i
	n := &TreeNode{Val: v}
	n.Left = this.helper(index, data)
	*index++

	n.Right = this.helper(index, data)
	*index++
	return n
}
