package offer

import "testing"

func TestReverseString(t *testing.T) {
	testStr := "I am a student."
	result := "student. a am I"
	if res := ReverseString(testStr); res != result {
		t.Fatalf("expect %s get %s", result, res)
	}
}

func TestLeftRotateString(t *testing.T) {
	testStr := "abcdefg"
	for kTh := 0; kTh < 8; kTh++ {
		result := LeftRotateString(testStr, kTh)
		t.Logf("kth: %d result: %s", kTh, result)
	}
}
