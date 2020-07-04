package string

import "testing"

func TestRabinKarp(t *testing.T) {
	s := "31415"
	_t := "2359023141526739921"
	RabinKarp(s, _t)
}

func TestKmp(t *testing.T) {
	s := "ababc"
	_t := "abaeababc"
	kmp(s, _t)
}