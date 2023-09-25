package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindDuplicate609(t *testing.T) {
	input := []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"}
	exp := [][]string{
		{"root/a/1.txt", "root/c/3.txt"},
		{"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"},
	}
	r := findDuplicate609(input)
	sort.Slice(r, func(i, j int) bool {
		a, b := len(r[i]), len(r[j])
		if a == b {
			for ii := 0; ii < a; ii++ {
				if r[i][ii] == r[j][ii] {
					continue
				}
				return r[i][ii] < r[j][ii]
			}
			return true
		}
		return a < b
	})
	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	input = []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)"}
	exp = [][]string{
		{"root/a/1.txt", "root/c/3.txt"},
		{"root/a/2.txt", "root/c/d/4.txt"},
	}
	r = findDuplicate609(input)
	sort.Slice(r, func(i, j int) bool {
		a, b := len(r[i]), len(r[j])
		if a == b {
			for ii := 0; ii < a; ii++ {
				if r[i][ii] == r[j][ii] {
					continue
				}
				return r[i][ii] < r[j][ii]
			}
			return true
		}
		return a < b
	})
	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
