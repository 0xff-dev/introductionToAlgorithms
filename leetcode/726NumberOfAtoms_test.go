package leetcode

import "testing"

func TestCountOfAtoms(t *testing.T) {
	f := "H2O"
	exp := "H2O"
	if r := countOfAtoms(f); r != exp {
		t.Fatalf("expect '%s' get '%s'", exp, r)
	}

	f = "Mg(OH)2"
	exp = "H2MgO2"
	if r := countOfAtoms(f); r != exp {
		t.Fatalf("expect '%s' get '%s'", exp, r)
	}
	f = "K4(ON(SO3)2)2"
	exp = "K4N2O14S4"
	if r := countOfAtoms(f); r != exp {
		t.Fatalf("expect '%s' get '%s'", exp, r)
	}
	f = "H2(H3)2"
	exp = "H8"
	if r := countOfAtoms(f); r != exp {
		t.Fatalf("expect '%s' get '%s'", exp, r)
	}
	f = "H11He49NO35B7N46Li20"
	exp = "B7H11He49Li20N47O35"
	if r := countOfAtoms(f); r != exp {
		t.Fatalf("expect '%s' get '%s'", exp, r)
	}
}
