package leetcode

import "testing"

func TestThousandSeparator(t *testing.T) {
	n := 987
	if r := thousandSeparator(n); r != "987" {
		t.Fatalf("expect %s get %s", "987", r)
	}

	n = 1234
	if r := thousandSeparator(n); r != "1.234" {
		t.Fatalf("expect %s get %s", "1.234", r)
	}

	n = 123456789
	if r := thousandSeparator(n); r != "123.456.789" {
		t.Fatalf("expect %s get %s", "123.456.789", r)
	}
}
