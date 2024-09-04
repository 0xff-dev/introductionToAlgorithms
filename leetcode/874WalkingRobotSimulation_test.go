package leetcode

import "testing"

func TestRobotSim(t *testing.T) {
	commands, obstacles, exp := []int{4, -1, 3}, [][]int{}, 25
	if r := robotSim(commands, obstacles); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	commands, obstacles, exp = []int{4, -1, 4, -2, 4}, [][]int{{2, 4}}, 65
	if r := robotSim(commands, obstacles); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	commands, obstacles, exp = []int{6, -1, -1, 6}, [][]int{}, 36
	if r := robotSim(commands, obstacles); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
