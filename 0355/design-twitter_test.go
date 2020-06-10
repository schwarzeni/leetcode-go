package _0355

import (
	"reflect"
	"testing"
)

func TestTwitter(t *testing.T) {
	twitter := Constructor()
	twitter.Follow(1, 2)
	twitter.Follow(1, 3)
	twitter.Follow(1, 4)
	twitter.Follow(2, 4)
	twitter.PostTweet(1, 1)
	twitter.PostTweet(1, 2)
	twitter.PostTweet(1, 3)
	twitter.PostTweet(1, 4)
	twitter.PostTweet(1, 5)
	twitter.PostTweet(2, 6)
	twitter.PostTweet(3, 7)
	twitter.PostTweet(2, 8)
	twitter.PostTweet(4, 9)
	twitter.PostTweet(2, 10)
	twitter.PostTweet(3, 11)
	if !reflect.DeepEqual(twitter.GetNewsFeed(1), []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}) {
		t.Errorf("get feed 1 failed: %v", twitter.GetNewsFeed(1))
	}
	twitter.Unfollow(1, 3)
	if !reflect.DeepEqual(twitter.GetNewsFeed(1), []int{10, 9, 8, 6, 5, 4, 3, 2, 1}) {
		t.Errorf("get feed 2 failed: %v", twitter.GetNewsFeed(1))
	}
}
