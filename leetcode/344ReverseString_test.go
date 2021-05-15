package leetcode

import "testing"

func TestReverseString(t *testing.T) {
	r := []byte("helloworld")
	reverseString([]byte(r))
	t.Logf("%v", string(r))
}
