package leetcode

import "testing"

func TestRemoveComments(t *testing.T) {
	code := []string{"/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test", "   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"}
	r := removeComments(code)
	t.Logf("%q", r)
	code = []string{"a/*comment", "line", "more_comment*/b"}
	r = removeComments(code)
	t.Logf("%q", r)
	code = []string{"/*asdfsdf*/"}
	r = removeComments(code)
	t.Logf("%q", r)

	code = []string{"asdf/* asdf", "aaaa", "bbbbb*/ coa // fuc"}
	r = removeComments(code)
	t.Logf("%q", r)

	code = []string{"int comments;// there are comments", "/*todo*/", "/* one, two", "three, four", " five*/ six, // seven"}
	r = removeComments(code)
	t.Logf("%q", r)

	code = []string{"struct Node{", "    /*/ declare members;/**/", "    int size;", "    /**/int val;", "};"}
	r = removeComments(code)
	t.Logf("%q", r)

	code = []string{"void func(int k) {", "// this function does nothing /*", "   k = k*2/4;", "   k = k/2;*/", "}"}
	r = removeComments(code)
	t.Logf("%q", r)

	code = []string{"eed*//cd//**///*/*//e//*bbcbbaedb*//aabb//*badb*//d/*/e*//ade//*bacbc*//ea//*a"}
	r = removeComments(code)
	t.Logf("%q", r)
}
