package service

import (
	"backend/global"
	"backend/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (service *RetweetService) Create(c *gin.Context) Response {
	retweet := model.Retweet{
		TweetID: service.TweetID,
		UserID:  CurrentUser(c).ID,
	}

	// create retweet
	if err := global.GORM_DB.Create(&retweet).Error; err != nil {
		return Response{
			Code:  50001,
			Msg:   "creating retweet failed",
			Error: err.Error(),
		}
	}

	return Response{
		Data: retweet,
		Msg:  "retweet created",
	}
}

func DeleteRetweet(c *gin.Context, id string) Response {
	tweetId, _ := strconv.Atoi(id)
	// delete retweet
	if err := global.GORM_DB.Where("tweet_id = ? AND user_id = ?", uint(tweetId), CurrentUser(c).ID).Delete(&model.Retweet{}).Error; err != nil {
		return Response{
			Code:  50001,
			Msg:   "deleteting retweet failed",
			Error: err.Error(),
		}
	}

	return Response{
		Msg:  "retweet deleted",
	}
}
