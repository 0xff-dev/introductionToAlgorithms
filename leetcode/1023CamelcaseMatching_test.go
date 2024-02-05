package leetcode

import (
	"reflect"
	"testing"
)

func TestCamelMatch(t *testing.T) {
	queries := []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}
	pattern := "FB"
	exp := []bool{true, false, true, true, false}
	if r := camelMatch(queries, pattern); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	queries = []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}
	pattern = "FoBa"
	exp = []bool{true, false, true, false, false}
	if r := camelMatch(queries, pattern); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	queries = []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}
	pattern = "FoBaT"
	exp = []bool{false, true, false, false, false}
	if r := camelMatch(queries, pattern); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
