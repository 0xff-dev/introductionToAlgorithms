package leetcode

import (
	"reflect"
	"testing"
)

func TestUpdateBoard(t *testing.T) {
	b := [][]byte{
		{'E'},
	}
	exp := [][]byte{
		{'B'},
	}
	if r := updateBoard(b, []int{0, 0}); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	b = [][]byte{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
	}
	exp = [][]byte{
		{'B', '1', 'E', '1', 'B'},
		{'B', '1', 'M', '1', 'B'},
		{'B', '1', '1', '1', 'B'},
		{'B', 'B', 'B', 'B', 'B'},
	}
	if r := updateBoard(b, []int{3, 0}); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	b = [][]byte{
		{'B', '1', 'E', '1', 'B'},
		{'B', '1', 'M', '1', 'B'},
		{'B', '1', '1', '1', 'B'},
		{'B', 'B', 'B', 'B', 'B'},
	}
	exp = [][]byte{
		{'B', '1', 'E', '1', 'B'},
		{'B', '1', 'X', '1', 'B'},
		{'B', '1', '1', '1', 'B'},
		{'B', 'B', 'B', 'B', 'B'},
	}
	if r := updateBoard(b, []int{1, 2}); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

}
