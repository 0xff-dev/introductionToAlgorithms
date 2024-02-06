package leetcode

import (
	"reflect"
	"testing"
)

func TestSuggestedProducts(t *testing.T) {
	p := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	s := "mouse"
	exp := [][]string{
		{"mobile", "moneypot", "monitor"},
		{"mobile", "moneypot", "monitor"},
		{"mouse", "mousepad"},
		{"mouse", "mousepad"},
		{"mouse", "mousepad"},
	}
	if r := suggestedProducts(p, s); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	p = []string{"havana"}
	s = "havana"
	exp = [][]string{
		{"havana"},
		{"havana"},
		{"havana"},
		{"havana"},
		{"havana"},
		{"havana"},
	}
	if r := suggestedProducts(p, s); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
