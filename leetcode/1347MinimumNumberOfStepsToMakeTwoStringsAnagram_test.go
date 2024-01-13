package leetcode

import "testing"

func TestMinSteps1347(t *testing.T) {
	s := "leetcode"
	ts := "practice"
	if r := minSteps1347(s, ts); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
	s = "gctcxyuluxjuxnsvmomavutrrfb"
	ts = "qijrjrhqqjxjtprybrzpyfyqtzf"
	if r := minSteps1347(s, ts); r != 18 {
		t.Fatalf("expect 18 get %d", r)
	}
}
