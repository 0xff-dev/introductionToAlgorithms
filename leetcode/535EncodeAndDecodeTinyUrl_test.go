package leetcode

import "testing"

func TestCodec(t *testing.T) {
	u := "https://leetcode.com/problems/design-tinyurl"
	c := FConstructor()
	expect := c.encode(u)
	if expect != "https://leetcode.com/1" {
		t.Fatalf("expect %s get %s", "https://leetcode.com/1", expect)
	}

	uexpect := c.decode(expect)
	if uexpect != u {
		t.Fatalf("expect %s get %s", u, uexpect)
	}
}
