package golang

import "testing"

func TestCalculation(t *testing.T) {
	input := "(5*3+4/2)"
	if r := calculation(input); r != 17 {
		t.Fatalf("expect 17 get %d", r)
	}

	input = "45/9+12*3-2/2+33"
	if r := calculation(input); r != 73 {
		t.Fatalf("expect 73 get %d", r)
	}

	input = "1*2*3*4"
	if r := calculation(input); r != 24 {
		t.Fatalf("expect 24 get %d", r)
	}

	input = "(2+4)*(2*44/2)+22"
	if r := calculation(input); r != 286 {
		t.Fatalf("expect 286 get %d", r)
	}

	input = "(2*5+10)/2+33/3-4/2+3*6"
	if r := calculation(input); r != 37 {
		t.Fatalf("expect 37 get %d", r)
	}

	input = "(10/5*3)/2+(2+4+(2*5))/4"
	if r := calculation(input); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}

	input = "(((2+3)+4)*6)"
	if r := calculation(input); r != 54 {
		t.Fatalf("expect 54 get %d", r)
	}

	input = "(((2*50)/2+5)*3)+3"
	if r := calculation(input); r != 168 {
		t.Fatalf("expect 168 get %d", r)
	}

	input = "3(2+4)"
	if r := calculation(input); r != 18 {
		t.Fatalf("expect 18 get %d", r)
	}

	input = "(2*3)+((3*4)+3)/5-2/2-3/3"
	if r := calculation(input); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}

	input = "(((2*3)*4)*5)-4-2*2"
	if r := calculation(input); r != 112 {
		t.Fatalf("expect 112 get %d", r)
	}
}
