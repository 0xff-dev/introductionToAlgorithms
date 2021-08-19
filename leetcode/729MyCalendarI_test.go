package leetcode

import "testing"

func TestMyCalendar(t *testing.T) {
	/*
	   o := Constructor2()
	   if r := o.Book(10, 20); !r {
	       t.Fatalf("expect true get  false")
	   }

	   if r := o.Book(15, 25); r {
	       t.Fatalf("expect false get true")
	   }

	   if r := o.Book(20, 30); !r {
	       t.Fatalf("expect true get false")
	   }
	*/

	b := Constructor2()
	if r := b.Book(2, 6); !r {
		t.Fatalf("expect true get false")
	}

	if r := b.Book(0, 2); !r {
		t.Fatalf("expect true get flase")
	}

	if r := b.Book(1, 3); r {
		t.Fatalf("expect false get true")
	}
	if r := b.Book(6, 10); !r {
		t.Fatalf("expect true get false")
	}

	c := Constructor2()
	if r := c.Book(3, 5); !r {
		t.Fatalf("expect true get false")
	}

	if r := c.Book(1, 6); r {
		t.Fatalf("expect false get true")
	}

	d := Constructor2()
	if r := d.Book(47, 50); !r {
		t.Fatalf("expect true get false")
	}

	if r := d.Book(33, 41); !r {
		t.Fatalf("expect true get false1")
	}

	if r := d.Book(39, 45); r {
		t.Fatalf("expect false get true")
	}

	if r := d.Book(33, 42); r {
		t.Fatalf("expect false get true")
	}

	if r := d.Book(25, 32); !r {
		t.Fatalf("expect true get false")
	}

	if r := d.Book(26, 35); r {
		t.Fatalf("expect false get true")
	}

	if r := d.Book(19, 25); !r {
		t.Fatalf("expect true get false")
	}
	if r := d.Book(3, 8); !r {
		t.Fatalf("expect true get false")
	}

	if r := d.Book(8, 13); !r {
		t.Fatalf("expect true get false")
	}

	if r := d.Book(18, 27); r {
		t.Fatalf("expect true get false")
	}

	e := Constructor2()
	if r := e.Book(37, 50); !r {
		t.Fatalf("expect true get false")
	}

	if r := e.Book(33, 50); r {
		t.Fatalf("expect false get true")
	}

	if r := e.Book(4, 17); !r {
		t.Fatalf("expect true get fale")
	}

	if r := e.Book(35, 48); r {
		t.Fatalf("expect false get true")
	}

	if r := e.Book(8, 25); r {
		t.Fatalf("expect false get true")
	}
}
