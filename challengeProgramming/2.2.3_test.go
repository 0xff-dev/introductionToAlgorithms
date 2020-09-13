package challengeProgramming

import "testing"

func TestLexicalOrder(t *testing.T) {
	s := "ACDBCB"
	if res := LexicalOrder(s); res != "ABCBCD" {
		t.Fatalf("expect ABCBCD, get %s", res)
	}
}
