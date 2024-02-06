package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestRemoveSubfolders(t *testing.T) {
	folder := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
	exp := []string{"/a", "/c/d", "/c/f"}
	r := removeSubfolders(folder)
	sort.Strings(r)
	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	folder = []string{"/a", "/a/b/c", "/a/b/d"}
	exp = []string{"/a"}
	r = removeSubfolders(folder)
	sort.Strings(r)
	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	folder = []string{"/a/b/c", "/a/b/ca", "/a/b/d"}
	exp = []string{"/a/b/c", "/a/b/ca", "/a/b/d"}
	r = removeSubfolders(folder)
	sort.Strings(r)
	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
