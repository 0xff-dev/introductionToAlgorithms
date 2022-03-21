package leetcode

import (
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	s, expect := "ababcbacadefegdehijhklij", []int{9, 7, 8}
	if r := partitionLabels(s); !reflect.DeepEqual(r, expect) {
		t.Fatalf("expect %v get %v", expect, r)
	}

	s, expect = "eccbbbbdec", []int{10}
	if r := partitionLabels(s); !reflect.DeepEqual(r, expect) {
		t.Fatalf("expect %v get %v", expect, r)
	}
}
