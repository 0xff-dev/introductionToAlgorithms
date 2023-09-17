package leetcode

import "testing"

func TestLengthLongestPath(t *testing.T) {
	input := "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
	if r := lengthLongestPath(input); r != 20 {
		t.Fatalf("expect 20 get %d", r)
	}
	input = "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
	if r := lengthLongestPath(input); r != 32 {
		t.Fatalf("expect 32 get %d", r)
	}
	input = "a"
	if r := lengthLongestPath(input); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	input = "a.txt"
	if r := lengthLongestPath(input); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
	input = "file1.txt\nfile2.txt\tlongfile.txt"
	if r := lengthLongestPath(input); r != 22 {
		t.Fatalf("expect 22 get %d", r)
	}
	input = "a\n\tb.txt\na2\n\tb2.txt"
	if r := lengthLongestPath(input); r != 9 {
		t.Fatalf("expect 9 get %d", r)
	}
	input = "a\n\taa\n\t\taaa\n\t\t\tfile1.txt\naaaaaaaaaaaaaaaaaaaaa\n\tsth.png"
	if r := lengthLongestPath(input); r != 29 {
		t.Fatalf("expect 29 get %d", r)
	}
}
