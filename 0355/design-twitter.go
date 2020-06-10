package _0355

// 2020.06.10
// https://leetcode-cn.com/problems/design-twitter/
type Twitter struct {
	follow map[int]map[int]struct{}
	posts  []Post
}

type Post struct {
	PostID int
	UserID int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{follow: make(map[int]map[int]struct{})}
}

func (this *Twitter) newUser(userID int) {
	if _, ok := this.follow[userID]; !ok {
		this.follow[userID] = make(map[int]struct{})
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.newUser(userId)
	this.posts = append(this.posts, Post{PostID: tweetId, UserID: userId})
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	this.newUser(userId)
	count := 10
	var result []int
	for i := len(this.posts) - 1; i >= 0 && count > 0; i-- {
		_, ok := this.follow[userId][this.posts[i].UserID]
		if this.posts[i].UserID == userId || ok {
			result = append(result, this.posts[i].PostID)
			count--
		}
	}
	return result
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	this.newUser(followerId)
	this.newUser(followeeId)
	this.follow[followerId][followeeId] = struct{}{}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	this.newUser(followerId)
	delete(this.follow[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
