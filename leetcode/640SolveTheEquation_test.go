package leetcode

import "testing"

func TestParseExpression(t *testing.T) {
	expression := "3x-2x+x-5+3"
	ea, eb := 2, -2
	if a, b := parseExpression(expression); a != ea || b != eb {
		t.Fatalf("expression: %s expect %d %d get %d %d", expression, ea, eb, a, b)
	}
	expression = "-2x+x-5+3"
	ea, eb = -1, -2
	if a, b := parseExpression(expression); a != ea || b != eb {
		t.Fatalf("expression: %s expect %d %d get %d %d", expression, ea, eb, a, b)
	}
	expression = "x+5-3+x"
	ea, eb = 2, 2
	if a, b := parseExpression(expression); a != ea || b != eb {
		t.Fatalf("expression: %s expect %d %d get %d %d", expression, ea, eb, a, b)
	}
	expression = "x"
	ea, eb = 1, 0
	if a, b := parseExpression(expression); a != ea || b != eb {
		t.Fatalf("expression: %s expect %d %d get %d %d", expression, ea, eb, a, b)
	}
	expression = "6+x-2"
	ea, eb = 1, 4
	if a, b := parseExpression(expression); a != ea || b != eb {
		t.Fatalf("expression: %s expect %d %d get %d %d", expression, ea, eb, a, b)
	}
}

func TestGCD640(t *testing.T) {
	if r := gcd640(1, 1); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	if r := gcd640(8, 4); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	if r := gcd640(12, 2); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
func TestSolveEquation(t *testing.T) {
	equation := "x+5-3+x=6+x-2"
	exp := "x=2"
	if r := solveEquation(equation); r != exp {
		t.Fatalf("with %s expect %s get %s", equation, exp, r)
	}
	equation = "x=x"
	exp = "Infinite solutions"
	if r := solveEquation(equation); r != exp {
		t.Fatalf("with %s expect %s get %s", equation, exp, r)
	}
	equation = "2x=x"
	exp = "x=0"
	if r := solveEquation(equation); r != exp {
		t.Fatalf("with %s expect %s get %s", equation, exp, r)
	}
	equation = "2x+3x-6x=x+2"
	exp = "x=-1"
	if r := solveEquation(equation); r != exp {
		t.Fatalf("with %s expect %s get %s", equation, exp, r)
	}
	equation = "2x+0=0"
	exp = "x=0"
	if r := solveEquation(equation); r != exp {
		t.Fatalf("with %s expect %s get %s", equation, exp, r)
	}
	equation = "0x=0"
	exp = "Infinite solutions"
	if r := solveEquation(equation); r != exp {
		t.Fatalf("with %s expect %s get %s", equation, exp, r)
	}
}
