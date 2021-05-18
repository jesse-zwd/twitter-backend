package service

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Response basic serializer
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

// TrackedErrorResponse
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// Error code
const (
	// CodeCheckLogin
	CodeCheckLogin = 401
	// CodeNoRightErr
	CodeNoRightErr = 403
	// CodeDBError database error
	CodeDBError = 50001
	// CodeEncryptError encrypting failed
	CodeEncryptError = 50002
	//CodeParamErr other param error
	CodeParamErr = 40001
)

// Err handling common error
func Err(errCode int, msg string, err error) Response {
	res := Response{
		Code: errCode,
		Msg:  msg,
	}

	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}

// DBErr database error
func DBErr(msg string, err error) Response {
	if msg == "" {
		msg = "database error"
	}
	return Err(CodeDBError, msg, err)
}

// ParamErr other errors
func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "param error"
	}
	return Err(CodeParamErr, msg, err)
}

// DataList basic list
type DataList struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

type APIUser struct {
	ID       uint   `form:"id" json:"id"`
	UserName string `form:"username" json:"username"`
	Nickname string `form:"nickname" json:"nickname"`
	Avatar   string `form:"avatar" json:"avatar"`
}

type APIListUser struct {
	ID          uint   `form:"id" json:"id"`
	UserName    string `form:"username" json:"username"`
	Nickname    string `form:"nickname" json:"nickname"`
	Avatar      string `form:"avatar" json:"avatar"`
	IsFollowing bool   `form:"isFollowing" json:"isFollowing"`
	IsMe        bool   `form:"isMe" json:"isMe"`
}

type APIComment struct {
	ID          uint      `form:"id" json:"id"`
	Text        string    `form:"text" json:"text"`
	CommentUser APIUser   `form:"user" json:"user"`
	CreatedAt   time.Time `form:"created_at" json:"created_at"`
}

type APITweetFile struct {
	Url string `form:"url" json:"url"`
}

type APITweet struct {
	ID           uint           `form:"id" json:"id"`
	Text         string         `form:"text" json:"text"`
	Tags         string         `form:"tags" json:"tags"`
	TweetFiles   []APITweetFile `form:"files" json:"files"`
	LikeCount    int64          `form:"likeCount" json:"likeCount"`
	TweetUser    APIUser        `form:"user" json:"user"`
	IsTweetMine  bool           `form:"isTweetMine" json:"isTweetMine"`
	IsLiked      bool           `form:"isLiked" json:"isLiked"`
	CommentCount int64          `form:"commentCount" json:"commentCount"`
	IsRetweet    bool           `form:"isRetweet" json:"isRetweet"`
	RetweetCount int64          `form:"retweetCount" json:"retweetCount"`
	CreatedAt    time.Time      `form:"created_at" json:"created_at"`
}

type APITweetByID struct {
	ID           uint           `form:"id" json:"id"`
	Text         string         `form:"text" json:"text"`
	Tags         string         `form:"tags" json:"tags"`
	TweetFiles   []APITweetFile `form:"files" json:"files"`
	Comments     []APIComment   `form:"comments" json:"comments"`
	LikeCount    int64          `form:"likeCount" json:"likeCount"`
	TweetUser    APIUser        `form:"user" json:"user"`
	IsTweetMine  bool           `form:"isTweetMine" json:"isTweetMine"`
	IsLiked      bool           `form:"isLiked" json:"isLiked"`
	CommentCount int64          `form:"commentCount" json:"commentCount"`
	IsRetweet    bool           `form:"isRetweet" json:"isRetweet"`
	RetweetCount int64          `form:"retweetCount" json:"retweetCount"`
	CreatedAt    time.Time      `form:"created_at" json:"created_at"`
}

type APIProfile struct {
	ID             uint       `form:"id" json:"id"`
	UserName       string     `form:"username" json:"username"`
	Nickname       string     `form:"nickname" json:"nickname"`
	Avatar         string     `form:"avatar" json:"avatar"`
	Status         string     `form:"status" json:"status"`
	Email          string     `form:"email" json:"email"`
	Bio            string     `form:"bio" json:"bio"`
	Website        string     `form:"website" json:"website"`
	CoverPhoto     string     `form:"coverPhoto" json:"coverPhoto"`
	Dob            time.Time  `form:"dob" json:"dob"`
	Location       string     `form:"location" json:"location"`
	Tweets         []APITweet `form:"tweets" json:"tweets"`
	Followers      []APIUser  `form:"followers" json:"followers"`
	Followings     []APIUser  `form:"followings" json:"followings"`
	TweetCount     int64      `form:"tweetCount" json:"tweetCount"`
	FollowingCount int64      `form:"followingCount" json:"followingCount"`
	FollowerCount  int64      `form:"followerCount" json:"followerCount"`
	IsFollowing    bool       `form:"isFollowing" json:"isFollowing"`
	IsMe           bool       `form:"isMe" json:"isMe"`
}
