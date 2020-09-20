package dp

import "testing"

type RY struct {
	Str string
	Ans int
}

func TestToRYR(t *testing.T) {
	for _, testCase := range []RY{
		{"rrr", 1},
		{"ryr", 0},
		{"yry", 3},
		{"yrry", 3},
		{"yryry", 2},
		{"ryrrrrrrrryr", 1},
		{"rrryyyryyyyrr", 1},
		{"rrryyyrryyyrr", 2},
		{"yyy", 2},
		{"yyrryyryryry", 5},
	} {
		if ans := ToRYR([]byte(testCase.Str)); ans != testCase.Ans {
			t.Fatalf("expect %d get %d\n", testCase.Ans, ans)
		}
	}
}
