package leetcode

import "testing"

func TestKcodes(t *testing.T) {
	k := 4
	bytes := make([]byte, k)
	kcodes(bytes, 0, k)
}

func TestPow2(t *testing.T) {
	if r := pow2(3); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}

	if r := pow2(1); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	if r := pow2(10); r != 1024 {
		t.Fatalf("expect 1024 get %d", r)
	}
}

func TestHasAllCodes(t *testing.T) {
	s, k := "00110110", 2
	if r := hasAllCodes(s, k); !r {
		t.Fatalf("expect true get false")
	}

	s, k = "0110", 1
	if r := hasAllCodes(s, k); !r {
		t.Fatalf("expect true get false")
	}

	s, k = "0110", 2
	if r := hasAllCodes(s, k); r {
		t.Fatalf("expect false get true")
	}
	s, k = "00110", 2
	if r := hasAllCodes(s, k); !r {
		t.Fatalf("expect true get false")
	}
}
