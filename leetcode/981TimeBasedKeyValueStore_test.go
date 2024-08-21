package leetcode

import (
	"reflect"
	"testing"
)

type input981 struct {
	key, val  string
	timestamp int
}

func TestConstructor981(t *testing.T) {
	inputs, exp := []input981{{"foo", "bar", 1}, {"foo", "", 1}, {"foo", "", 3}, {"foo", "bar2", 4}, {"foo", "", 4}, {"foo", "", 5}}, []string{"bar", "bar", "bar2", "bar2"}
	c := Constructor981()
	ans := make([]string, 0)
	for _, input := range inputs {
		if input.val == "" {
			ans = append(ans, c.Get(input.key, input.timestamp))
			continue
		}
		c.Set(input.key, input.val, input.timestamp)
	}
	if !reflect.DeepEqual(exp, ans) {
		t.Fatalf("expect %v get %v", exp, ans)
	}
}
