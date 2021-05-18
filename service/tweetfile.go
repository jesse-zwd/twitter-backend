package service

import (
	"backend/global"
	"backend/model"
)

func CreateTweetFile(url string, id uint) Response {
	tweetfile := model.TweetFile{
		Url:     url,
		TweetID: id,
	}

	if err := global.GORM_DB.Create(&tweetfile).Error; err != nil {
		return Response{
			Code:  50001,
			Msg:   "creating tweetfile failed",
			Error: err.Error(),
		}
	}

	return Response{
		Msg: "tweetfile created",
	}
}
