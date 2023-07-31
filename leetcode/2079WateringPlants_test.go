package leetcode

import "testing"

func TestWateringPlants(t *testing.T) {
	plants := []int{2, 2, 3, 3}
	c := 5
	if r := wateringPlants(plants, c); r != 14 {
		t.Fatalf("expect 14 get %d", r)
	}
	plants = []int{1, 1, 1, 4, 2, 3}
	c = 4
	if r := wateringPlants(plants, c); r != 30 {
		t.Fatalf("expect 30 get %d", r)
	}

	plants = []int{7, 7, 7, 7, 7, 7, 7}
	c = 8
	if r := wateringPlants(plants, c); r != 49 {
		t.Fatalf("expect 49 get %d", r)
	}
}
