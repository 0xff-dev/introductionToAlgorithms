package leetcode

import "testing"

func TestConstrutor2296(t *testing.T) {
	c := Constructor2296()
	c.AddText("leetcode")
	del := c.DeleteText(4)
	if del != 4 {
		t.Fatalf("expect 4 get %d", del)
	}
	c.AddText("practice")
	r := c.CursorRight(3)
	if r != "etpractice" {
		t.Fatalf("expect etpractice get %s", r)
	}

	r = c.CursorLeft(8)
	if r != "leet" {
		t.Fatalf("expect leet get %s", r)
	}

	del = c.DeleteText(10)
	if del != 4 {
		t.Fatalf("expect 4 get %d", del)
	}

	r = c.CursorRight(6)
	if r != "practi" {
		t.Fatalf("expect practi get %s", r)
	}
}

func TestConstrutor22961(t *testing.T) {
	c := Constructor2296()
	c.AddText("jxarid")
	r := c.CursorLeft(5)
	if r != "j" {
		t.Fatalf("expect j get %s", r)
	}
	r = c.CursorLeft(10)
	if r != "" {
		t.Fatalf("expect '' get %s", r)
	}

	c.AddText("du")
	del := c.DeleteText(20)
	if del != 2 {
		t.Fatalf("expect 2 get %d", del)
	}
}
