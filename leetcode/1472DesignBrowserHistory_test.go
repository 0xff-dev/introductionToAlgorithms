package leetcode

import "testing"

func TestConstrcutor1472(t *testing.T) {
	c := Constructor1472("leetcode.com")
	c.Visit("google.com")
	c.Visit("facebook.com")
	c.Visit("youtube.com")
	if r := c.Back(1); r != "facebook.com" {
		t.Fatalf("expect 'facebookc.com' get '%s'", r)
	}

	if r := c.Back(1); r != "google.com" {
		t.Fatalf("expect 'google.com' get '%s'", r)
	}

	if r := c.Forward(1); r != "facebook.com" {
		t.Fatalf("expect 'facebook.com' get '%s'", r)
	}

	c.Visit("linkedin.com")
	if r := c.Forward(2); r != "linkedin.com" {
		t.Fatalf("expect 'linkedin.com' get '%s'", r)
	}

	if r := c.Back(2); r != "google.com" {
		t.Fatalf("expect 'google.com' get '%s'", r)
	}

	if r := c.Back(7); r != "leetcode.com" {
		t.Fatalf("expect 'leetcode.com' get '%s'", r)
	}
}
