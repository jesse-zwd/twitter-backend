package service

import (
	"backend/global"
	"backend/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (service *LikeService) Create(c *gin.Context) Response {
	like := model.Like{
		TweetID: service.TweetID,
		UserID:  CurrentUser(c).ID,
	}

	// create like
	if err := global.GORM_DB.Create(&like).Error; err != nil {
		return Response{
			Code:  50001,
			Msg:   "creating like failed",
			Error: err.Error(),
		}
	}

	return Response{
		Data: like,
		Msg:  "like created",
	}
}

func DeleteLike(c *gin.Context, id string) Response {
	tweetId, _ := strconv.Atoi(id)
	// delete like
	if err := global.GORM_DB.Where("tweet_id = ? AND user_id = ?", uint(tweetId), CurrentUser(c).ID).Delete(&model.Like{}).Error; err != nil {
		return Response{
			Code:  50001,
			Msg:   "deleteting like failed",
			Error: err.Error(),
		}
	}

	return Response{
		Msg: "like deleted",
	}
}
