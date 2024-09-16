package leetcode

import "testing"

func TestMaskPII(t *testing.T) {
	s, exp := "LeetCode@LeetCode.com", "l*****e@leetcode.com"
	if r := maskPII(s); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
	s, exp = "AB@qq.com", "a*****b@qq.com"
	if r := maskPII(s); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
	s, exp = "1(234)567-890", "***-***-7890"
	if r := maskPII(s); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
