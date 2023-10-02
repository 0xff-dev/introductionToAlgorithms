package leetcode

import (
	"reflect"
	"testing"
)

func TestConstructor355(t *testing.T) {
	twitter := Constructor355()
	// post
	twitter.PostTweet(1, 5)
	r := twitter.GetNewsFeed(1)
	exp := []int{5}
	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)
	r = twitter.GetNewsFeed(1)
	exp = []int{6, 5}
	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	twitter.Unfollow(1, 2)
	r = twitter.GetNewsFeed(1)
	exp = []int{5}
	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
