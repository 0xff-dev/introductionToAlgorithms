package leetcode

import (
	"reflect"
	"testing"
)

func TestSmallestSufficientTeam(t *testing.T) {
	req := []string{"java", "nodejs", "reactjs"}
	people := [][]string{
		{"java"}, {"nodejs"}, {"nodejs", "reactjs"},
	}
	exp := []int{0, 2}
	if r := smallestSufficientTeam(req, people); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	req = []string{"algorithms", "math", "java", "reactjs", "csharp", "aws"}
	people = [][]string{
		{"algorithms", "math", "java"},
		{"algorithms", "math", "reactjs"},
		{"java", "csharp", "aws"},
		{"reactjs", "csharp"},
		{"csharp", "math"},
		{"aws", "java"},
	}
	exp = []int{1, 2}
	if r := smallestSufficientTeam(req, people); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
