package leetcode

import "testing"

func TestGetImportance(t *testing.T) {
	employees := []*Employee{
		&Employee{
			Id:           1,
			Importance:   5,
			Subordinates: []int{2, 3},
		},
		&Employee{Id: 2, Importance: 3},
		&Employee{Id: 3, Importance: 3},
	}
	if r := getImportance(employees, 1); r != 11 {
		t.Fatalf("expect 11 get %d", r)
	}

	employees = []*Employee{
		&Employee{Id: 1, Importance: 1, Subordinates: []int{5}},
		&Employee{Id: 5, Importance: -3},
	}
	if r := getImportance(employees, 5); r != -3 {
		t.Fatalf("expect -3 get %d", r)
	}
}
