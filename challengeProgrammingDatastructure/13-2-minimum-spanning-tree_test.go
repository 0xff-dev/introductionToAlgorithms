package challengeprogrammingdatastructure

import "testing"

func TestMinimumSpanningTree(t *testing.T) {
	graps := [][]int{
		{-1, 2, 3, 1, -1},
		{2, -1, -1, 4, -1},
		{3, -1, -1, 1, 1},
		{1, 4, 1, -1, 3},
		{-1, -1, 1, 3, -1},
	}
	if ans1, ans2 := MinimumSpanningTree(5, graps); ans1 != 5 || ans2 != 5 {
		t.Fatalf("expect 5 get ans1 %d, ans2 %d", ans1, ans2)
	}

	// a,b,c,d,e,f,g,h,i
	graps = [][]int{
		// 0,1,2,3,4,5,6,7,8
		//a  b . c . d . e . f . g . h . i
		{-1, 4, -1, -1, -1, -1, -1, 8, -1},
		{4, -1, 8, -1, -1, -1, -1, 11, -1},
		{-1, 8, -1, 7, -1, 4, -1, -1, 2},
		{-1, -1, 7, -1, 9, 14, -1, -1, -1},
		{-1, -1, -1, 9, -1, 10, -1, -1, -1},
		{-1, -1, 4, 14, 10, -1, 2, -1, -1},
		{-1, -1, -1, -1, -1, 2, -1, 1, 6},
		{8, 11, -1, -1, -1, -1, 1, -1, 7},
		{-1, -1, 2, -1, -1, -1, 6, 7, -1},
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if graps[i][j] != graps[j][i] {
				t.Fatalf("ij=%d, ji=%d", graps[i][j], graps[j][i])
			}
		}
	}
	if ans1, ans2 := MinimumSpanningTree(9, graps); ans1 != 37 || ans2 != 37 {
		t.Fatalf("expect 37 ans1 %d, ans2 %d", ans1, ans2)
	}
}
