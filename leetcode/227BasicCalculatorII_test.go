package leetcode

import "testing"

func TestCalculate(t *testing.T) {
	s := "3+2*2"
	exp := 7
	if r := calculate(s); r != exp {
		t.Fatalf("input %s expect %d get %d", s, exp, r)
	}
	s = "3/2 "
	exp = 1
	if r := calculate(s); r != exp {
		t.Fatalf("input %s expect %d get %d", s, exp, r)

	}

	s = " 3+5 / 2 "
	exp = 5
	if r := calculate(s); r != exp {
		t.Fatalf("input %s expect %d get %d", s, exp, r)

	}
	s = "1*2-3/4+5*6-7*8+9/10"
	exp = -24
	if r := calculate(s); r != exp {
		t.Fatalf("input %s expect %d get %d", s, exp, r)

	}
	s = "282-1*2*13-30-2*2*2/2-95/5*2+55+804+3024"
	exp = 4067
	if r := calculate(s); r != exp {
		t.Fatalf("input %s expect %d get %d", s, exp, r)

	}
}
