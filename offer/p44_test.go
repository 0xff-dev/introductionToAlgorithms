package offer

import "testing"

func TestIsStraight(t *testing.T) {
	outputTemplate := "input: %s, expect %v"
	input := "460"
	if IsStraight(input) != true {
		t.Fatalf(outputTemplate, input, true)
	}
	input = "470"
	if IsStraight(input) != false {
		t.Fatalf(outputTemplate, input, false)
	}
	input = "246A"
	if IsStraight(input) != false {
		t.Fatalf(outputTemplate, input, false)
	}
	input = "A2345"
	if IsStraight(input) != true {
		t.Fatalf(outputTemplate, input, true)
	}
	input = "JQK"
	if IsStraight(input) != true {
		t.Fatalf(outputTemplate, input, true)
	}
}
