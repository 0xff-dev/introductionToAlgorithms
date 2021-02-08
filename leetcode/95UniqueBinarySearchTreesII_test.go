package leetcode

import "testing"

func TestGenerateTrees(t *testing.T) {
	for idx := 1; idx <= 4; idx++ {
		r := generateTrees(idx)
		t.Logf("idx[%d] totla trees: [%d]", idx, len(r))
		for _, tree := range r {
			t.Logf("%v\n", levelOrder(tree))
		}
	}
}
