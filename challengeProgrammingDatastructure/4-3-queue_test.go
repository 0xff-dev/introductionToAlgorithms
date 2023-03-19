package challengeprogrammingdatastructure

import "testing"

func TestCputasks(t *testing.T) {
	n, q := 5, 100
	tasks := []taskDesc{
		{"p1", 150},
		{"p2", 80},
		{"p3", 200},
		{"p4", 350},
		{"p5", 20},
	}
	exp := "p2180p5400p1450p3550p4800"
	if r := CpuTasks(n, q, tasks); r != exp {
		t.Fatalf("expect '%s' get %s", exp, r)
	}
}
