package offer

import "testing"

func TestReplaceSpace(t *testing.T) {
	testStr := "We are happy"
	res := ReplaceSpace(testStr)
	if res != "We%20are%20happy" {
		t.Fatalf("except %s get %s", "We%20are%20happy", res)
	}
	testStr = " we are happy "
	res = ReplaceSpace(testStr)
	if res != "%20we%20are%20happy%20" {
		t.Fatalf("except %s get %s", "%20we%20are%20happy%20", res)
	}
}
