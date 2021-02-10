package leetcode

import "testing"

func TestSimplifyPath(t *testing.T) {
	path := "/a"
	t.Log(simplifyPath(path))
	path = "/a/b/c"
	t.Log(simplifyPath(path))
	path = "/a/./b/c/.."
	t.Log(simplifyPath(path))
	path = "./a"
	t.Log(simplifyPath(path))
	path = "/../"
	t.Log(simplifyPath(path))
	path = "/"
	t.Log(simplifyPath(path))
	path = "/a/../../b/../c/./d"
	t.Log(simplifyPath(path))
	path = "/a/./b/../../c/"
	t.Log(simplifyPath(path))
}
