package leetcode

import "container/heap"

type Tweet struct {
	tweetId, id int
}

type Tweets []Tweet

func (t *Tweets) Len() int {
	return len(*t)
}

func (t *Tweets) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

func (t *Tweets) Less(i, j int) bool {
	return (*t)[i].id < (*t)[j].id
}

func (t *Tweets) Push(x interface{}) {
	*t = append(*t, x.(Tweet))
}

func (t *Tweets) Pop() interface{} {
	old := *t
	l := len(old)
	x := old[l-1]
	*t = old[:l-1]
	return x
}

type Twitter struct {
	tid        int
	userTweets map[int][]Tweet
	following  map[int]map[int]struct{}
}

func Constructor355() Twitter {
	return Twitter{
		tid:        0,
		userTweets: make(map[int][]Tweet),
		following:  make(map[int]map[int]struct{}),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.userTweets[userId]; !ok {
		this.userTweets[userId] = make([]Tweet, 0)
	}
	this.userTweets[userId] = append(this.userTweets[userId], Tweet{tweetId: tweetId, id: this.tid})
	this.tid++
}

func (this *Twitter) getUserTweets(userId int, tweets *Tweets, limit int) {
	for _, tt := range this.userTweets[userId] {
		// 1,2 3
		if len(*tweets) < limit {
			heap.Push(tweets, tt)
			continue
		}
		top := heap.Pop(tweets).(Tweet)
		if tt.id > top.id {
			heap.Push(tweets, tt)
		} else {
			heap.Push(tweets, top)
		}
	}
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	tweets := Tweets{}
	this.getUserTweets(userId, &tweets, 10)
	for following := range this.following[userId] {
		this.getUserTweets(following, &tweets, 10)
	}
	tids := make([]int, 0)
	for tweets.Len() > 0 {
		top := heap.Pop(&tweets).(Tweet)
		tids = append(tids, top.tweetId)
	}
	for s, e := 0, len(tids)-1; s < e; s, e = s+1, e-1 {
		tids[s], tids[e] = tids[e], tids[s]
	}
	return tids
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.following[followerId]; !ok {
		this.following[followerId] = make(map[int]struct{})
	}
	this.following[followerId][followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if v, ok := this.following[followerId]; ok {
		delete(v, followeeId)
		this.following[followerId] = v
	}
}
