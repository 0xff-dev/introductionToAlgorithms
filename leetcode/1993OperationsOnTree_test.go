package leetcode

import (
	"reflect"
	"testing"
)

func TestConstructor1993(t *testing.T) {
	p := []int{-1, 0, 0, 1, 1, 2, 2}
	c := Constructor1993(p)
	exp := []bool{true, false, true, true, true, false}
	ans := make([]bool, 0)
	ans = append(ans, c.Lock(2, 2))
	ans = append(ans, c.Unlock(2, 3))
	ans = append(ans, c.Unlock(2, 2))
	ans = append(ans, c.Lock(4, 5))
	ans = append(ans, c.Upgrade(0, 1))
	ans = append(ans, c.Lock(0, 1))
	if !reflect.DeepEqual(exp, ans) {
		t.Fatalf("expect %v get %v", exp, ans)
	}

}
