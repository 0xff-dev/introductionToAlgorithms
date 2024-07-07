package leetcode

import (
	"reflect"
	"testing"
)

func TestConstructor1348(t *testing.T) {
	c := Constructor1348()
	c.RecordTweet("tweet3", 0)
	c.RecordTweet("tweet3", 60)
	c.RecordTweet("tweet3", 10)
	exp := []int{2}
	if r := c.GetTweetCountsPerFrequency("minute", "tweet3", 0, 59); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	exp = []int{2, 1}
	if r := c.GetTweetCountsPerFrequency("minute", "tweet3", 0, 60); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	exp = []int{4}
	c.RecordTweet("tweet3", 120)
	if r := c.GetTweetCountsPerFrequency("hour", "tweet3", 0, 210); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
