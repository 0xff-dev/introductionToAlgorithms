package leetcode

import (
	"fmt"
	"testing"
)

func debug2353(c FoodRatings) {
	for food, idx := range c.indies {
		fmt.Printf("food %s in %d\n", food, *idx)
	}
	for c, what := range c.h {
		fmt.Printf("cuisenes: %s\n", c)
		for _, n := range *what {
			fmt.Printf("\t name: %s, rating: %d, index: %d\n", n.name, n.rating, *n.index)
		}
	}
}
func TestConstructor2353(t *testing.T) {
	c := Constructor2352([]string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"},
		[]string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}, []int{9, 12, 8, 15, 14, 7})

	food := c.HighestRated("korean")
	if food != "kimchi" {
		t.Fatalf("expect kimchi get %s", food)
	}
	food = c.HighestRated("japanese")
	if food != "ramen" {
		t.Fatalf("expect ramen get %s", food)
	}

	c.ChangeRating("sushi", 16)
	food = c.HighestRated("japanese")
	if food != "sushi" {
		t.Fatalf("expect sushi get %s", food)
	}

	c.ChangeRating("ramen", 16)
	food = c.HighestRated("japanese")
	if food != "ramen" {
		t.Fatalf("expect ramen get %s", food)
	}
}
