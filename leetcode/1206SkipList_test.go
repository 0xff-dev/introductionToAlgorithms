package leetcode

import (
	"fmt"
	"testing"
)

func TestConstructor1206(t *testing.T) {
	c := Constructor1206()
	tc := []int{5, 3, 1, 2, 2, 6}
	for _, item := range tc {
		c.Add(item)
	}

	for _, item := range append(tc, 7) {
		delSuccess := c.Erase(item)
		t.Logf("del: %v", delSuccess)
	}

	fmt.Println()
	for i := 1; i <= 6; i++ {
		r := c.Search(i)
		t.Logf("%v", r)
	}

	b := Constructor1206()
	b.Add(1)
	b.Add(2)
	b.Add(3)
	if b.Search(0) {
		t.Fatal("expect false get true")
	}
	b.Add(4)
	if !b.Search(1) {
		t.Fatalf("expect true get false")
	}

	if b.Erase(0) {
		t.Fatalf("expect false get true")
	}

	if !b.Erase(1) {
		t.Fatalf("expect true get false")
	}

	if b.Search(1) {
		t.Fatalf("expect false get true")
	}

	d := Constructor1206()
	//d.Add(3)
	//d.Add(14)
	//d.Add(3)
	//d.Add(18)
	//d.Add(1)
	//d.Add(18)
	//d.Add(17)
	//d.Add(8)
	//d.Add(5)
	//d.Add(16)
	//d.Add(3)
	//
	//if d.Search(10) {
	//	t.Fatalf("Search 10 expect false get true")
	//}
	//
	//if d.Search(9) {
	//	t.Fatalf("Search 9 expect false get true")
	//}
	//
	//if !d.Erase(14) {
	//	t.Fatalf("Erase 14 expect true get false")
	//}
	//
	//if !d.Erase(13) {
	//	t.Fatalf("Erase 13 expect true get false")
	//}
	//
	//d.Add(18)
	//
	//if d.Search(11) {
	//	t.Fatalf("Search 11 expect false get true")
	//}
	//
	//if !d.Erase(16) {
	//	t.Fatalf("Erase 16 expect true get false")
	//}
	//
	//if !d.Search(1) {
	//	t.Fatalf("Serach 1 expect true get false")
	//}
	//
	//if d.Search(4) {
	//	t.Fatalf("Search 4 expect false get true")
	//}
	//
	//d.Add(7)
	opts := []string{"add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "search", "search", "erase", "erase", "add", "search", "erase", "search", "search", "add", "search", "add", "erase", "erase", "add", "erase", "erase", "erase", "search", "search", "erase", "erase", "erase", "search", "add", "add", "add", "add", "erase", "erase", "add", "add", "erase", "erase", "erase", "search", "search", "add", "add", "erase", "erase", "search", "add", "search", "search", "search", "search", "search", "search", "search", "search", "search", "search", "search", "search"}
	nums := [][]int{{3}, {14}, {3}, {18}, {1}, {18}, {17}, {8}, {5}, {16}, {3}, {10}, {9}, {14}, {13}, {18}, {11}, {16}, {1}, {4}, {7}, {0}, {5}, {12}, {11}, {20}, {17}, {2}, {1}, {4}, {17}, {12}, {1}, {20}, {7}, {6}, {21}, {2}, {13}, {16}, {17}, {6}, {5}, {10}, {13}, {10}, {17}, {0}, {1}, {8}, {7}, {4}, {15}, {14}, {13}, {0}, {13}, {0}, {19}, {18}, {1}, {20}, {15}, {0}, {21}, {18}}
	ans := []string{"null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "false", "false", "true", "false", "null", "false", "true", "true", "false", "null", "false", "null", "false", "false", "null", "true", "false", "true", "false", "false", "false", "false", "true", "true", "null", "null", "null", "null", "false", "false", "null", "null", "false", "true", "false", "false", "false", "null", "null", "true", "false", "false", "null", "false", "false", "false", "false", "false", "true", "true", "false", "false", "false", "true", "true"}

	for idx, opt := range opts {
		fmt.Printf("-----> %s %d exp %s\n", opt, nums[idx][0], ans[idx])
		if opt == "add" {
			d.Add(nums[idx][0])
			continue
		}
		get := false
		if opt == "erase" {
			get = d.Erase(nums[idx][0])
		} else if opt == "search" {
			get = d.Search(nums[idx][0])
		}
		getStr := "false"
		if get {
			getStr = "true"
		}
		if getStr != ans[idx] {
			t.Fatalf("%s %d exp %s get %s", opt, nums[idx][0], ans[idx], getStr)
		}
	}

}
