package service

import (
	"backend/global"
	"backend/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BuildAPITweet(c *gin.Context, tweet model.Tweet) APITweet {
	return APITweet{
		ID:           tweet.ID,
		Text:         tweet.Text,
		Tags:         tweet.Tags,
		TweetFiles:   GetTweetFiles(&tweet),
		LikeCount:    GetLikeCount(&tweet),
		CommentCount: GetCommentCount(&tweet),
		TweetUser:    GetTweetUser(&tweet),
		IsTweetMine:  IsTweetMine(c, &tweet),
		IsLiked:      IsLiked(c, &tweet),
		IsRetweet:    IsRetweet(c, &tweet),
		RetweetCount: GetRetweetCount(&tweet),
		CreatedAt:    tweet.CreatedAt.Local(),
	}
}

func BuildAPITweets(c *gin.Context, tweets []model.Tweet) (apiTweets []APITweet) {
	for _, tweet := range tweets {
		apiTweet := BuildAPITweet(c, tweet)
		apiTweets = append(apiTweets, apiTweet)
	}
	return
}

func BuildAPITweetByID(c *gin.Context, tweet model.Tweet) APITweetByID {
	return APITweetByID{
		ID:           tweet.ID,
		Text:         tweet.Text,
		Tags:         tweet.Tags,
		TweetFiles:   GetTweetFiles(&tweet),
		Comments:     BuildAPIComments(tweet.Comments),
		LikeCount:    GetLikeCount(&tweet),
		CommentCount: GetCommentCount(&tweet),
		TweetUser:    GetTweetUser(&tweet),
		IsTweetMine:  IsTweetMine(c, &tweet),
		IsLiked:      IsLiked(c, &tweet),
		IsRetweet:    IsRetweet(c, &tweet),
		RetweetCount: GetRetweetCount(&tweet),
		CreatedAt:    tweet.CreatedAt.Local(),
	}
}

func (service *ListTweetService) List(c *gin.Context) Response {
	tweets := []model.Tweet{}
	var user_ids []uint

	currentUser := CurrentUser(c)
	user_ids = append(user_ids, currentUser.ID)

	followings := GetFollowings(currentUser)
	for _, following := range followings {
		user_ids = append(user_ids, following.ID)
	}

	if service.Limit == 0 {
		service.Limit = 12
	}

	if err := global.GORM_DB.Where("user_id IN ?", user_ids).Limit(service.Limit).Offset(service.Start).Order("created_at desc").Find(&tweets).Error; err != nil {
		return Response{
			Code:  50001,
			Msg:   "database connection error",
			Error: err.Error(),
		}
	}

	apiTweets := BuildAPITweets(c, tweets)

	return Response{
		Data: apiTweets,
		Msg:  "Got tweets",
	}
}

func (service *CreateTweetService) Create(c *gin.Context) Response {
	currentUser := CurrentUser(c)
	tweet := model.Tweet{
		Text:   service.Text,
		Tags:   service.Tags,
		UserID: currentUser.ID,
	}

	// Create post
	if err := global.GORM_DB.Create(&tweet).Error; err != nil {
		return Response{
			Code:  50001,
			Msg:   "creating tweet failed",
			Error: err.Error(),
		}
	}

	if service.TweetFiles != nil {
		for _, tweetfile := range service.TweetFiles {
			url := tweetfile.Url
			CreateTweetFile(url, tweet.ID)
		}
	}

	apiTweet := BuildAPITweet(c, tweet)

	return Response{
		Data: apiTweet,
		Msg:  "tweet created",
	}
}

func (service *DeleteTweetService) Delete(id string) Response {
	var tweet model.Tweet

	tweetId, _ := strconv.Atoi(id)
	err := global.GORM_DB.First(&tweet, uint(tweetId)).Error
	if err != nil {
		return Response{
			Code:  404,
			Msg:   "tweet doesn't exist",
			Error: err.Error(),
		}
	}

	// delete comments of tweet
	err = global.GORM_DB.Where("tweet_id = ?", tweet.ID).Delete(&model.Comment{}).Error
	if err != nil {
		return Response{
			Code:  50001,
			Msg:   "deleting comments of tweet failed",
			Error: err.Error(),
		}
	}

	// delete likes of tweet
	err = global.GORM_DB.Where("tweet_id = ?", tweet.ID).Delete(&model.Like{}).Error
	if err != nil {
		return Response{
			Code:  50001,
			Msg:   "deleting likes of tweet failed",
			Error: err.Error(),
		}
	}

	// delete retweet of tweet
	err = global.GORM_DB.Where("tweet_id = ?", tweet.ID).Delete(&model.Retweet{}).Error
	if err != nil {
		return Response{
			Code:  50001,
			Msg:   "deleting retweets of tweet failed",
			Error: err.Error(),
		}
	}

	// delete tweetfiles and tweet
	err = global.GORM_DB.Select("TweetFiles").Delete(&tweet).Error
	if err != nil {
		return Response{
			Code:  50001,
			Msg:   "deleting tweet failed",
			Error: err.Error(),
		}
	}

	return Response{
		Code: 0,
		Msg:  "tweet deleted",
	}
}

func GetTweetUser(tweet *model.Tweet) APIUser {
	tweetUser, _ := GetUser(tweet.UserID)
	return BuildAPIUser(tweetUser)
}

func GetTweetFiles(tweet *model.Tweet) (tweetfiles []APITweetFile) {
	global.GORM_DB.Model(&tweet).Association("TweetFiles").Find(&tweetfiles)
	return
}

func GetComments(tweet *model.Tweet) (comments []model.Comment) {
	global.GORM_DB.Model(&model.Comment{}).Where("tweet_id=?", tweet.ID).Order("created_at desc").Find(&comments)
	return
}

func GetLikeCount(tweet *model.Tweet) (count int64) {
	global.GORM_DB.Where("tweet_id = ?", tweet.ID).Find(&model.Like{}).Count(&count)
	return
}

func GetCommentCount(tweet *model.Tweet) int64 {
	return global.GORM_DB.Model(&tweet).Association("Comments").Count()
}

func IsTweetMine(c *gin.Context, tweet *model.Tweet) bool {
	var count int64
	currentUser := CurrentUser(c)
	global.GORM_DB.Where("user_id = ? AND id = ?", currentUser.ID, tweet.ID).Find(&model.Tweet{}).Count(&count)
	return count > 0
}

func IsLiked(c *gin.Context, tweet *model.Tweet) bool {
	var count int64
	currentUser := CurrentUser(c)
	global.GORM_DB.Where("user_id = ? AND tweet_id = ?", currentUser.ID, tweet.ID).Find(&model.Like{}).Count(&count)
	return count > 0
}

func IsRetweet(c *gin.Context, tweet *model.Tweet) bool {
	var count int64
	currentUser := CurrentUser(c)
	global.GORM_DB.Where("user_id = ? AND tweet_id = ?", currentUser.ID, tweet.ID).Find(&model.Retweet{}).Count(&count)
	return count > 0
}

func GetRetweetCount(tweet *model.Tweet) (count int64) {
	global.GORM_DB.Where("tweet_id = ?", tweet.ID).Find(&model.Retweet{}).Count(&count)
	return
}

func GetTweetByID(c *gin.Context, id string) Response {
	var tweet model.Tweet
	tweetId, _ := strconv.Atoi(id)
	global.GORM_DB.Where("id = ?", uint(tweetId)).Find(&tweet)
	commets := GetComments(&tweet)
	tweet.Comments = commets

	return Response{
		Data: BuildAPITweetByID(c, tweet),
		Msg:  "got tweet by ID",
	}
}
