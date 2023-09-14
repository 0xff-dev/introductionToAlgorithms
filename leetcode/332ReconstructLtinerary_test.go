package leetcode

import (
	"reflect"
	"testing"
)

func TestFindItinerary(t *testing.T) {
	tic := [][]string{
		{"MUC", "LHR"},
		{"JFK", "MUC"},
		{"SFO", "SJC"},
		{"LHR", "SFO"},
	}
	exp := []string{"JFK", "MUC", "LHR", "SFO", "SJC"}
	if r := findItinerary(tic); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
