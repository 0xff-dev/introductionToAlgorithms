package leetcode

import "testing"

func TestConstructor1396(t *testing.T) {
	c := Constructor1396()
	c.CheckIn(45, "Leyton", 3)
	c.CheckIn(32, "Paradise", 8)
	c.CheckIn(27, "Leyton", 10)
	c.CheckOut(45, "Waterloo", 15)
	c.CheckOut(27, "Waterloo", 20)
	c.CheckOut(32, "Cambridge", 22)
	if r := c.GetAverageTime("Paradise", "Cambridge"); r != 14 {
		t.Fatalf("expect 14 get %f", r)
	}
	if r := c.GetAverageTime("Leyton", "Waterloo"); r != 11 {
		t.Fatalf("expect 11 get %f", r)
	}
	c.CheckIn(10, "Leyton", 24)
	if r := c.GetAverageTime("Leyton", "Waterloo"); r != 11 {
		t.Fatalf("expect 11 get %f", r)
	}
	c.CheckOut(10, "Waterloo", 38)
	if r := c.GetAverageTime("Leyton", "Waterloo"); r != 12 {
		t.Fatalf("expect 12 get %f", r)
	}
}
