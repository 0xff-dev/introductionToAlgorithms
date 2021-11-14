package leetcode

import "testing"

func TestConstructorForCom(t *testing.T) {
	s, length := "abc", 2
	o := ConstructorForCom(s, length)
	if r := o.Next(); r != "ab" {
		t.Fatalf("expect ab get %s", r)
	}
	if !o.HasNext() {
		t.Fatalf("expect true get false")
	}
	if r := o.Next(); r != "ac" {
		t.Fatalf("expect ac get %s", r)
	}

	if !o.HasNext() {
		t.Fatalf("expect true get false")
	}

	if r := o.Next(); r != "bc" {
		t.Fatalf("expect bc get %s", r)
	}

	if o.HasNext() {
		t.Fatalf("expect false get true")
	}
}
