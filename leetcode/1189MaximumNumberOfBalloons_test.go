package leetcode

import "testing"

func TestMaxNumberOfBalloons(t *testing.T) {
	text := "nlaebolko"
	if r := maxNumberOfBalloons(text); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	text = "leetcode"
	if r := maxNumberOfBalloons(text); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	text = "loonbalxballpoon"
	if r := maxNumberOfBalloons(text); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
