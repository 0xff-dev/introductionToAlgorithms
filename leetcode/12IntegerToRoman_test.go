package leetcode

import "testing"

func TestInToRoman(t *testing.T) {
	n, roman := 3, "III"
	if r := intToRoman(n); r != roman {
		t.Fatalf("expect %s get %s", roman, r)
	}

	n, roman = 58, "LVIII"
	if r := intToRoman(n); r != roman {
		t.Fatalf("expect %s get %s", roman, r)
	}

	n, roman = 3999, "MMMCMXCIX"
	if r := intToRoman(n); r != roman {
		t.Fatalf("expect %s get %s", roman, r)
	}

	n, roman = 789, "DCCLXXXIX"
	if r := intToRoman(n); r != roman {
		t.Fatalf("expect %s get %s", roman, r)
	}
}
