package service

import (
	"backend/global"
	"backend/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (service *FollowService) Create(c *gin.Context) Response {
	follow := model.Follow{
		Following: service.Following,
		Follower:  CurrentUser(c).ID,
	}

	if err := global.GORM_DB.Create(&follow).Error; err != nil {
		return Response{
			Code:  50001,
			Msg:   "creating follow failed",
			Error: err.Error(),
		}
	}

	return Response{
		Data: follow,
		Msg:  "follow created",
	}
}

func DeleteFollow(c *gin.Context, id string) Response {
	following, _ := strconv.Atoi(id)

	if err := global.GORM_DB.Where("follower = ? AND following = ?", CurrentUser(c).ID, uint(following)).Delete(&model.Follow{}).Error; err != nil {
		return Response{
			Code:  50001,
			Msg:   "deleting follow failed",
			Error: err.Error(),
		}
	}

	return Response{
		Msg:  "follow deleted",
	}
}
